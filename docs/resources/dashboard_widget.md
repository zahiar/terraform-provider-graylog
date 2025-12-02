# Resource: graylog_dashboard_widget

## Example Usage
```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count change"
  dashboard_id = graylog_dashboard.test.id
  type         = "STREAM_SEARCH_RESULT_COUNT"
  cache_time   = 10
  config = jsonencode({
    timerange = {
      type  = "relative"
      range = 400
    }
    lower_is_better = true
    trend           = true
    stream_id       = graylog_stream.test.id
    query           = ""
  })
}
```
```hcl
resource "graylog_dashboard_widget" "test2" {
  description  = "Quick values"
  dashboard_id = graylog_dashboard.test.id
  type         = "QUICKVALUES"
  cache_time   = 10

  config = jsonencode({
    timerange = {
      type  = "relative"
      range = 300
    }
    stream_id        = graylog_stream.test.id
    query            = ""
    field            = "status"
    show_data_table  = true
    show_pie_chart   = true
    limit            = 5
    sort_order       = "desc"
    data_table_limit = 60
  })
}
```
```hcl
resource "graylog_dashboard_widget" "stacked_chart" {
  description  = "stacked chart"
  dashboard_id = graylog_dashboard.test.id
  type         = "STACKED_CHART"
  cache_time   = 10

  config = jsonencode({
    interval = "hour"
    timerange = {
      type  = "relative"
      range = 86400
    },
    renderer      = "bar"
    interpolation = "linear"
    series = [
      {
        query                = ""
        field                = "AccessMask"
        statistical_function = "count"
      }
    ]
  })
}
```

## Argument Reference

* `type` - (Required) widget type. The data type is `string`.
* `description` - (Required) description of the widget. The data type is `string`.
* `dashboard_id` - (Required, Forces new resource) id of the dashboard which the widget is associated with. The data type is `string`.
* `configuration` - (Required) configuration of the widget. The data type is `JSON string`.
* `cache_time` - (Optional) The data type is `int`.

## Attributes Reference

* `creator_user_id` - The user id who created the widget. The data type is `string`.

## Import

`graylog_dashboard_widget` can be imported using `<Dashboard id>/<Widget id>`, e.g.

```console
$ terraform import graylog_dashboard_widget.test 5c4acaefc9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
