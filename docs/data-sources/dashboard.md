# Data Source: graylog_dashboard

## Example Usage
```hcl
data "graylog_dashboard" "test" {
  title = "test"
}
```
```hcl
data "graylog_dashboard" "test" {
  dashboard_id = "123test"
}
```

## Argument Reference
One of `dashboard_id` or `title` must be set.

## Attributes Reference
* `title` - The data type is `string`
* `dashboard_id` - The data type is `string`
* `description` - The data type is `string`
