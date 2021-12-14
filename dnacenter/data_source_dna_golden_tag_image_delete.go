package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGoldenTagImageDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Software Image Management (SWIM).

- Remove golden tag. Set siteId as -1 for Global site.
`,

		ReadContext: dataSourceGoldenTagImageDeleteRead,
		Schema: map[string]*schema.Schema{
			"device_family_identifier": &schema.Schema{
				Description: `deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"device_role": &schema.Schema{
				Description: `deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": &schema.Schema{
				Description: `imageId path parameter. Image Id in uuid format.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.
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
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5 
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGoldenTagImageDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vDeviceFamilyIDentifier, okDeviceFamilyIDentifier := d.GetOk("device_family_identifier")
	vDeviceRole, okDeviceRole := d.GetOk("device_role")
	vImageID, okImageID := d.GetOk("image_id")

	method1 := []bool{okSiteID, okDeviceFamilyIDentifier, okDeviceRole, okImageID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okSiteID, okDeviceFamilyIDentifier, okDeviceRole, okImageID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveGoldenTagForImage")
		vvSiteID := vSiteID.(string)
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier.(string)
		vvDeviceRole := vDeviceRole.(string)
		vvImageID := vImageID.(string)

		response1, _, err := client.SoftwareImageManagementSwim.RemoveGoldenTagForImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveGoldenTagForImage", err,
				"Failure at RemoveGoldenTagForImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSoftwareImageManagementSwimRemoveGoldenTagForImageItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveGoldenTagForImage response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetGoldenTagStatusOfAnImage")
		vvSiteID := vSiteID.(string)
		vvDeviceFamilyIDentifier := vDeviceFamilyIDentifier.(string)
		vvDeviceRole := vDeviceRole.(string)
		vvImageID := vImageID.(string)

		response2, _, err := client.SoftwareImageManagementSwim.GetGoldenTagStatusOfAnImage(vvSiteID, vvDeviceFamilyIDentifier, vvDeviceRole, vvImageID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGoldenTagStatusOfAnImage", err,
				"Failure at GetGoldenTagStatusOfAnImage, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

	}
	return diags
}

func flattenSoftwareImageManagementSwimRemoveGoldenTagForImageItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
