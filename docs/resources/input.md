# Resource: graylog_input

## Example Usage
```hcl
resource "graylog_input" "gelf_udp" {
  title  = "gelf udp"
  type   = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
  global = true

  attributes = jsonencode({
    bind_address          = "0.0.0.0"
    port                  = 12201
    recv_buffer_size      = 262144
    decompress_size_limit = 8388608
  })
}
```
```hcl
resource "graylog_input" "json_path" {
  title  = "json path"
  type   = "org.graylog2.inputs.misc.jsonpath.JsonPathInput"
  global = "true"

  attributes = jsonencode({
    interval           = 30
    path               = "$.userId"
    throttling_allowed = true
    target_url         = "http://jsonplaceholder.typicode.com/posts/1"
    source             = "id"
    timeunit           = "SECONDS"
  })
}
```
```hcl
resource "graylog_input" "kube_events" {
  title = "Kube Events Input"
  type  = "org.graylog2.inputs.raw.tcp.RawTCPInput"

  global = true

  attributes = jsonencode({
    bind_address = "0.0.0.0"
    port         = 5555
  })
}
```

## Argument Reference

* `title` - (Required) the title of the Input. The data type is `string`.
* `type` - (Required) the type of the Input. The data type is `string`.
* `attributes` - (Required) the attributes of the Input. The data type is `JSON string`.
* `global` - (Optional) The data type is `bool`. The default value is `false`.
* `node` - (Optional) The data type is `string`.

## Attributes Reference

* `created_at` - The date time when the Index Set is created. The data type is `string`.
* `creator_user_id` - The user id who created the Input. The data type is `string`.

## Import

`graylog_input` can be imported using the Input id, e.g.

```console
$ terraform import graylog_input.test 5c4acaefc9e77bbbbbbbbbbb
```
