package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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
	vNetworkProfileID := d.Get("network_profile_id")
	vSiteID := d.Get("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Disassociate")
		vvNetworkProfileID := vNetworkProfileID.(string)
		vvSiteID := vSiteID.(string)

		response1, restyResp1, err := client.SiteDesign.Disassociate(vvNetworkProfileID, vvSiteID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Disassociate", err,
				"Failure at Disassociate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignDisassociateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
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
