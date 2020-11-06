package policy_sentry

import (
	"context"
    "strconv"
    "time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	policySentryRest "github.com/reetasingh/terraform-provider-policy-sentry/policysentry_rest"
)


func dataSourcePolicySentryDocument() *schema.Resource {
	setOfString := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	return &schema.Resource{
		ReadContext: dataSourcePolicySentryDocumentRead,

		Schema: map[string]*schema.Schema{
			"statement": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"effect": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Allow",
							ValidateFunc: validation.StringInSlice([]string{"Allow", "Deny"}, false),
						},
						"actions":        setOfString,
						"resources":      setOfString,
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2012-10-17",
				ValidateFunc: validation.StringInSlice([]string{
					"2008-10-17",
					"2012-10-17",
				}, false),
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePolicySentryDocumentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {


  client := m.(*policySentryRest.Client)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  policyDocumentJsonString, err := client.GetPolicyDocumentJsonString()
  if err != nil {
		return nil, err
  }

 if err := d.Set("json", policyDocumentJsonString); err != nil {
    return diag.FromErr(err)
  }
  // always run
  d.SetId(strconv.FormatInt(time.Now().Unix(), 10))


  return diags
}