package sidecar

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/client"
	"github.com/zahiar/terraform-provider-graylog/graylog/util"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.Sidecar.GetAll(ctx)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf("failed to get all sidecars: %w", err))
	}
	return setDataToResourceData(d, data)
}
