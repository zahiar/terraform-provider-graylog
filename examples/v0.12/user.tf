# For Graylog V3
resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "Test Test"
  roles = [
    "Reader",
  ]
}

# For Graylog V4
# resource "graylog_user" "test" {
#   username   = "test"
#   email      = "test@example.com"
#   password   = "password"
#   first_name = "Test"
#   last_name  = "Test"
#   roles = [
#     "Reader",
#   ]
# }