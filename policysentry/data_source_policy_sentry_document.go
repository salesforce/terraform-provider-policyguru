package policysentry

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


func dataSourcePolicySentryDocument() *schema.Resource {
	setOfString := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	return &schema.Resource{
		ReadContext: dataSourcePolicySentryDocumentRead,

		Schema: map[string]*schema.Schema{
			"statement": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"effect": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Allow",
							ValidateFunc: validation.StringInSlice([]string{"Allow", "Deny"}, false),
						},
						"actions":        setOfString,
						"resources":      setOfString,
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2012-10-17",
				ValidateFunc: validation.StringInSlice([]string{
					"2008-10-17",
					"2012-10-17",
				}, false),
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePolicySentryDocumentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  requestBody, err := json.Marshal(map[string] interface{} {
    "mode": "crud",
    "read": []string{"arn:aws:s3:::example-org-s3-access-logs"} })

  if err != nil {
    return diag.FromErr(err)
  }

  r, err := http.Post("https://zeok878mvj.execute-api.us-east-1.amazonaws.com/dev/write", "application/json", bytes.NewBuffer(requestBody))
  if err != nil {
    return diag.FromErr(err)
  }

  if r.StatusCode != 200 {
    bodyBytes, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return diag.FromErr(err)
    }
    bodyString := string(bodyBytes)
    diags = append(diags, diag.Diagnostic{
        Severity: diag.Error,
        Summary:  "got non 200",
        Detail:   bodyString,
    })
    return diags
  }

  body, err := ioutil.ReadAll(r.Body)

//  policy := PolicyDocument{}

//  err = json.Unmarshal(body, &policy)
//	if err != nil {
//		return diag.FromErr(err)
//  }


//  b, err := json.Marshal(policy)
 //   if err != nil {
//       return diag.FromErr(err)
 //  }

 if err := d.Set("json", string(body)); err != nil {
    return diag.FromErr(err)
  }
  // always run
  d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

  defer r.Body.Close()

  return diags
}