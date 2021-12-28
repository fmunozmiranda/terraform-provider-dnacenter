package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaMulticast() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Add multicast in SDA fabric

- Delete multicast from SDA fabric
`,

		CreateContext: resourceSdaMulticastCreate,
		ReadContext:   resourceSdaMulticastRead,
		UpdateContext: resourceSdaMulticastUpdate,
		DeleteContext: resourceSdaMulticastDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'siteNameHierarchy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Full path of sda fabric siteNameHierarchy\n'}, 'multicastMethod': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Multicast Methods\n'}, 'muticastType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Muticast Type\n'}, 'multicastVnInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'virtualNetworkName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Virtual Network Name, that is associated to fabricSiteNameHierarchy\n'}, 'ipPoolName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ip Pool Name, that is reserved to fabricSiteNameHierarchy\n'}, 'externalRpIpAddress': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'External Rp Ip Address, required for muticastType=asm_with_external_rp\n'}, 'ssmInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Source-specific multicast information, required if muticastType=ssm\n', 'Elem': {'Schema': {}}}, 'ssmGroupRange': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Valid SSM group range ip address(e.g., 230.0.0.0)\n'}, 'ssmWildcardMask': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)\n'}}}}}}}}, 'metadata': {'item': {'operation_id': ['AddMulticastInSDAFabric'], 'new_flat_structure': [{'RequestSdaAddMulticastInSDAFabric': {'type': 'obj', 'data': [{'name': 'siteNameHierarchy', 'description': 'Full path of sda fabric siteNameHierarchy\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'multicastMethod', 'description': 'Multicast Methods\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'muticastType', 'description': 'Muticast Type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'multicastVnInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestSdaAddMulticastInSdaFabricMulticastVnInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSdaAddMulticastInSDAFabricMulticastVnInfo': {'type': 'obj', 'data': [{'name': 'virtualNetworkName', 'description': 'Virtual Network Name, that is associated to fabricSiteNameHierarchy\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipPoolName', 'description': 'Ip Pool Name, that is reserved to fabricSiteNameHierarchy\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'externalRpIpAddress', 'description': 'External Rp Ip Address, required for muticastType=asm_with_external_rp\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ssmInfo', 'description': 'Source-specific multicast information, required if muticastType=ssm\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo'}, {'name': 'ssmGroupRange', 'description': 'Valid SSM group range ip address(e.g., 230.0.0.0)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ssmWildcardMask', 'description': 'Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSdaAddMulticastInSDAFabricMulticastVnInfoSsmInfo': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestSdaAddMulticastInSDAFabric'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"multicast_method": &schema.Schema{
							Description: `Multicast Methods
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"multicast_vn_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_rp_ip_address": &schema.Schema{
										Description: `External Rp Ip Address, required for muticastType=asm_with_external_rp
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_pool_name": &schema.Schema{
										Description: `Ip Pool Name, that is reserved to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssm_group_range": &schema.Schema{
										Description: `Valid SSM group range ip address(e.g., 230.0.0.0)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssm_info": &schema.Schema{
										Description: `Source-specific multicast information, required if muticastType=ssm
`,
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
									},
									"ssm_wildcard_mask": &schema.Schema{
										Description: `Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name, that is associated to fabricSiteNameHierarchy
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"muticast_type": &schema.Schema{
							Description: `Muticast Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Full path of sda fabric siteNameHierarchy
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

func resourceSdaMulticastCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaMulticastAddMulticastInSdaFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddMulticastInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddMulticastInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddMulticastInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSdaMulticastRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy, okSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetMulticastDetailsFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetMulticastDetailsFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetMulticastDetailsFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMulticastDetailsFromSdaFabric", err,
				"Failure at GetMulticastDetailsFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceSdaMulticastUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaMulticastRead(ctx, d, m)
}

func resourceSdaMulticastDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSdaMulticastAddMulticastInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddMulticastInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_method")))) {
		request.MulticastMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".muticast_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".muticast_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".muticast_type")))) {
		request.MuticastType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_vn_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_vn_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_vn_info")))) {
		request.MulticastVnInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfo(ctx, key+".multicast_vn_info.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo {
	request := dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_rp_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_rp_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_rp_ip_address")))) {
		request.ExternalRpIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_info")))) {
		request.SsmInfo = expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx, key+".ssm_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_group_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_group_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_group_range")))) {
		request.SsmGroupRange = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssm_wildcard_mask")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssm_wildcard_mask")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssm_wildcard_mask")))) {
		request.SsmWildcardMask = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaMulticastAddMulticastInSdaFabricMulticastVnInfoSsmInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo {
	var request dnacentersdkgo.RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
