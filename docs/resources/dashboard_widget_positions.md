# Resource: graylog_dashboard_widget_positions

## Example Usage
```hcl
resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = graylog_dashboard_widget.test.dashboard_id

  positions = jsonencode({
    "${graylog_dashboard_widget.test.widget_id}" = {
      row    = 0
      col    = 0
      height = 1
      width  = 1
    }
    "${graylog_dashboard_widget.test2.widget_id}" = {
      row    = 0
      col    = 1
      height = 2
      width  = 2
    }
  })
}
```

## Argument Reference

* `dashboard_id` - (Required, Forces new resource) id of the dashboard which widgets are associated with. The data type is `string`.
* `positions` - (Required) positions of widgets. The data type is `JSON string`.

### Attributes Reference

Nothing.

### Import

`graylog_dashboard_widget_positions` can be imported using the Dashboard id, e.g.

```console
$ terraform import graylog_dashboard_widget_positions.test 5c4acaefc9e77bbbbbbbbbbb
```
