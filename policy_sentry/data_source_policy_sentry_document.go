package policy_sentry

import (
	"context"
    "strconv"
    "time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"
)


func dataSourcePolicySentryDocument() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicySentryDocumentRead,

		Schema: map[string]*schema.Schema{
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
		return diag.FromErr(err)
  }

 if err := d.Set("json", policyDocumentJsonString); err != nil {
    return diag.FromErr(err)
  }
  // always run
  d.SetId(strconv.FormatInt(time.Now().Unix(), 10))


  return diags
}