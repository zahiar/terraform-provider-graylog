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
