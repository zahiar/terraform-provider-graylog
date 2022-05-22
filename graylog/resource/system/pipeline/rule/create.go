package rule

import (
	"context"
	"fmt"

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

	ds, _, err := cl.PipelineRule.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a pipeline rule: %w", err)
	}
	d.SetId(ds[keyID].(string))
	return nil
}
