package hashicups

import (
	policySentryRest "github.com/reetasingh/policysentry_rest"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"policy_sentry_document": dataSourcePolicySentryDocument(),
		},
		ConfigureFunc: configureFunc(),
	}
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		client := policySentryRest.NewClient()
		return client, nil
	}
}
