# Resource: graylog_role

## Example Usage
```hcl
resource "graylog_role" "read-stream-test" {
  name        = "read-stream-test"
  description = "read the stream 'test'"

  permissions = [
    "streams:read:${graylog_stream.test.id}"
  ]
}
```
```hcl
resource "graylog_role" "terraform" {
  name        = "terraform"
  description = "terraform"

  permissions = [
    "dashboards:*",
    "indexsets:*",
    "inputs:*",
    "roles:*",
    "streams:*",
    "users:*",
    "pipeline_rule:*",
  ]
}
```

## Argument Reference

* `name` - (Required) The Role name. The data type is `string`.
* `permissions` - (Required) The permissions of the Role. The data type is `[]string`.
* `description` - (Optional) description of the Role. The data type is `string`.

## Attributes Reference

* `read_only` - The data type is `bool`.

## Import

`graylog_role` can be imported using the Role name, e.g.

```
$ terraform import graylog_role.foo foo
```
