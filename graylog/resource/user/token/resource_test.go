package token

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"

	"github.com/zahiar/terraform-provider-graylog/graylog/testutil"
)

const testUserId = "62aca4d45036cb55b711f258"

func TestAccUserToken(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get user tokens",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
    "tokens": [
        {
            "id": "62aca4d45036cb55b711f259",
            "last_access": "1970-01-01T00:00:00.000Z",
            "name": "toto"
        },
        {
            "id": "62acf0095036cb55b71202f7",
            "last_access": "1970-01-01T00:00:00.000Z",
            "name": "truc"
        }
    ]
}`)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a user token",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/users/" + testUserId + "/tokens/toto",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "id": "62aca4d45036cb55b711f259",
  "last_access": "1970-01-01T00:00:00.000Z",
  "name": "toto",
  "token": "10asfgsjp43de53io0cjc9ttj3r002tq26qai278ac57lhrb8p3r"
}`)),
				}, nil
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a user token",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			Path: "/api/users/" + testUserId + "/tokens/62aca4d45036cb55b711f259",
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
				}, nil
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_user_token.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_user_token" "test" {
  user_id = "` + testUserId + `"
	name    = "toto"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_user_token.test", "name", "toto"),
			resource.TestCheckResourceAttr("graylog_user_token.test", "user_id", testUserId),
			resource.TestCheckResourceAttr("graylog_user_token.test", "token", "10asfgsjp43de53io0cjc9ttj3r002tq26qai278ac57lhrb8p3r"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_user_token", Resource()),
		Steps: []resource.TestStep{
			createStep,
		},
	})
}
