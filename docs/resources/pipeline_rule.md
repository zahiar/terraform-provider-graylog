# Resource: graylog_pipeline_rule

## Example Usage
```hcl
resource "graylog_pipeline_rule" "test" {
  source = <<EOF
rule "test"
when
    to_long($message.status) < 500
then
    set_field("status_01", 1);
end
EOF

  description = "test"
}
```

## Argument Reference

* `source` - (Required) The source of the Pipeline Rule. The data type is `string`.
* `description` - (Optional) description of the Pipeline Rule. The data type is `string`.

## Attributes Reference

Nothing.

## Import

`graylog_pipeline_rule` can be imported using the Pipeline Rule id, e.g.

```console
$ terraform import graylog_pipeline_rule.test 5c4acaefc9e77bbbbbbbbbbb
```
