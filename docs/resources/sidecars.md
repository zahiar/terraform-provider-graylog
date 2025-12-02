# Resource: graylog_sidecars

Manages to assign Sidecars's configuration to Sidecars.
Due to the Graylog API's restriction, we have to manage all assignments by one Terraform resource,
which means we shouldn't use this resource more than once.

## Example Usage
```hcl
resource "graylog_sidecars" "all" { 
  sidecars {
    node_id = data.graylog_sidecar.test.id
    assignments {
      collector_id     = graylog_sidecar_collector.test.id
      configuration_id = graylog_sidecar_configuration.test.id
    }
  }
}
```

## Argument Reference

* `sidecars` - (Required) The data type is `[]object (set)`
* `sidecars[].node_id` - (Required) The data type is `string`
* `sidecars[].assignments` - (Required) The data type is `[]object (set)`
* `sidecars[].assignments[].collector_id` - (Required) The data type is `string`
* `sidecars[].assignments[].configuration_id` - (Required) The data type is `string`

## Attributes Reference

None.

## Import

Unlike other resources, the given ID is ignored so please specify any string as ID.

e.g.

```console
$ terraform import graylog_sidecars.all all
```
