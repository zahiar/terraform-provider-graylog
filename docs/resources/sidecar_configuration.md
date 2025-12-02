# Resource: graylog_sidecar_configuration

## Example Usage
```hcl
resource "graylog_sidecar_configuration" "test" {
  name         = "foo"
  color        = "#00796b"
  collector_id = graylog_sidecar_collector.test.id
  template     = <<EOF
fields_under_root: true
EOF
}
```

## Argument Reference

* `name` - (Required) The name of the Sidecar Configuration. The data type is `string`.
* `color` - (Required) The data type is `string`.
* `collector_id` - (Required) The data type is `string`.
* `template` - (Required) The data type is `string`.

## Attributes Reference

None.

## Import

`graylog_sidecar_configuration` can be imported using the Collector id, e.g.

```console
$ terraform import graylog_sidecar_configuration.test 5c4acaefc9e77bbbbbbbbbbb
```
