package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns brief summary of device info such as hostname, management IP address for the given device Id
`,

		ReadContext: dataSourceNetworkDeviceSummaryRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"role_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceSummary")
		vvID := vID.(string)

		response1, _, err := client.Devices.GetDeviceSummary(vvID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceSummary", err,
				"Failure at GetDeviceSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenDevicesGetDeviceSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceSummary response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())

	}
	return diags
}

func flattenDevicesGetDeviceSummaryItem(item *dnacentersdkgo.ResponseDevicesGetDeviceSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["role"] = item.Role
	respItem["role_source"] = item.RoleSource
	return []map[string]interface{}{
		respItem,
	}
}
