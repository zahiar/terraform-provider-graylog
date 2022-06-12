package token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/client"
)

func create(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	data, _, err = cl.UserToken.Create(ctx, data)
	if err != nil {
		return err
	}
	err = d.Set("token", data["token"])
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))
	return nil
}
