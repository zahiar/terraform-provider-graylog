# Resource: graylog_dashboard

## Example Usage
```hcl
resource "graylog_dashboard" "test" {
  title       = "test"
  description = "test"
}
```

## Argument Reference

* `title` - (Required) Dashboard title. The data type is `string`.
* `description` - (Required) Dashboard description. The data type is `string`.

## Attrs Reference

* `created_at` - The date time when the Dashboard is created. The data type is `string`.

## Import

`graylog_dashboard` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard.test 5c4acaefc9e77bbbbbbbbbbb
```
