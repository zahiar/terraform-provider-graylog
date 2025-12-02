# Resource: graylog_input_static_fields

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

resource "graylog_input_static_fields" "gelf_udp" {
  input_id = graylog_input.gelf_udp.id
  fields = {
    foo = "bar"
  }
}
```

## Argument Reference

* `input_id` - (Required) id of the Input which the static fields are associated with. The data type is `string`.
* `fields` - (Optional) The data type is `map[string]string`.

## Attributes Reference

None.

## Import

`graylog_input_static_fields` can be imported using the Input id, e.g.

```console
$ terraform import graylog_input_static_fields.test 5c4acaefc9e77bbbbbbbbbbb
```
