package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	var update_id string
	var save_full_name interface{}
	if err != nil {
		return err
	}
	oldName, newName := d.GetChange(keyUsername)
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	if cl.APIVersion == "v4" {
		save_full_name = data["full_name"]
		delete(data, "full_name")
		update_id = d.Id()
	} else {
		delete(data, "first_name")
		delete(data, "last_name")
		update_id = oldName.(string)
	}
	if _, ok := data[keyPermissions]; !ok {
		data[keyPermissions] = []string{}
	}
	if _, err := cl.User.Update(ctx, update_id, data); err != nil {
		return err
	}
	if cl.APIVersion == "v4" {
		err = d.Set("full_name", save_full_name)
		if err != nil {
			return err
		}
	} else {
		d.SetId(newName.(string))
	}
	return nil
}
