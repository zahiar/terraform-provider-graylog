package entities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zahiar/terraform-provider-graylog/graylog/convert"
)

const grn_base = "grn::::"

func parseGrn(grn string) (string, string, error) {
	parts := strings.SplitN(grn, ":", 6)
	if len(parts) != 6 || parts[0] != "grn" || parts[4] == "" || parts[5] == "" {
		return "", "", fmt.Errorf("unexpected format of (%s), expected: grn::::type:id, got grn::::%s:%s", grn, parts[4], parts[5])
	}
	return parts[4], parts[5], nil
}

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	new_data := make(map[string]interface{})
	if err != nil {
		return nil, err
	}
	new_data["id"] = data["entity_type"].(string) + ":" + data["entity_id"].(string)
	new_data["entity_type"] = data["entity_type"]
	new_data["entity_id"] = data["entity_id"]
	payload := make(map[string]string, 0)
	for k := range data["grant"].([]interface{}) {
		entity_type := data["grant"].([]interface{})[k].(map[string]interface{})["entity_type"].(string)
		entity_id := data["grant"].([]interface{})[k].(map[string]interface{})["entity_id"].(string)
		capability := data["grant"].([]interface{})[k].(map[string]interface{})["right"].(string)
		payload[fmt.Sprintf("%s%s:%s", grn_base, entity_type, entity_id)] = capability
	}
	new_data["selected_grantee_capabilities"] = payload
	return new_data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	new_data := make(map[string]interface{})
	active_shares := data["active_shares"]
	new_data_grant := make([]map[string]interface{}, 0)
	for e := range active_shares.([]interface{}) {
		share := active_shares.([]interface{})[e]
		grant := make(map[string]interface{})
		entity_type, entity_id, err := parseGrn(share.(map[string]interface{})["grantee"].(string))
		if err != nil {
			return err
		}
		grant["entity_type"] = entity_type
		grant["entity_id"] = entity_id
		grant["right"] = share.(map[string]interface{})["capability"]
		new_data_grant = append(new_data_grant, grant)
	}
	new_data["grant"] = new_data_grant
	entity_type, entity_id, err := parseGrn(data["entity"].(string))
	if err != nil {
		return err
	}

	new_data["entity_type"] = entity_type
	new_data["entity_id"] = entity_id
	if err := convert.SetResourceData(d, Resource(), new_data); err != nil {
		return err
	}

	d.SetId(new_data["entity_type"].(string) + ":" + new_data["entity_id"].(string))
	return nil
}
