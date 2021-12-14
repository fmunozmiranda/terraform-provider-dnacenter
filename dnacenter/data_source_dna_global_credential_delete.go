package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGlobalCredentialDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Discovery.

- Deletes global credential for the given ID
`,

		ReadContext: dataSourceGlobalCredentialDeleteRead,
		Schema: map[string]*schema.Schema{
			"global_credential_id": &schema.Schema{
				Description: `globalCredentialId path parameter. ID of global-credential
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalCredentialDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vGlobalCredentialID, okGlobalCredentialID := d.GetOk("global_credential_id")

	method1 := []bool{okGlobalCredentialID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okGlobalCredentialID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteGlobalCredentialsByID")
		vvGlobalCredentialID := vGlobalCredentialID.(string)

		response1, _, err := client.Discovery.DeleteGlobalCredentialsByID(vvGlobalCredentialID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteGlobalCredentialsByID", err,
				"Failure at DeleteGlobalCredentialsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenDiscoveryDeleteGlobalCredentialsByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteGlobalCredentialsByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: UpdateGlobalCredentials")
		vvGlobalCredentialID := vGlobalCredentialID.(string)
		request2 := expandRequestGlobalCredentialDeleteUpdateGlobalCredentials(ctx, "", d)

		response2, _, err := client.Discovery.UpdateGlobalCredentials(vvGlobalCredentialID, request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalCredentials", err,
				"Failure at UpdateGlobalCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

	}
	return diags
}

func flattenDiscoveryDeleteGlobalCredentialsByIDItem(item *dnacentersdkgo.ResponseDiscoveryDeleteGlobalCredentialsByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}

func expandRequestGlobalCredentialDeleteUpdateGlobalCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentials {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_uuids")))) {
		request.SiteUUIDs = interfaceToSliceString(v)
	}
	return &request
}
