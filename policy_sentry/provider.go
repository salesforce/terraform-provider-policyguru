package policy_sentry

import (
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["endpoint"],
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"policy-sentry_document": dataSourcePolicySentryDocument(),
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: configureFunc(),
	}
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {

	return func(d *schema.ResourceData) (interface{}, error) {
		client := policySentryRest.NewClient()
		return client, nil
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"endpoint": "Use this to override the default service endpoint URL",
	}
}
