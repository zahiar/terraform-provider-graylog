package alarmcallback

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

	streamID := data[keyStreamID].(string)
	delete(data, keyID)
	delete(data, keyStreamID)
	delete(data, keyAlarmCallbackID)
	ac, _, err := cl.AlarmCallback.Create(ctx, streamID, data)
	if err != nil {
		return fmt.Errorf("failed to create a alarm callback (stream id: %s): %w", streamID, err)
	}
	acID := ac[keyAlarmCallbackID].(string)
	if err := d.Set(keyAlarmCallbackID, acID); err != nil {
		return err
	}
	d.SetId(streamID + "/" + acID)
	return nil
}
