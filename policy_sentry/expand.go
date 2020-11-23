package policy_sentry

import (
	"fmt"
	policySentryRest "terraform-provider-policy-sentry/policy_sentry_rest"
)

func expandActionForResourcesAtAccessLevel(s []interface{}) (*policySentryRest.ActionsForResourcesAtAccessLevel, error) {

	if len(s) == 0 || s[0] == nil {
		return nil, fmt.Errorf("got empty list")
	}

	data := s[0].(map[string]interface{})

	actionForResources := new(policySentryRest.ActionsForResourcesAtAccessLevel)

	if v, ok := data["read"]; ok {
		actionForResources.Read = expandStringList(v.([]interface{}))
	}
	if v, ok := data["write"]; ok {
		actionForResources.Write = expandStringList(v.([]interface{}))
	}
	if v, ok := data["tagging"]; ok {
		actionForResources.Tagging = expandStringList(v.([]interface{}))
	}
	if v, ok := data["permissions_management"]; ok {
		actionForResources.PermissionsManagement = expandStringList(v.([]interface{}))
	}
	if v, ok := data["list"]; ok {
		actionForResources.List = expandStringList(v.([]interface{}))
	}

	return actionForResources, nil

}

func expandActionForServicesWithoutResourceConstraints(s []interface{}) (*policySentryRest.ActionsForServicesWithoutResourceConstraints, error) {

	if len(s) == 0 || s[0] == nil {
		return nil, fmt.Errorf("got empty list")
	}

	data := s[0].(map[string]interface{})

	actionForServices := new(policySentryRest.ActionsForServicesWithoutResourceConstraints)

	if v, ok := data["read"]; ok {
		actionForServices.Read = expandStringList(v.([]interface{}))
	}
	if v, ok := data["write"]; ok {
		actionForServices.Write = expandStringList(v.([]interface{}))
	}
	if v, ok := data["tagging"]; ok {
		actionForServices.Tagging = expandStringList(v.([]interface{}))
	}
	if v, ok := data["permissions_management"]; ok {
		actionForServices.PermissionsManagement = expandStringList(v.([]interface{}))
	}
	if v, ok := data["list"]; ok {
		actionForServices.List = expandStringList(v.([]interface{}))
	}

	if v, ok := data["include_single_actions"]; ok {
		actionForServices.SingleActions = expandStringList(v.([]interface{}))
	}

	return actionForServices, nil
}

func expandOverrides(s []interface{}) (*policySentryRest.Overrides, error) {

	if len(s) == 0 || s[0] == nil {
		return nil, fmt.Errorf("got empty list")
	}

	data := s[0].(map[string]interface{})
	overrides := new(policySentryRest.Overrides)
	if v, ok := data["skip_resource_constraints_for_actions"]; ok {
		overrides.SkipResourceConstraints = expandStringList(v.([]interface{}))
	}
	return overrides, nil
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
