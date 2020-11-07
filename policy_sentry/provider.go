package policy_sentry

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"
)

func Provider() *schema.Provider {
	return &schema.Provider{
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
