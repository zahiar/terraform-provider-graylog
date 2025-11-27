# Data Source: graylog_sidecar

## Example Usage
```hcl
data "graylog_sidecar" "test" {
  node_name = "test"
}
```
```hcl
data "graylog_sidecar" "test" {
  node_id = "123"
}
```

## Argument Reference
One of `node_id` or `node_name` must be set.

## Attributes Reference
* `node_id` - The data type is `string`.
* `node_name` - The data type is `string`.
