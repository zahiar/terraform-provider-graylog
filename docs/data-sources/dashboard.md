# Data Source: graylog_dashboard

* [Example](https://github.com/zahiar/terraform-provider-graylog/blob/master/examples/v0.12/dashboard.tf)
* [Source Code](https://github.com/zahiar/terraform-provider-graylog/blob/master/graylog/datasource/dashboard/data_source.go)

## Example Usage

```tf
data "graylog_dashboard" "test" {
  title = "test"
}
```

## Argument Reference

One of `dashboard_id` or `title` must be set.

## Attributes Reference

* `title` - The data type is `string`
* `dashboard_id` - The data type is `string`
* `description` - The data type is `string`
