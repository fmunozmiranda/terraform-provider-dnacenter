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
func dataSourceWirelessProvisionAccessPoint() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Access Point Provision and ReProvision
`,

		ReadContext: dataSourceWirelessProvisionAccessPointRead,
		Schema: map[string]*schema.Schema{
			"custom_ap_group_name": &schema.Schema{
				Description: `Custom AP group name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_flex_group_name": &schema.Schema{
				Description: `["Custom flex group name"]
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"device_name": &schema.Schema{
				Description: `Device name
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
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_url": &schema.Schema{
							Description: `Execution Url`,
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
			"rf_profile": &schema.Schema{
				Description: `Radio frequency profile name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `Site name hierarchy(ex: Global/...)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_name_hierarchy": &schema.Schema{
				Description: `Site name hierarchy(ex: Global/...)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `ApWirelessConfiguration
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceWirelessProvisionAccessPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ApProvision")
		headerParams1 := dnacentersdkgo.ApProvisionHeaderParams{}
		request1 := expandRequestWirelessProvisionAccessPointApProvision(ctx, "", d)
		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, _, err := client.Wireless.ApProvision(request1, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ApProvision", err,
				"Failure at ApProvision, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenWirelessApProvisionItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ApProvision response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestWirelessProvisionAccessPointApProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessApProvision {
	request := dnacentersdkgo.RequestWirelessApProvision{}
	if v := expandRequestWirelessProvisionAccessPointApProvisionArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionAccessPointApProvisionArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessApProvision {
	request := []dnacentersdkgo.RequestItemWirelessApProvision{}
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
		i := expandRequestItemWirelessProvisionAccessPointApProvision(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemWirelessProvisionAccessPointApProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessApProvision {
	request := dnacentersdkgo.RequestItemWirelessApProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get("rf_profile"))) {
		request.RfProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get("site_id"))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get("device_name"))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_ap_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_ap_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get("custom_ap_group_name"))) {
		request.CustomApGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_flex_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_flex_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get("custom_flex_group_name"))) {
		request.CustomFlexGroupName = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get("type"))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get("site_name_hierarchy"))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	return &request
}

func flattenWirelessApProvisionItems(items *dnacentersdkgo.ResponseWirelessApProvision) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["execution_id"] = item.ExecutionID
		respItem["execution_url"] = item.ExecutionURL
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}
