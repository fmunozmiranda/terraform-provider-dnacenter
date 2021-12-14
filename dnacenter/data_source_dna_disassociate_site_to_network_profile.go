package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceDisassociateSiteToNetworkProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Site Design.

- Disassociate a Site from a Network Profile
`,

		ReadContext: dataSourceDisassociateSiteToNetworkProfileRead,
		Schema: map[string]*schema.Schema{
			"network_profile_id": &schema.Schema{
				Description: `networkProfileId path parameter. Network-Profile Id to be associated
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site Id to be associated
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
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDisassociateSiteToNetworkProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkProfileID, okNetworkProfileID := d.GetOk("network_profile_id")
	vSiteID, okSiteID := d.GetOk("site_id")

	method1 := []bool{okNetworkProfileID, okSiteID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okNetworkProfileID, okSiteID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Associate")
		vvNetworkProfileID := vNetworkProfileID.(string)
		vvSiteID := vSiteID.(string)

		response1, _, err := client.SiteDesign.Associate(vvNetworkProfileID, vvSiteID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Associate", err,
				"Failure at Associate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: Disassociate")
		vvNetworkProfileID := vNetworkProfileID.(string)
		vvSiteID := vSiteID.(string)

		response2, _, err := client.SiteDesign.Disassociate(vvNetworkProfileID, vvSiteID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Disassociate", err,
				"Failure at Disassociate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSiteDesignDisassociateItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Disassociate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignDisassociateItem(item *dnacentersdkgo.ResponseSiteDesignDisassociateResponse) []map[string]interface{} {
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
