package dashboard

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/convert"
	"github.com/zahiar/terraform-provider-graylog/graylog/resource/dashboard"
	"github.com/zahiar/terraform-provider-graylog/graylog/util"
)

const (
	keyID          = "id"
	keyDashboardID = "dashboard_id"
	keyDashboards  = "dashboards"
	keyViews       = "views"
	keyTitle       = "title"
)

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	id, ok := util.RenameKey(data, keyID, keyDashboardID)

	if err := convert.SetResourceData(d, dashboard.Resource(), data); err != nil {
		return err
	}

	if ok {
		d.SetId(id.(string))
	}
	return nil
}
