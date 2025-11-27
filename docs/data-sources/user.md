# Data Source: graylog_user

Please use `"v4"` version of API when using a graylog version >=4.1.

## Example Usage
```hcl
data "graylog_user" "test" {
  username = "test"
}
```

## Argument Reference
* `username` - (Required) The data type is `string`.

## Attributes Reference
* `email` - The data type is `string`.
* `full_name` - The data type is `string`. For v3 api only.
* `first_name` - The data type is `string`. For v4 api only.
* `last_name` -  The data type is `string`. For v4 api only.
* `roles` - The data type is `set of string`.
* `timezone` - The data type is `string`.
* `session_timeout_ms` - The data type is `int`. The default value is `3600000`.
* `user_id` - The data type is `string`.
* `read_only` - The data type is `bool`.
* `external` - The data type is `bool`.
* `client_address` - The data type is `string`.
* `session_active` - The data type is `bool`.
* `last_activity` - The data type is `string`.
* `permissions` - The data type is `set of string`.