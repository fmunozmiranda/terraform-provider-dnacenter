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
func dataSourceNfvProvisionDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Checks the provisioning detail of an ENCS device including log information.
`,

		ReadContext: dataSourceNfvProvisionDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_ip": &schema.Schema{
				Description: `Device Ip`,
				Type:        schema.TypeString,
				Optional:    true,
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

func dataSourceNfvProvisionDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vRunsync := d.Get("runsync")
	vRunsynctimeout, okRunsynctimeout := d.GetOk("runsynctimeout")
	vPersistbapioutput := d.Get("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: NfvProvisioningDetail")
		headerParams1 := dnacentersdkgo.NfvProvisioningDetailHeaderParams{}
		request1 := expandRequestNfvProvisionDetailsNfvProvisioningDetail(ctx, "", d)
		headerParams1.Runsync = vRunsync.(string)

		if okRunsynctimeout {
			headerParams1.Runsynctimeout = vRunsynctimeout.(string)
		}
		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		response1, _, err := client.SiteDesign.NfvProvisioningDetail(request1, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing NfvProvisioningDetail", err,
				"Failure at NfvProvisioningDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSiteDesignNfvProvisioningDetailItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting NfvProvisioningDetail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNfvProvisionDetailsNfvProvisioningDetail(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail {
	request := dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ip")))) {
		request.DeviceIP = interfaceToString(v)
	}
	return &request
}

func flattenSiteDesignNfvProvisioningDetailItem(item *dnacentersdkgo.ResponseSiteDesignNfvProvisioningDetail) []map[string]interface{} {
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
