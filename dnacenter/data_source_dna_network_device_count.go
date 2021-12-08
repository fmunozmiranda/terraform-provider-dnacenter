package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the interface count for the given device

- Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and
location name
`,

		ReadContext: dataSourceNetworkDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")

	method1 := []bool{okDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceInterfaceCount2")
		vvDeviceID := vDeviceID.(string)

		response1, _, err := client.Devices.GetDeviceInterfaceCount2(vvDeviceID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceInterfaceCount2", err,
				"Failure at GetDeviceInterfaceCount2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceCount2")

		response2, _, err := client.Devices.GetDeviceCount2()

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceCount2", err,
				"Failure at GetDeviceCount2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenDevicesGetDeviceCount2ItemName(response2)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCount2 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		vItemID2 := flattenDevicesGetDeviceCount2ItemID(response2)
		if err := d.Set("item_id", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCount2 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())

	}
	return diags
}

func flattenDevicesGetDeviceCount2ItemName(item *dnacentersdkgo.ResponseDevicesGetDeviceCount2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetDeviceCount2ItemID(item *dnacentersdkgo.ResponseDevicesGetDeviceCount2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
