package indexset

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/client"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	if _, err := cl.IndexSet.Delete(ctx, d.Id()); err != nil {
		return fmt.Errorf("failed to delete a index set %s: %w", d.Id(), err)
	}
	return nil
}
