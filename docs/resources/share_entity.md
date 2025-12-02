# Resource: graylog_share_entity

## Example Usage
```hcl
resource "graylog_share_entity" "example" {
  entity_type = "stream"
  entity_id   = graylog_stream.test.id
  grant {
    entity_type = "user"
    entity_id   = graylog_user.test.id
    right       = "view"
  }
}
```

Will update user permissions, don't mix both

## Argument Reference

* `entity_type` - (Required, Force new resource) The data type is `string`. The kind of entity to share for example `stream`.
* `entity_id` - (Required, Force new resource) The data type is `string`. The id of the entity to share.
* `grant` - (Optional) Describe a share to create. Can be specified multiple time
  * `entity_type` - (Required) The data type is `string`. The entity type of the grantee, for example `user`
  * `entity_id` - (Required) The data type is `string`. The id of the granteee
  * `right` - (Required) The data type is `string`. The right to grant. In (`manage`, `view`, `own`)

## Import

`graylog_share_entity` can be imported using the `entity_type:entity_id`, e.g.
```
$ terraform import graylog_share_entity.foo stream:123456
```
