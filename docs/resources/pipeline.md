# Resource: graylog_pipeline

## Example Usage
```hcl
resource "graylog_pipeline" "test" {
  source = <<EOF
pipeline "test"
  stage 0 match either
end
EOF

  description = "test"
}
```

## Argument Reference

### Required Argument

* `source` - (Required) The source of the Pipeline. The data type is `string`.
* `description` - (Optional) The description of the Pipeline. The data type is `string`.

## Attributes Reference

Nothing.

## Import

`graylog_pipeline` can be imported using the Pipeline id, e.g.

```console
$ terraform import graylog_pipeline.test 5c4acaefc9e77bbbbbbbbbbb
```
