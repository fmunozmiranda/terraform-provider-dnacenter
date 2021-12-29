package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Adds border device in SDA Fabric

- Deletes border device from SDA Fabric
`,

		CreateContext: resourceSdaFabricBorderDeviceCreate,
		ReadContext:   resourceSdaFabricBorderDeviceRead,
		UpdateContext: resourceSdaFabricBorderDeviceUpdate,
		DeleteContext: resourceSdaFabricBorderDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"border_session_type": &schema.Schema{
							Description: `Border Session Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"connected_to_internet": &schema.Schema{
							Description: `Connected to Internet
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_autonomou_system_number": &schema.Schema{
							Description: `External Autonomous System Number  will be used to automate IP routing between Border Node and remote peer (e.g.,1-65535)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_connectivity_ip_pool_name": &schema.Schema{
							Description: `IP pool to use to automate IP routing between the border node and remote peer.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_connectivity_settings": &schema.Schema{
							Description: `External Connectivity Settings information of L3 Handoff
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
						},
						"external_domain_routing_protocol_name": &schema.Schema{
							Description: `External Domain Routing Protocol  Name. (Example: BGP)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"internal_autonomou_system_number": &schema.Schema{
							Description: `Internal Autonomouns System Number used by border node to communicate with remote peer (e.g.,1-65535)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"l3_handoff": &schema.Schema{
							Description: `L3 Handoff information
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy for device location(site should be fabric site)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network": &schema.Schema{
							Description: `Virtual Network information of L3 Hand off
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name assigned to site
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id (e.g.,2-4096 except for reserved VLANs (1002-1005, 2046, 4095))
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

func resourceSdaFabricBorderDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddsBorderDeviceInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddsBorderDeviceInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddsBorderDeviceInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress, okDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsBorderDeviceDetailFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetsBorderDeviceDetailFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		response1, restyResp1, err := client.Sda.GetsBorderDeviceDetailFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsBorderDeviceDetailFromSdaFabric", err,
				"Failure at GetsBorderDeviceDetailFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO

	}
	return diags
}

func resourceSdaFabricBorderDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricBorderDeviceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_domain_routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) {
		request.ExternalDomainRoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) {
		request.ExternalConnectivityIPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) {
		request.InternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_session_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_session_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_session_type")))) {
		request.BorderSessionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_to_internet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_to_internet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_to_internet")))) {
		request.ConnectedToInternet = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_settings")))) {
		request.ExternalConnectivitySettings = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricExternalConnectivitySettings(ctx, key+".external_connectivity_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) {
		request.ExternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l3_handoff")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l3_handoff")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l3_handoff")))) {
		request.L3Handoff = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricL3Handoff(ctx, key+".l3_handoff.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network")))) {
		request.VirtualNetwork = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricVirtualNetwork(ctx, key+".virtual_network.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricExternalConnectivitySettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricL3Handoff(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricL3Handoff {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricL3Handoff
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricVirtualNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricVirtualNetwork {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricVirtualNetwork
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
