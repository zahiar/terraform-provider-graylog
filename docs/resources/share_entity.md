# Resource: graylog_share_entity

* [Source Code](https://github.com/terraform-provider-graylog/terraform-provider-graylog/blob/master/graylog/resource/authz/shares/entities/resource.go)

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
