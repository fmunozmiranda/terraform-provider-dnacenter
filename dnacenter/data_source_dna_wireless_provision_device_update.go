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
func dataSourceWirelessProvisionDeviceUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Wireless.

- Updates wireless provisioning
`,

		ReadContext: dataSourceWirelessProvisionDeviceUpdateRead,
		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Description: `Device Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dynamic_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_gateway": &schema.Schema{
							Description: `Interface Gateway`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"interface_ipaddress": &schema.Schema{
							Description: `Interface IPAddress`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"interface_netmask_in_cid_r": &schema.Schema{
							Description: `Interface Netmask In CIDR`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"lag_or_port_number": &schema.Schema{
							Description: `Lag Or Port Number`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
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
						"execution_url": &schema.Schema{
							Description: `Execution Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"provisioning_tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failed": &schema.Schema{
										Description: `Failed`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"success": &schema.Schema{
										Description: `Success`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"managed_aplocations": &schema.Schema{
				Description: `Managed APLocations`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceWirelessProvisionDeviceUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	method1 := []bool{okPersistbapioutput}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ProvisionUpdate")
		headerParams1 := dnacentersdkgo.ProvisionUpdateHeaderParams{}
		request1 := expandRequestWirelessProvisionDeviceUpdateProvisionUpdate(ctx, "", d)
		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, _, err := client.Wireless.ProvisionUpdate(request1, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ProvisionUpdate", err,
				"Failure at ProvisionUpdate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenWirelessProvisionUpdateItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ProvisionUpdate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: Provision")
		request2 := expandRequestWirelessProvisionDeviceUpdateProvision(ctx, "", d)

		response2, _, err := client.Wireless.Provision(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Provision", err,
				"Failure at Provision, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

	}
	return diags
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessProvisionUpdate {
	request := dnacentersdkgo.RequestWirelessProvisionUpdate{}
	if v := expandRequestWirelessProvisionDeviceUpdateProvisionUpdateArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionUpdate {
	request := []dnacentersdkgo.RequestItemWirelessProvisionUpdate{}
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
		i := expandRequestItemWirelessProvisionDeviceUpdateProvisionUpdate(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemWirelessProvisionDeviceUpdateProvisionUpdate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionUpdate {
	request := dnacentersdkgo.RequestItemWirelessProvisionUpdate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get("device_name"))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_aplocations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_aplocations")))) && (ok || !reflect.DeepEqual(v, d.Get("managed_aplocations"))) {
		request.ManagedApLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get("dynamic_interfaces"))) {
		request.DynamicInterfaces = expandRequestWirelessProvisionDeviceUpdateProvisionUpdateDynamicInterfacesArray(ctx, key+".dynamic_interfaces", d)
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateDynamicInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces {
	request := []dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces{}
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
		i := expandRequestWirelessProvisionDeviceUpdateProvisionUpdateDynamicInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionUpdateDynamicInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces {
	request := dnacentersdkgo.RequestItemWirelessProvisionUpdateDynamicInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_ipaddress"))) {
		request.InterfaceIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_netmask_in_cid_r")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_netmask_in_cid_r"))) {
		request.InterfaceNetmaskInCIDR = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_gateway"))) {
		request.InterfaceGateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lag_or_port_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lag_or_port_number")))) && (ok || !reflect.DeepEqual(v, d.Get("lag_or_port_number"))) {
		request.LagOrPortNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get("vlan_id"))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_name"))) {
		request.InterfaceName = interfaceToString(v)
	}
	return &request
}

func flattenWirelessProvisionUpdateItem(item *dnacentersdkgo.ResponseWirelessProvisionUpdate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_url"] = item.ExecutionURL
	respItem["provisioning_tasks"] = flattenWirelessProvisionUpdateItemProvisioningTasks(item.ProvisioningTasks)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessProvisionUpdateItemProvisioningTasks(item *dnacentersdkgo.ResponseWirelessProvisionUpdateProvisioningTasks) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = item.Success
	respItem["failed"] = item.Failed

	return []map[string]interface{}{
		respItem,
	}

}

func expandRequestWirelessProvisionDeviceUpdateProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessProvision {
	request := dnacentersdkgo.RequestWirelessProvision{}
	if v := expandRequestWirelessProvisionDeviceUpdateProvisionArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvision {
	request := []dnacentersdkgo.RequestItemWirelessProvision{}
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
		i := expandRequestItemWirelessProvisionDeviceUpdateProvision(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemWirelessProvisionDeviceUpdateProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvision {
	request := dnacentersdkgo.RequestItemWirelessProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get("device_name"))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get("site"))) {
		request.Site = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_aplocations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_aplocations")))) && (ok || !reflect.DeepEqual(v, d.Get("managed_aplocations"))) {
		request.ManagedApLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get("dynamic_interfaces"))) {
		request.DynamicInterfaces = expandRequestWirelessProvisionDeviceUpdateProvisionDynamicInterfacesArray(ctx, key+".dynamic_interfaces", d)
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionDynamicInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces {
	request := []dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces{}
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
		i := expandRequestWirelessProvisionDeviceUpdateProvisionDynamicInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionDeviceUpdateProvisionDynamicInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces {
	request := dnacentersdkgo.RequestItemWirelessProvisionDynamicInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_ipaddress"))) {
		request.InterfaceIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_netmask_in_cid_r")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_netmask_in_cid_r"))) {
		request.InterfaceNetmaskInCIDR = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_gateway"))) {
		request.InterfaceGateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lag_or_port_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lag_or_port_number")))) && (ok || !reflect.DeepEqual(v, d.Get("lag_or_port_number"))) {
		request.LagOrPortNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get("vlan_id"))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get("interface_name"))) {
		request.InterfaceName = interfaceToString(v)
	}
	return &request
}
