package extractor

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

	inputID := d.Get(keyInputID).(string)
	eID := d.Get(keyExtractorID).(string)
	data, resp, err := cl.Extractor.Get(ctx, inputID, eID)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf(
				"failed to get a extractor (input id: %s, id: %s): %w", inputID, eID, err))
	}
	return setDataToResourceData(d, data)
}
