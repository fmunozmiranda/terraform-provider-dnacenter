package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSiteAssignDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sites.

- Assigns list of devices to a site
`,

		ReadContext: dataSourceSiteAssignDeviceRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site id to which site the device to assign
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Description: `Device ip (eg: 10.104.240.64)
`,
				Type:     schema.TypeString,
				Optional: true,
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

func dataSourceSiteAssignDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vRunsync := d.Get("runsync")
	vPersistbapioutput := d.Get("persistbapioutput")
	vRunsynctimeout, okRunsynctimeout := d.GetOk("runsynctimeout")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AssignDeviceToSite")
		vvSiteID := vSiteID.(string)
		headerParams1 := dnacentersdkgo.AssignDeviceToSiteHeaderParams{}
		request1 := expandRequestSiteAssignDeviceAssignDeviceToSite(ctx, "", d)
		headerParams1.Runsync = vRunsync.(string)

		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		if okRunsynctimeout {
			headerParams1.Runsynctimeout = vRunsynctimeout.(string)
		}

		response1, _, err := client.Sites.AssignDeviceToSite(vvSiteID, request1, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AssignDeviceToSite", err,
				"Failure at AssignDeviceToSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSitesAssignDeviceToSiteItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AssignDeviceToSite response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSiteAssignDeviceAssignDeviceToSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDeviceToSite {
	request := dnacentersdkgo.RequestSitesAssignDeviceToSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestSiteAssignDeviceAssignDeviceToSiteDeviceArray(ctx, key+".device", d)
	}
	return &request
}

func expandRequestSiteAssignDeviceAssignDeviceToSiteDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice {
	request := []dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSiteAssignDeviceAssignDeviceToSiteDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSiteAssignDeviceAssignDeviceToSiteDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice {
	request := dnacentersdkgo.RequestSitesAssignDeviceToSiteDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip")))) && (ok || !reflect.DeepEqual(v, d.Get("ip"))) {
		request.IP = interfaceToString(v)
	}
	return &request
}

func flattenSitesAssignDeviceToSiteItem(item *dnacentersdkgo.ResponseSitesAssignDeviceToSite) []map[string]interface{} {
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
