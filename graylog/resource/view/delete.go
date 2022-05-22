package view

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
	if _, err := cl.View.Delete(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
