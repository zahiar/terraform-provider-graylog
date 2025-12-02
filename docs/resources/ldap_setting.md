# Resource: graylog_ldap_setting

## Example Usage
```hcl
resource "graylog_ldap_setting" "system" {
  enabled                   = false
  system_username           = ""
  ldap_uri                  = "ldap://localhost:389"
  use_start_tls             = false
  trust_all_certificates    = false
  active_directory          = false
  search_base               = ""
  search_pattern            = ""
  display_name_attribute    = ""
  default_group             = "Reader"
  group_mapping             = null
  group_search_base         = null
  group_id_attribute        = null
  additional_default_groups = null
  group_search_pattern      = null
}
```

## Argument Reference

* `system_username` - (Required) The data type is `string`.
* `ldap_uri` - (Required) The data type is `string`.
* `search_base` - (Required) The data type is `string`.
* `search_pattern` - (Required) The data type is `string`.
* `display_name_attribute` - (Required) The data type is `string`.
* `default_group` - (Required) The data type is `string`.
* `description` - (Optional) The data type is `string`.
* `enabled` - (Optional) The data type is `bool`.
* `use_start_tls` - (Optional) The data type is `bool`.
* `trust_all_certificates` - (Optional) The data type is `bool`.
* `active_directory` - (Optional) The data type is `bool`.
* `group_search_base` - (Optional) The data type is `string`.
* `group_id_attribute` - (Optional) The data type is `string`.
* `group_search_pattern` - (Optional) The data type is `string`.
* `group_mapping` - (Optional) The data type is `map[string]string`.
* `system_password` - (Optional) The data type is `string`.

Note that `system_passoword` is optional as Terraform schema but is required to create a LDAP setting.
If we make `system_password` required as Terrafrom schema, we have to store `system_password` in the Terraform state file, which some users wouldn't want it.

## Attributes Reference

* `system_password_set` - The data type is `bool`.

## Import

Unlike other resources, LDAP settings has no id,
so when you import the LDAP settings, please specify some string as id.

```console
$ terraform import graylog_ldap_setting.foo bar
```
