package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zahiar/terraform-provider-graylog/graylog/resource/user"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			d.SetId(d.Get("username").(string))
			return user.Read(d, m)
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},

			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_timeout_ms": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"session_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_activity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
