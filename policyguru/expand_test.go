package policyguru

import (
	"testing"
)

func assertNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Errorf("var a %v is nil", a)
	}
}

func TestExpandActionForResourcesAtAccessLevel(t *testing.T) {

	inputMap := make(map[string]interface{})

	var data []interface{}
	data = append(data, "s3")
	inputMap["read"] = data
	inputMap["write"] = data
	inputMap["tagging"] = data
	inputMap["permissions_management"] = data
	inputMap["list"] = data

	input := make([]interface{}, 0, 1)
	input = append(input, inputMap)

	actionForResources, err := expandActionForResourcesAtAccessLevel(input)
	if err != nil {
		t.Fatal(err)
	}
	assertNotNil(t, actionForResources.Read)
	assertNotNil(t, actionForResources.Write)
	assertNotNil(t, actionForResources.List)
	assertNotNil(t, actionForResources.Tagging)
	assertNotNil(t, actionForResources.PermissionsManagement)

}

func TestExpandActionForServicesWithoutResourceConstraints(t *testing.T) {

	inputMap := make(map[string]interface{})

	var data []interface{}
	data = append(data, "s3")
	inputMap["read"] = data
	inputMap["write"] = data
	inputMap["tagging"] = data
	inputMap["permissions_management"] = data
	inputMap["list"] = data

	input := make([]interface{}, 0, 1)
	input = append(input, inputMap)

	actionForServices, err := expandActionForServicesWithoutResourceConstraints(input)
	if err != nil {
		t.Fatal(err)
	}
	assertNotNil(t, actionForServices.Read)
	assertNotNil(t, actionForServices.Write)
	assertNotNil(t, actionForServices.List)
	assertNotNil(t, actionForServices.Tagging)
	assertNotNil(t, actionForServices.PermissionsManagement)

}

func TestExpandOverrides(t *testing.T) {

	inputMap := make(map[string]interface{})

	var data []interface{}
	data = append(data, "s3")
	inputMap["skip_resource_constraints_for_actions"] = data

	input := make([]interface{}, 0, 1)
	input = append(input, inputMap)

	overrides, err := expandOverrides(input)
	if err != nil {
		t.Fatal(err)
	}
	assertNotNil(t, overrides.SkipResourceConstraints)
}

func TestExpandStringList(t *testing.T) {

	input := make([]interface{}, 0, 1)
	input = append(input, "abc")
	result := expandStringList(input)

	assertNotNil(t, result)

}
