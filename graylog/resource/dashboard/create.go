package dashboard

import (
	"context"
	"errors"
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

	ds, _, err := cl.Dashboard.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a dashboard: %w", err)
	}

	a, ok := ds[keyDashboardID]
	if !ok {
		return errors.New("response body of Graylog API is unexpected. 'dashboard_id' isn't found")
	}
	dID, ok := a.(string)
	if !ok {
		return fmt.Errorf(
			"response body of Graylog API is unexpected. 'dashboard_id' should be string: %v", a)
	}

	d.SetId(dID)
	return nil
}
