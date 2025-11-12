# Data Source: graylog_sidecar

* [Example](https://github.com/zahiar/terraform-provider-graylog/blob/master/examples/v0.12/sidecar.tf)
* [Source Code](https://github.com/zahiar/terraform-provider-graylog/blob/master/graylog/datasource/sidecar/data_source.go)


## Example Usage

```tf
data "graylog_sidecar" "test" {
  node_name = "test"
}
```

## Argument Reference

One of `node_id` or `node_name` must be set.

## Attributes Reference

* `node_id` - The data type is `string`.
* `node_name` - The data type is `string`.
