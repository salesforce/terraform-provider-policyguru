package policy_sentry

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"
	"time"
)

func dataSourcePolicySentryDocument() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicySentryDocumentRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"read" : {
			    Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
				},

			},
			"write" : {
			    Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
				},

			},
			"tagging" : {
			    Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
				},

			},
			"permissions_management" : {
			    Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
				},
			},
			"list" : {
			    Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
				},
			},
		},
	}
}

func dataSourcePolicySentryDocumentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*policySentryRest.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	policyDocumentInput := new(policySentryRest.PolicyDocumentInput)

	if v, ok := d.GetOk("read"); ok {
		policyDocumentInput.Read = expandStringList(v.([]interface{}))
	}
	if v, ok := d.GetOk("write"); ok {
		policyDocumentInput.Write = expandStringList(v.([]interface{}))
	}
	if v, ok := d.GetOk("tagging"); ok {
		policyDocumentInput.Tagging = expandStringList(v.([]interface{}))
	}
	if v, ok := d.GetOk("permissions_management"); ok {
		policyDocumentInput.PermissionsManagement = expandStringList(v.([]interface{}))
	}
	if v, ok := d.GetOk("list"); ok {
		policyDocumentInput.List = expandStringList(v.([]interface{}))
	}

	policyDocumentJsonString, err := client.GetPolicyDocumentJsonString(policyDocumentInput)
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

func expandStringList(configured []interface{}) []*string {
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, &val)
		}
	}
	return vs
}
