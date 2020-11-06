package policysentry_rest

import (
	"encoding/json"
	"context"
	"net/http"
    "bytes"
    "io/ioutil"
    "strconv"
    "time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type PolicyDocument struct {
	Statement []struct {
		Action   []string `json:"Action"`
		Effect   string   `json:"Effect"`
		Resource []string `json:"Resource"`
		Sid      string   `json:"Sid"`
	} `json:"Statement"`
	Version string `json:"Version"`
}