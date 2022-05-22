package widget

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
	dsID := d.Get(keyDashboardID).(string)
	wID := d.Get(keyWidgetID).(string)
	data, resp, err := cl.DashboardWidget.Get(ctx, dsID, wID)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf(
				"failed to get a dashboard widget(dashboard id: %s widget id: %s): %w", dsID, wID, err))
	}
	return setDataToResourceData(d, data)
}
