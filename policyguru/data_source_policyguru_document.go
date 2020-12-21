package policyguru

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	policyGuruRest "terraform-provider-policyguru/policygururest"
	"time"
)

func dataSourcePolicyGuruDocument() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyGuruDocumentRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"overrides": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"skip_resource_constraints_for_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"actions_for_resources_at_access_level": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"write": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tagging": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"permissions_management": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"actions_for_service_without_resource_constraint_support": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"write": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tagging": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"permissions_management": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_single_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePolicyGuruDocumentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*policyGuruRest.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	policyDocumentInput := new(policyGuruRest.PolicyDocumentInput)

	// Read input and set policyDocumentInput
	if v, ok := d.GetOk("exclude_actions"); ok {
		policyDocumentInput.ExcludeActions = expandStringList(v.([]interface{}))
	}

	if v, ok := d.GetOk("actions_for_resources_at_access_level"); ok {
		actionsForResources, err := expandActionForResourcesAtAccessLevel(v.([]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		policyDocumentInput.ActionsForResources = actionsForResources
	}

	if v, ok := d.GetOk("overrides"); ok {

		overrides, err := expandOverrides(v.([]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		policyDocumentInput.Overrides = overrides
	}

	if v, ok := d.GetOk("actions_for_service_without_resource_constraint_support"); ok {

		actionForServices, err := expandActionForServicesWithoutResourceConstraints(v.([]interface{}))

		if err != nil {
			return diag.FromErr(err)
		}

		policyDocumentInput.ActionsForServices = actionForServices
	}

	// get json
	policyDocumentJsonString, err := client.GetPolicyDocumentJsonString(policyDocumentInput)
	if err != nil {
		return diag.FromErr(err)
	}

	// field to be exported
	if err := d.Set("json", policyDocumentJsonString); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
