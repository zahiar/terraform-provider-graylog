package entities

import (
	"context"
	"net/http"

	"github.com/zahiar/terraform-provider-graylog/graylog/client/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(
	ctx context.Context, id string,
) (map[string]interface{}, *http.Response, error) {
	req := map[string]string{}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/authz/shares/entities/grn::::" + id + "/prepare",
		RequestBody:  &req,
		ResponseBody: &body,
	})

	return body, resp, err
}

func (cl Client) Post(
	ctx context.Context, data map[string]interface{},
) (*http.Response, error) {
	req := map[string]interface{}{"selected_grantee_capabilities": data["selected_grantee_capabilities"]}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "POST",
		Path:        "/authz/shares/entities/grn::::" + data["id"].(string),
		RequestBody: &req,
	})

	return resp, err
}
