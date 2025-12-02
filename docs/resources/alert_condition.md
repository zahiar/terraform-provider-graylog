# Resource: graylog_alert_condition

## Example Usage
```hcl
resource "graylog_alert_condition" "test" {
  type      = "field_content_value"
  stream_id = graylog_stream.test.id
  in_grace  = false
  title     = "test"

  parameters = jsonencode({
    field                = "message"
    value                = "hoge hoge"
    backlog              = 2
    repeat_notifications = false
    query                = "*"
    grace                = 0
  })
}
```

## Argument Reference

* `type` - (Required) Alert Condition type. The data type is `string`.
* `title` - (Required) Alert Condition title. The data type is `string`.
* `stream_id` - (Required, Forces new resource) Stream id which the Alert Condition is associated with. The data type is `string`.
* `parameters` - (Required) parameters of Alert Condition. The data type is `JSON string`.
* `in_grace` - (Optional) The data type is `bool`. The default value is `false`.

### parameters

`parameters` is a JSON string whose type is `object`.
The data structure of JSON is different per AlertCondition `type`.
Please see the example above.

### Attributes Reference

* `alert_condition_id` - Alert Condition id. The data type is `string`.

## Import

`graylog_alert_condition` can be imported using the User `<stream id>/<alert condition id>`, e.g.

```console
$ terraform import graylog_alert_condition.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
