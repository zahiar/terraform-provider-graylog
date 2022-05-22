package indexset

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
	delete(data, keyDefault)

	is, _, err := cl.IndexSet.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a index set: %w", err)
	}
	d.SetId(is[keyID].(string))
	return nil
}
