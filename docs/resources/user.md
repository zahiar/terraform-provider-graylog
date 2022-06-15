# Resource: graylog_user

* [Example](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/examples/v0.12/user.tf)
* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/user/resource.go)

Please use `"v4"` version of API when using a graylog version >=4.1.

## Argument Reference

* `username` - (Required, Forces new resource) The data type is `string`.
* `email` - (Required) The data type is `string`.
* `full_name` - (Required for v3) The data type is `string`. For v3 api only, incompatible with `first_name` and `last_name`.
* `first_name` - (Required for v4) The data type is `string`. For v4 api only, incompatible with `full_name`.
* `last_name` - (Required for v4) The data type is `string`. For v4 api only, incompatible with `full_name`.
* `password` - (Optonal, Sensitive) The data type is `string`.
* `roles` - (Optional) The data type is `set of string`.
* `timezone` - (Optional, Computed) The data type is `string`.
* `session_timeout_ms` - (Optional) The data type is `int`. The default value is `3600000`.

### password

`password` is required to create a resource.
Once the user is created, `password` is optional.

## Attributes Reference

* `user_id` - The data type is `string`.
* `read_only` - The data type is `bool`.
* `external` - The data type is `bool`.
* `client_address` - The data type is `string`.
* `session_active` - The data type is `bool`.
* `last_activity` - The data type is `string`.
* `permissions` - The data type is `set of string`.

## Import

`graylog_user` can be imported using the User `username`, e.g.

```
$ terraform import graylog_user.foo foo
```
