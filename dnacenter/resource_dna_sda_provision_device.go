package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaProvisionDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Delete provisioned Wired Device

- Re-Provision Wired Device

- Provision Wired Device
`,

		CreateContext: resourceSdaProvisionDeviceCreate,
		ReadContext:   resourceSdaProvisionDeviceRead,
		UpdateContext: resourceSdaProvisionDeviceUpdate,
		DeleteContext: resourceSdaProvisionDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'deviceManagementIpAddress': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Management Ip Address of the device to be provisioned\n'}, 'siteNameHierarchy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Name Hierarchy for device location(only building / floor level)\n'}}}}}, 'metadata': {'item': {'operation_id': [['ProvisionWiredDevice', 'ReProvisionWiredDevice']], 'new_flat_structure': [[{'RequestSdaProvisionWiredDevice': {'type': 'obj', 'data': [{'name': 'deviceManagementIpAddress', 'description': 'Management Ip Address of the device to be provisioned\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'siteNameHierarchy', 'description': 'Site Name Hierarchy for device location(only building / floor level)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestSdaReProvisionWiredDevice': {'type': 'obj', 'data': [{'name': 'deviceManagementIpAddress', 'description': 'Management Ip Address of the device to be re-provisioned\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'siteNameHierarchy', 'description': 'Site Name Hierarchy for device location(only building / floor level)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestSdaProvisionWiredDevice', 'RequestSdaReProvisionWiredDevice']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the device to be provisioned
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy for device location(only building / floor level)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaProvisionDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaProvisionDeviceProvisionWiredDevice(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.ProvisionWiredDevice(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ProvisionWiredDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ProvisionWiredDevice", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSdaProvisionDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress, okDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetProvisionedWiredDevice")
		queryParams1 := dnacentersdkgo.GetProvisionedWiredDeviceQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetProvisionedWiredDevice(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetProvisionedWiredDevice", err,
				"Failure at GetProvisionedWiredDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceSdaProvisionDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress, okDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestSdaProvisionDeviceReProvisionWiredDevice(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.ReProvisionWiredDevice(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ReProvisionWiredDevice", err, restyResp1.String(),
					"Failure at ReProvisionWiredDevice, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReProvisionWiredDevice", err,
				"Failure at ReProvisionWiredDevice, unexpected response", ""))
			return diags
		}
	}

	return resourceSdaProvisionDeviceRead(ctx, d, m)
}

func resourceSdaProvisionDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSdaProvisionDeviceProvisionWiredDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaProvisionWiredDevice {
	request := dnacentersdkgo.RequestSdaProvisionWiredDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaProvisionDeviceReProvisionWiredDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaReProvisionWiredDevice {
	request := dnacentersdkgo.RequestSdaReProvisionWiredDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
