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
func dataSourceBusinessSdaWirelessControllerCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Fabric Wireless.

- Add WLC to Fabric Domain
`,

		ReadContext: dataSourceBusinessSdaWirelessControllerCreateRead,
		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Description: `EWLC Device Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Status of the job for wireless state change in fabric domain
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_status_url": &schema.Schema{
							Description: `executionStatusURL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"site_name_hierarchy": &schema.Schema{
				Description: `Site Name Hierarchy
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceBusinessSdaWirelessControllerCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceIPAddress, okDeviceIPAddress := d.GetOk("device_ipaddress")

	method1 := []bool{okDeviceIPAddress}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveWLCFromFabricDomain")
		queryParams1 := dnacentersdkgo.RemoveWLCFromFabricDomainQueryParams{}

		if okDeviceIPAddress {
			queryParams1.DeviceIPAddress = vDeviceIPAddress.(string)
		}

		response1, _, err := client.FabricWireless.RemoveWLCFromFabricDomain(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveWLCFromFabricDomain", err,
				"Failure at RemoveWLCFromFabricDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: AddWLCToFabricDomain")
		request2 := expandRequestBusinessSdaWirelessControllerCreateAddWLCToFabricDomain(ctx, "", d)

		response2, _, err := client.FabricWireless.AddWLCToFabricDomain(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddWLCToFabricDomain", err,
				"Failure at AddWLCToFabricDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItems2 := flattenFabricWirelessAddWLCToFabricDomainItems(response2)
		if err := d.Set("items", vItems2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AddWLCToFabricDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBusinessSdaWirelessControllerCreateAddWLCToFabricDomain(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessAddWLCToFabricDomain {
	request := dnacentersdkgo.RequestFabricWirelessAddWLCToFabricDomain{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	return &request
}

func flattenFabricWirelessAddWLCToFabricDomainItems(items *dnacentersdkgo.ResponseFabricWirelessAddWLCToFabricDomain) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["execution_id"] = item.ExecutionID
		respItem["execution_status_url"] = item.ExecutionStatusURL
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}
