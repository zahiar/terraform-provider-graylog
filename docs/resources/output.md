# Resource: graylog_output

## Example Usage
```hcl
resource "graylog_output" "stdout" {
  title = "stdout"
  type  = "org.graylog2.outputs.LoggingOutput"

  configuration = jsonencode({
    prefix = "Writing message: "
  })
}
```
```hcl
resource "graylog_output" "gelf" {
  title = "gelf"
  type  = "org.graylog2.outputs.GelfOutput"

  configuration = jsonencode({
    "hostname" : "localhost",
    "protocol" : "TCP",
    "connect_timeout" : 1000,
    "reconnect_delay" : 500,
    "queue_size" : 512,
    "port" : 12201,
    "max_inflight_sends" : 512,
    "tcp_no_delay" : false,
    "tcp_keep_alive" : false,
    "tls_trust_cert_chain" : "",
    "tls_verification_enabled" : false
  })
}
```
```hcl
resource "graylog_output" "slack" {
  title = "slack"
  type  = "org.graylog2.plugins.slack.output.SlackMessageOutput"

  configuration = jsonencode({
    "icon_url" : "",
    "graylog2_url" : "",
    "link_names" : true,
    "color" : "#FF0000",
    "webhook_url" : "http://example.com",
    "icon_emoji" : "",
    "user_name" : "Graylog",
    "proxy_address" : "",
    "channel" : "#channel",
    "custom_message" : "message",
    "notify_channel" : false,
    "short_mode" : false,
    "add_details" : true
  })
}
```

## Argument Reference

* `title` - (Required) The title of the Output. The data type is `string`.
* `type` - (Required) The type of the Output. The data type is `string`.
* `configuration` - (Required) The configuration of the Output. The data type is `JSON string`.

`configuration` is a JSON string.
The format of `configuration` depends on the output type.
Please see the example above.
Using the [Graylog's API browser](https://docs.graylog.org/en/3.1/pages/configuration/rest_api.html) you can check the format of `configuration`.

## Attributes Reference

None.

## Import

`graylog_output` can be imported using the Output id, e.g.

```console
$ terraform import graylog_output.test 5c4acaefc9e77bbbbbbbbbbb
```
