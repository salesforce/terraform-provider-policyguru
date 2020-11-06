package policysentry

import (
	policySentryRest "github.com/reetasingh/terraform-provider-policy-sentry/policysentry_rest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"policy-sentry-document": dataSourcePolicySentryDocument(),
		},
		ResourcesMap:   map[string]*schema.Resource{},
		ConfigureFunc: configureFunc(),
	}
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		client := policySentryRest.NewClient()
		return client, nil
	}
}
