# Resource: graylog_extractor

## Example Usage
```hcl
resource "graylog_extractor" "test_grok" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test_grok"
  type            = "grok"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  extractor_config = jsonencode({
    grok_pattern = "%%%{DATA}"
  })
}
```
```hcl
resource "graylog_extractor" "test_json" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test"
  type            = "json"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  extractor_config = jsonencode({
    list_separator             = ", "
    kv_separator               = "="
    key_prefix                 = "visit_"
    key_separator              = "_"
    replace_key_whitespace     = false
    key_whitespace_replacement = "_"
  })
}
```
```hcl
resource "graylog_extractor" "test_regex" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test_regex"
  type            = "regex"
  cursor_strategy = "copy"

  source_field   = "message"
  condition_type = "none"
  order          = 0

  extractor_config = jsonencode({
    regex_value = ".*"
  })

  converters {
    type = "date"

    config = jsonencode({
      date_format = "yyyy/MM/ddTHH:mm:ss"
      time_zone   = "Japan"
      locale      = "en"
    })
  }
}
```
```hcl
resource "graylog_extractor" "test_split_and_index" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test_split_and_index"
  type            = "split_and_index"
  cursor_strategy = "copy"

  source_field    = "message"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  extractor_config = jsonencode({
    split_by = "."
    index    = 1
  })
}
```
```hcl
resource "graylog_extractor" "http_response_code" {
  input_id        = graylog_input.gelf_udp.id
  title           = "Apache http_response_code"
  type            = "regex"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "http_response_code"
  condition_type  = "regex"
  condition_value = "[1-5]\\d{2}"
  order           = 0

  converters {
    type   = "numeric"
    config = "{}"
  }

  extractor_config = jsonencode({
    regex_value = "HTTP/1.[0-1]\" (\\d{3}) "
  })
}
```

## Argument Reference

### Required Argument

* `input_id` - (Required) the Input id which the Extractor is associated with. The data type is `string`.
* `type` - (Required) the type of the Extractor. The data type is `string`.
* `title` - (Required) the title of the Extractor. The data type is `string`.
* `cursor_strategy` - (Required) The data type is `string`.
* `source_field` - (Required) The data type is `string`.
* `condition_type` - (Required) the condition type of the Extractor. The data type is `string`.
* `extractor_config` - (Required) The data type is `JSON string`.
* `converters[].type` - (Required) the type of the converter. The data type is `string`.
* `converters[].config` - (Required) the configuration of the converter. The data type is `JSON string`.
* `converters` - (Optional) The data type is `[]object`. The default value is `[]`.
* `target_field` - (Optional) The data type is `string`.
* `condition_value` - (Optional) The data type is `string`.
* `order` - (Optional) The data type is `int`.

### Attributes Reference

* `extractor_id` - The id of the extractor. The data type is `string`.

## Import

`graylog_extractor` can be imported using the User `<input id>/<extractor id>`, e.g.

```console
$ terraform import graylog_extractor.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```
