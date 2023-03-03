resource "graylog_user" "test" {
  username   = "test"
  email      = "test@example.com"
  password   = "password"
  first_name = "Test"
  last_name  = "Test"
  roles = [
    "Reader",
  ]
}
