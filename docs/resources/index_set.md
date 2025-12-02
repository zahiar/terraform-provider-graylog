# Resource: graylog_index_set

## Example Usage
```hcl
resource "graylog_index_set" "test" {
  title                               = "test"
  index_prefix                        = "terraform-provider-graylog-test"
  rotation_strategy_class             = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  retention_strategy_class            = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  description                         = "test"
  index_analyzer                      = "standard"
  index_optimization_disabled         = true
  writable                            = true
  shards                              = 4
  replicas                            = 0
  index_optimization_max_num_segments = 1
  field_type_refresh_interval         = 5000

  retention_strategy = jsonencode({
    max_number_of_indices = 30
    type                  = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
  })

  rotation_strategy = jsonencode({
    max_docs_per_index = 30000000
    type               = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
  })
}
```

## Argument Reference

* `title` - (Required) the title of the Index Set. The data type is `string`.
* `index_prefix` - (Required, Forces new resource) the index prefix of the Index Set. The data type is `string`.
* `rotation_strategy_class` - (Required) the rotation strategy class of the Index Set. The data type is `string`.
* `rotation_strategy` - (Required) the rotation strategy of the Index Set. The data type is `JSON string`.
* `retention_strategy_class` - (Required) the retention strategy class of the Index Set. The data type is `string`.
* `retention_strategy` - (Required) the retention strategy of the Index Set. The data type is `JSON string`.
* `index_analyzer` - (Required) the Index Analyzer of the Index Set. The data type is `string`.
* `shards` - (Required) the number of shards of the Index Set. The data type is `int`.
* `description` - (Optional) the description of the Index Set. The data type is `string`.
* `replicas` - (Optional) the number of the replicas of the Index Set. The data type is `int`.
* `index_optimization_disabled` - (Optional) The data type is `bool`.
* `index_optimization_max_num_segments` - (Required) The data type is `int`.
* `default` - (Optional) The data type is `bool`.
* `field_type_refresh_interval` - (Optional) The data type is `int`.

## Attributes Reference

* `writable` - The data type is `bool`.
* `creation_date` - The date time when the Index Set is created. The data type is `string`.

## Import

`graylog_index_set` can be imported using the Index Set id, e.g.

```console
$ terraform import graylog_index_set.test 5c4acaefc9e77bbbbbbbbbbb
```
