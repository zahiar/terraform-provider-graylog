# Resource: graylog_pipeline_connection

## Example Usage
```hcl
resource "graylog_pipeline_connection" "test" {
  stream_id    = graylog_stream.test.id
  pipeline_ids = [graylog_pipeline.test.id]
}
```

## Argument Reference

* `stream_id` - (Required, Forces new resource) The stream id which the Pipelines are associated with. The data type is `string`.
* `pipeline_ids` - (Required) The pipeline ids. The data type is `[]string`.

### Note

This resource treats the stream id as the resource id,
because there is no Graylog API to operate resource by connection pipeline id.
So please make the stream id unique in all `graylog_pipeline_connection` resources.

## Attributes Reference

Nothing.

## Import

`graylog_pipeline_connection` can be imported using the Stream id, e.g.

```console
$ terraform import graylog_pipeline_connection.test <stream id>
```
