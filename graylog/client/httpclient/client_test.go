package httpclient

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/v2/flute"
)

func TestError_Error(t *testing.T) {
	data := []struct {
		title string
		exp   string
		err   Error
	}{
		{
			title: "normal",
			err: Error{
				statusCode: 500,
				bodyByte:   []byte(`{"error": "Internal Server Error"}`),
				err:        errors.New("status code >= 300"),
			},
			exp: `status code: 500, {"error": "Internal Server Error"}: status code >= 300`,
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			require.Equal(t, d.exp, d.err.Error())
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	exp := errors.New("foo")
	err := &Error{
		statusCode: 500,
		bodyByte:   []byte(`{"error": "Internal Server Error"}`),
		err:        exp,
	}
	if e := errors.Unwrap(err); e != exp {
		t.Errorf(`errors.Unwrap(err) = %v; want %v`, e, exp)
	}
}

func TestClient_Call(t *testing.T) {
	client := New("http://example.com/api")
	token := "xxx"
	client.SetRequest = func(req *http.Request) error {
		req.Header.Add("Authorization", "token "+token)
		return nil
	}
	client.Timeout = 1 * time.Second

	userAgent := "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
	routeCreateUser := flute.Route{
		Name: "create a user",
		Matcher: flute.Matcher{
			Method: "POST",
			Path:   "/api/users",
		},
		Tester: flute.Tester{
			BodyJSONString: `{
			  "name": "foo",
			  "email": "foo@example.com"
			}`,
			Header: http.Header{
				"Authorization": []string{"token " + token},
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
			  "id": 10,
			  "name": "foo",
			  "email": "foo@example.com"
			}`,
		},
	}
	routeCreateUserTimeout := flute.Route{
		Name: "create a user",
		Matcher: flute.Matcher{
			Method: "POST",
			Path:   "/api/users",
		},
		Tester: flute.Tester{
			BodyJSONString: `{
			  "name": "foo",
			  "email": "foo@example.com"
			}`,
			Header: http.Header{
				"Authorization": []string{"token " + token},
			},
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				ctx := req.Context()
				type resp struct {
					resp *http.Response
					err  error
				}
				respChan := make(chan resp, 1)
				go func() {
					time.Sleep(2 * time.Second)
					respChan <- resp{
						resp: &http.Response{
							StatusCode: 201,
							Body: ioutil.NopCloser(strings.NewReader(`{
		        	  "id": 10,
		        	  "name": "foo",
		        	  "email": "foo@example.com"
		        	}`)),
						},
						err: nil,
					}
				}()
				select {
				case resp := <-respChan:
					return resp.resp, resp.err
				case <-ctx.Done():
					return nil, ctx.Err()
				}
			},
		},
	}

	ctx := context.Background()
	data := []struct {
		title            string
		params           CallParams
		routes           []flute.Route
		exp              interface{}
		expErrorResponse interface{}
		isErr            bool
	}{
		{
			title: "request body is struct",
			params: CallParams{
				Method: "POST",
				Path:   "/users",
				RequestBody: struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}{
					Name:  "foo",
					Email: "foo@example.com",
				},
				ResponseBody: &map[string]interface{}{},
			},
			routes: []flute.Route{routeCreateUser},
			exp: &map[string]interface{}{
				"id":    10.0,
				"name":  "foo",
				"email": "foo@example.com",
			},
		},
		{
			title: "request body is string",
			params: CallParams{
				Method:       "POST",
				Path:         "/users",
				RequestBody:  `{"name": "foo", "email": "foo@example.com"}`,
				ResponseBody: &map[string]interface{}{},
			},
			routes: []flute.Route{routeCreateUser},
			exp: &map[string]interface{}{
				"id":    10.0,
				"name":  "foo",
				"email": "foo@example.com",
			},
		},
		{
			title: "request body is []byte",
			params: CallParams{
				Method:       "POST",
				Path:         "/users",
				RequestBody:  []byte(`{"name": "foo", "email": "foo@example.com"}`),
				ResponseBody: &map[string]interface{}{},
			},
			routes: []flute.Route{routeCreateUser},
			exp: &map[string]interface{}{
				"id":    10.0,
				"name":  "foo",
				"email": "foo@example.com",
			},
		},
		{
			title: "error response",
			params: CallParams{
				Method:            "GET",
				Path:              "/groups/foo",
				ResponseErrorBody: &map[string]interface{}{},
			},
			expErrorResponse: &map[string]interface{}{
				"error": "group foo isn't found",
			},
			routes: []flute.Route{
				{
					Name: "get a group",
					Matcher: flute.Matcher{
						Method: "GET",
						Path:   "/api/groups/foo",
					},
					Response: flute.Response{
						Base: http.Response{
							StatusCode: 404,
						},
						BodyString: `{
						  "error": "group foo isn't found"
						}`,
					},
				},
			},
			isErr: true,
		},
		{
			title: "error response with query and header",
			params: CallParams{
				Method: "GET",
				Path:   "/groups",
				Header: http.Header{
					"User-Agent": []string{userAgent},
				},
				Query: url.Values{
					"name": []string{"foo"},
				},
				ResponseErrorBody: &map[string]interface{}{},
			},
			routes: []flute.Route{
				{
					Name: "get a group with query",
					Matcher: flute.Matcher{
						Method: "GET",
						Path:   "/api/groups",
						Header: http.Header{
							"User-Agent":    []string{userAgent},
							"Authorization": []string{"token " + token},
						},
						Query: url.Values{
							"name": []string{"foo"},
						},
					},
					Response: flute.Response{
						Base: http.Response{
							StatusCode: 404,
						},
						BodyString: `{
						  "error": "group foo isn't found"
						}`,
					},
				},
			},
			expErrorResponse: &map[string]interface{}{
				"error": "group foo isn't found",
			},
			isErr: true,
		},
		{
			title: "client timeout",
			params: CallParams{
				Method: "POST",
				Path:   "/users",
				RequestBody: struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}{
					Name:  "foo",
					Email: "foo@example.com",
				},
				ResponseBody: &map[string]interface{}{},
			},
			routes: []flute.Route{routeCreateUserTimeout},
			isErr:  true,
		},
		{
			title: "params imeout",
			params: CallParams{
				Method: "POST",
				Path:   "/users",
				RequestBody: struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}{
					Name:  "foo",
					Email: "foo@example.com",
				},
				ResponseBody: &map[string]interface{}{},
				Timeout:      500 * time.Millisecond,
			},
			routes: []flute.Route{routeCreateUserTimeout},
			isErr:  true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client.HTTPClient.Transport = &flute.Transport{
				T: t,
				Services: []flute.Service{
					{
						Endpoint: "http://example.com",
						Routes:   d.routes,
					},
				},
			}

			_, err := client.Call(ctx, d.params) //nolint:bodyclose
			if d.isErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, d.exp, d.params.ResponseBody)
			}
			require.Equal(t, d.expErrorResponse, d.params.ResponseErrorBody)
			if err != nil {
				var e *Error
				if errors.As(err, &e) {
					require.Equal(t, d.expErrorResponse, e.Body())
				}
			}
		})
	}
}
