package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceWirelessProvisionSSIDDeleteReprovision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Wireless.

- Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center
`,

		ReadContext: dataSourceWirelessProvisionSSIDDeleteReprovisionRead,
		Schema: map[string]*schema.Schema{
			"managed_aplocations": &schema.Schema{
				Description: `managedAPLocations path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"ssid_name": &schema.Schema{
				Description: `ssidName path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessProvisionSSIDDeleteReprovisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSSIDName := d.Get("ssid_name")
	vManagedApLocations := d.Get("managed_aplocations")
	vPersistbapioutput := d.Get("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteSSIDAndProvisionItToDevices")
		vvSSIDName := vSSIDName.(string)
		vvManagedApLocations := vManagedApLocations.(string)
		headerParams1 := dnacentersdkgo.DeleteSSIDAndProvisionItToDevicesHeaderParams{}

		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		response1, _, err := client.Wireless.DeleteSSIDAndProvisionItToDevices(vvSSIDName, vvManagedApLocations, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteSSIDAndProvisionItToDevices", err,
				"Failure at DeleteSSIDAndProvisionItToDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenWirelessDeleteSSIDAndProvisionItToDevicesItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteSSIDAndProvisionItToDevices response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessDeleteSSIDAndProvisionItToDevicesItem(item *dnacentersdkgo.ResponseWirelessDeleteSSIDAndProvisionItToDevices) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
