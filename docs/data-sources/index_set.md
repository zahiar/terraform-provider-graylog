# Data Source: graylog_index_set

## Example Usage
```hcl
data "graylog_index_set" "test" {
  index_set_id = "123"
}
```
```hcl
data "graylog_index_set" "test" {
  title = "test"
}
```
```hcl
data "graylog_index_set" "test" {
  index_prefix = "test_"
}
```

## Argument Reference
One of `index_set_id` or `title` or `index_prefix` must be set.

## Attributes Reference
* `title` - the title of the Index Set. The data type is `string`.
* `index_prefix` - the index prefix of the Index Set. The data type is `string`.
* `rotation_strategy_class` - the rotation strategy class of the Index Set. The data type is `string`.
* `rotation_strategy` - the rotation strategy of the Index Set. The data type is `JSON string`.
* `retention_strategy_class` - the retention strategy class of the Index Set. The data type is `string`.
* `retention_strategy` - the retention strategy of the Index Set. The data type is `JSON string`.
* `index_analyzer` - the Index Analyzer of the Index Set. The data type is `string`.
* `shards` - the number of shards of the Index Set. The data type is `int`.
* `description` - the description of the Index Set. The data type is `string`.
* `replicas` - the number of the replicas of the Index Set. The data type is `int`.
* `index_optimization_disabled` - The data type is `bool`.
* `index_optimization_max_num_segments` - The data type is `int`.
* `default` - The data type is `bool`.
* `writable` - The data type is `bool`.
* `creation_date` - The date time when the Index Set is created. The data type is `string`.
* `id` - Index Set id. The data type is `string`.
* `field_type_refresh_interval` - The data type is `int`.
