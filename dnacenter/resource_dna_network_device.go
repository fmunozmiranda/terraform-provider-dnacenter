package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and delete operations on Devices.

- Deletes the network device for the given Id
`,

		CreateContext: resourceNetworkDeviceCreate,
		ReadContext:   resourceNetworkDeviceRead,
		UpdateContext: resourceNetworkDeviceUpdate,
		DeleteContext: resourceNetworkDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'id path parameter. Device ID\n'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'apManagerInterfaceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'associatedWlcIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'bootDateTime': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'collectionInterval': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'collectionStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'errorCode': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'errorDescription': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'family': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'hostname': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceCount': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inventoryStatusDetail': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'lastUpdateTime': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'lastUpdated': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'lineCardCount': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'lineCardId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'location': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'locationName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'macAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'managementIpAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'memorySize': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'platformId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'reachabilityFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'reachabilityStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'role': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'roleSource': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'serialNumber': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'series': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpContact': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpLocation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'softwareType': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'softwareVersion': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'tagCount': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'tunnelUdpPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'upTime': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'waasDeviceMode': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}, 'metadata': {'item': {'operation_id': ['GetDeviceByID'], 'new_flat_structure': [{'ResponseDevicesGetDeviceByIdResponse': {'type': 'obj', 'data': [{'name': 'apManagerInterfaceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'associatedWlcIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'bootDateTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'collectionInterval', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'collectionStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorCode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorDescription', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'family', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'hostname', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inventoryStatusDetail', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastUpdated', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lineCardCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lineCardId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'locationName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'managementIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'memorySize', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'platformId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reachabilityFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reachabilityStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'roleSource', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'series', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpLocation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tagCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tunnelUdpPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'upTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'waasDeviceMode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDevicesGetDeviceByIDResponse': {'type': 'obj', 'data': [{'name': 'apManagerInterfaceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'associatedWlcIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'bootDateTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'collectionInterval', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'collectionStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorCode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorDescription', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'family', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'hostname', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inventoryStatusDetail', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastUpdated', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lineCardCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lineCardId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'locationName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'managementIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'memorySize', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'platformId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reachabilityFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reachabilityStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'roleSource', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'series', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpLocation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tagCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tunnelUdpPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'upTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'waasDeviceMode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['ResponseDevicesGetDeviceByIdResponse'], 'access_list': [['response']]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_manager_interface_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_interval": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_status_detail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_status": &schema.Schema{
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
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"series": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_contact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_udp_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"up_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"waas_device_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. Device ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Devices.GetDeviceByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceByID", err,
				"Failure at GetDeviceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceNetworkDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDeviceRead(ctx, d, m)
}

func resourceNetworkDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
