package entities

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Delete: destroy,
		Update: update,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"entity_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"grant": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"entity_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"right": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
