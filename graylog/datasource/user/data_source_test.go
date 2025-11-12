package user

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"

	"github.com/zahiar/terraform-provider-graylog/graylog/testutil"
)

func TestAccUser(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getUserRoute := flute.Route{
		Name: "get a user",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/users/test",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "id": "12a4b5678901234cde567f8g",
  "username": "test",
  "email": "test@example.com",
  "first_name": null,
  "last_name": null,
  "full_name": "Test User",
  "permissions": [ "users:edit:test", "streams:read:123456789012345678901234" ],
  "grn_permissions": [ "entity:own:grn::::stream:11cf4cbacadd9c316802221b" ],
  "preferences": { "updateUnfocussed": true, "enableSmartSearch": true },
  "timezone": null,
  "session_timeout_ms": 28800000,
  "external": true,
  "roles": [ "Admin", "Reader", "Views Manager" ],
  "read_only": false,
  "session_active": true,
  "last_activity": "2025-11-12T00:00:00.000+0000",
  "client_address": "10.10.10.1",
  "account_status": "enabled",
  "auth_service_enabled": true,
  "service_account": false
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_user.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getUserRoute)
		},
		Config: `
data "graylog_user" "test" {
  username = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_user.test", "username", "test"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "email", "test@example.com"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "full_name", "Test User"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "session_timeout_ms", "28800000"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "external", "true"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "read_only", "false"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "session_active", "true"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "last_activity", "2025-11-12T00:00:00.000+0000"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "client_address", "10.10.10.1"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_user", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
