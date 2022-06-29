package entities

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	data["id"] = d.Id()
	data["selected_grantee_capabilities"] = map[string]interface{}{}
	if _, err := cl.AuthzSharesEntities.Post(ctx, data); err != nil {
		return err
	}
	return nil
}
