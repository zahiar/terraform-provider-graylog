package entities

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
const testStreamId = "628b9b556783e72e8b4782b1"

func TestAccShareEntity(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get entity shares",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
    "active_shares": [
        {
            "capability": "view",
            "grant": "grn::::grant:62ac81615036cb55b711e1f4",
            "grantee": "grn::::user:` + testUserId + `"
        }
    ],
    "available_capabilities": [
        {
            "id": "view",
            "title": "Viewer"
        },
        {
            "id": "manage",
            "title": "Manager"
        },
        {
            "id": "own",
            "title": "Owner"
        }
    ],
    "available_grantees": [
        {
            "id": "grn::::user:628b9b286783e72e8b4781b0",
            "title": "Sidecar System User (built-in)",
            "type": "user"
        },
        {
            "id": "grn::::user:` + testUserId + `",
            "title": "Test Test",
            "type": "user"
        },
        {
            "id": "grn::::builtin-team:everyone",
            "title": "Everyone",
            "type": "global"
        }
    ],
    "entity": "grn::::stream:` + testStreamId + `",
    "missing_permissions_on_dependencies": {},
    "selected_grantee_capabilities": {
        "grn::::user:` + testUserId + `": "view"
    },
    "sharing_user": "grn::::user:local:admin",
    "validation_result": {
        "error_context": {},
        "errors": {},
        "failed": false
    }
}`)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a grant",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/authz/shares/entities/grn::::user:" + testStreamId + "/",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
    "active_shares": [
        {
            "capability": "view",
            "grant": "grn::::grant:62ac81615036cb55b711e1f4",
            "grantee": "grn::::user:` + testUserId + `"
        }
    ],
    "available_capabilities": [
        {
            "id": "view",
            "title": "Viewer"
        },
        {
            "id": "manage",
            "title": "Manager"
        },
        {
            "id": "own",
            "title": "Owner"
        }
    ],
    "available_grantees": [
        {
            "id": "grn::::user:628b9b286783e72e8b4781b0",
            "title": "Sidecar System User (built-in)",
            "type": "user"
        },
        {
            "id": "grn::::user:` + testUserId + `",
            "title": "Test Test",
            "type": "user"
        },
        {
            "id": "grn::::builtin-team:everyone",
            "title": "Everyone",
            "type": "global"
        }
    ],
    "entity": "grn::::stream:` + testStreamId + `",
    "missing_permissions_on_dependencies": {},
    "selected_grantee_capabilities": {
        "grn::::user:` + testUserId + `": "view"
    },
    "sharing_user": "grn::::user:local:admin",
    "validation_result": {
        "error_context": {},
        "errors": {},
        "failed": false
    }
}`)),
				}, nil
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_share_entity.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_share_entity" "test" {
	entity_type = "stream"
	entity_id = "` + testStreamId + `"
	grant {
		entity_type = "user"
		entity_id = "` + testUserId + `"
		right = "view"
	}
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_share_entity.test", "id", "stream:"+testStreamId),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_share_entity", Resource()),
		Steps: []resource.TestStep{
			createStep,
		},
	})
}
