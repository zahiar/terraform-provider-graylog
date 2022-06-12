package token

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(
	ctx context.Context, user_id string, id string, api_version string,
) (map[string]interface{}, *http.Response, error) {

	body := map[string][]map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/users/"+user_id+"/tokens",
		ResponseBody: &body,
	})
		
	tokens := body["tokens"]
	for token := range tokens {
		if tokens[token]["id"] == id {
			return tokens[token], resp, err
		}
	}
	return nil, resp, err
}

func (cl Client) Create(ctx context.Context, token map[string]interface{}) (map[string]interface{},*http.Response, error) {
	if token == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "POST",
		Path:        "/users/"+token["user_id"].(string)+"/tokens/"+token["name"].(string),
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Delete(ctx context.Context, user_id string, id string) (*http.Response, error) {
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/users/" + user_id + "/tokens/" + id,
	})
	return resp, err
}
