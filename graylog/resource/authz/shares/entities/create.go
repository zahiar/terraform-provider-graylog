package entities

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
	_, err = cl.AuthzSharesEntities.Post(ctx, data)
	if err != nil {
		return err
	}
	d.SetId(data["entity_type"].(string) + ":" + data["entity_id"].(string))
	return nil
}
