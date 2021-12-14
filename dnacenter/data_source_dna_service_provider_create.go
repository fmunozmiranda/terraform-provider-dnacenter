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
func dataSourceServiceProviderCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to create service provider profile(QOS).
`,

		ReadContext: dataSourceServiceProviderCreateRead,
		Schema: map[string]*schema.Schema{
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
			"qos": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"model": &schema.Schema{
							Description: `Model`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"profile_name": &schema.Schema{
							Description: `Profile Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"wan_provider": &schema.Schema{
							Description: `Wan Provider`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServiceProviderCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetServiceProviderDetails")

		response1, _, err := client.NetworkSettings.GetServiceProviderDetails()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetServiceProviderDetails", err,
				"Failure at GetServiceProviderDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: CreateSpProfile")
		request2 := expandRequestServiceProviderCreateCreateSpProfile(ctx, "", d)

		response2, _, err := client.NetworkSettings.CreateSpProfile(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSpProfile", err,
				"Failure at CreateSpProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkSettingsCreateSpProfileItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSpProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: UpdateSpProfile")
		request3 := expandRequestServiceProviderCreateUpdateSpProfile(ctx, "", d)

		response3, _, err := client.NetworkSettings.UpdateSpProfile(request3)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSpProfile", err,
				"Failure at UpdateSpProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

	}
	return diags
}

func expandRequestServiceProviderCreateCreateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfile{}
	request.Settings = expandRequestServiceProviderCreateCreateSpProfileSettings(ctx, key, d)
	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get("qos"))) {
		request.Qos = expandRequestServiceProviderCreateCreateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}
	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
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
		i := expandRequestServiceProviderCreateCreateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get("profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get("model"))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get("wan_provider"))) {
		request.WanProvider = interfaceToString(v)
	}
	return &request
}

func flattenNetworkSettingsCreateSpProfileItem(item *dnacentersdkgo.ResponseNetworkSettingsCreateSpProfile) []map[string]interface{} {
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

func expandRequestServiceProviderCreateUpdateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile{}
	request.Settings = expandRequestServiceProviderCreateUpdateSpProfileSettings(ctx, key, d)
	return &request
}

func expandRequestServiceProviderCreateUpdateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get("qos"))) {
		request.Qos = expandRequestServiceProviderCreateUpdateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}
	return &request
}

func expandRequestServiceProviderCreateUpdateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
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
		i := expandRequestServiceProviderCreateUpdateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestServiceProviderCreateUpdateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get("profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get("model"))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get("wan_provider"))) {
		request.WanProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".old_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".old_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get("old_profile_name"))) {
		request.OldProfileName = interfaceToString(v)
	}
	return &request
}
