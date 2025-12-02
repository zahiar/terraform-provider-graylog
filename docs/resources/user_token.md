# Resource: graylog_user_token

## Example Usage
```hcl
resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "test test"
  roles = [
    "Reader",
  ]
}

resource "graylog_user_token" "token" {
  user_id = graylog_user.test.id
  name    = "mytoken"
}
```

## Argument Reference

* `user_id` - (Required, Forces new resource) The id of the user. The data type is string
* `name` - (Required, Forces new resource) The name of the token. The data type is string

## Attributes Reference

* `token` - The generated token. The data type is `string` 

## Import

`graylog_user_token` can be imported using the `user_id:id`, e.g.
```
$ terraform import graylog_user_token.foo 123456:12345678
```
Even it's not very useful since token won't be set due to graylog API
