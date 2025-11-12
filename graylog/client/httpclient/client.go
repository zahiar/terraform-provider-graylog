package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Endpoint   string
	HTTPClient *http.Client
	SetRequest func(req *http.Request) error
	Timeout    time.Duration
}

func New(endpoint string) Client {
	return Client{
		Endpoint:   endpoint,
		HTTPClient: http.DefaultClient,
	}
}

type Error struct {
	statusCode int
	bodyByte   []byte
	body       interface{}
	err        error
}

func (e Error) StatusCode() int {
	return e.statusCode
}

func (e Error) BodyByte() []byte {
	return e.bodyByte
}

func (e Error) Body() interface{} {
	return e.body
}

func (e Error) Error() string {
	a := ""
	if e.err != nil {
		a = e.err.Error()
	}
	return "status code: " + strconv.Itoa(e.statusCode) + ", " + string(e.bodyByte) + ": " + a
}

func (e Error) Unwrap() error {
	return e.err
}

type CallParams struct {
	Method            string
	Path              string
	Header            http.Header
	Query             url.Values
	RequestBody       interface{}
	ResponseBody      interface{}
	ResponseErrorBody interface{}
	Timeout           time.Duration
}

func (client Client) Call(ctx context.Context, params CallParams) (*http.Response, error) {
	if params.Timeout > 0 {
		c, cancel := context.WithTimeout(ctx, params.Timeout)
		defer cancel()
		ctx = c
	} else if client.Timeout > 0 {
		c, cancel := context.WithTimeout(ctx, client.Timeout)
		defer cancel()
		ctx = c
	}
	if client.Endpoint == "" {
		return nil, errors.New("endpoint is required")
	}
	if params.Method == "" {
		return nil, errors.New("method is required")
	}

	var body io.Reader
	if params.RequestBody != nil {
		switch b := params.RequestBody.(type) {
		case string:
			body = strings.NewReader(b)
		case []byte:
			body = bytes.NewBuffer(b)
		default:
			buf := &bytes.Buffer{}
			if err := json.NewEncoder(buf).Encode(params.RequestBody); err != nil {
				return nil, fmt.Errorf("failed to parse the request body as JSON: %w", err)
			}
			body = buf
		}
	}

	path := client.Endpoint + params.Path
	if len(params.Query) != 0 {
		path += "?" + params.Query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, params.Method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}
	for k, list := range params.Header {
		for _, v := range list {
			req.Header.Add(k, v)
		}
	}
	if client.SetRequest != nil {
		if err := client.SetRequest(req); err != nil {
			return nil, fmt.Errorf("failed to set a request: %w", err)
		}
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send a request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode >= 300 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return res, &Error{
				statusCode: res.StatusCode,
				err:        fmt.Errorf("status code >= 300: failed to read a response body: %w", err),
			}
		}
		if params.ResponseErrorBody != nil {
			if err := json.Unmarshal(body, params.ResponseErrorBody); err != nil {
				return res, &Error{
					statusCode: res.StatusCode,
					bodyByte:   body,
					err:        fmt.Errorf("status code >= 300: failed to parse an error response body as JSON: %w", err),
				}
			}
			return res, &Error{
				statusCode: res.StatusCode,
				bodyByte:   body,
				body:       params.ResponseErrorBody,
				err:        errors.New("status code >= 300"),
			}
		}
		return res, &Error{
			statusCode: res.StatusCode,
			bodyByte:   body,
			err:        errors.New("status code >= 300"),
		}
	}

	if params.ResponseBody != nil {
		if err := json.NewDecoder(res.Body).Decode(params.ResponseBody); err != nil {
			return res, fmt.Errorf("failed to read a response body as JSON: %w", err)
		}
		return res, nil
	}
	_, _ = io.Copy(ioutil.Discard, res.Body)
	return res, nil
}
