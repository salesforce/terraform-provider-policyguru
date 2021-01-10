package policyguru

import (
    "testing"
    "fmt"
    )

type ClientInput struct {
	endpoint string
	expected string
}

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

    fmt.Println(input)

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