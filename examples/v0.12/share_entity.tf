resource "graylog_stream" "test" {
  title         = "test"
  index_set_id  = data.graylog_index_set.default.id
  disabled      = true
  matching_type = "AND"
  description   = "test"
}

data "graylog_stream" "all_messages" {
  title = "All messages"
}

resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "test test"
  roles = [
    "Reader",
  ]
}

resource "graylog_share_entity" "example" {
  entity_type = "stream"
  entity_id   = graylog_stream.test.id
  grant {
    entity_type = "user"
    entity_id   = graylog_user.test.id
    right       = "view"
  }
}
