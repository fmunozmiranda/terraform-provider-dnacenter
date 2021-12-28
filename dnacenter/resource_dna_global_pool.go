package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalPool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to update global pool

- API to create global pool.

- API to delete global IP pool.
`,

		CreateContext: resourceGlobalPoolCreate,
		ReadContext:   resourceGlobalPoolRead,
		UpdateContext: resourceGlobalPoolUpdate,
		DeleteContext: resourceGlobalPoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'settings': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'ippool': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'dhcpServerIps': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Dhcp Server Ips'}, 'ipPoolCidr': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ip Pool Cidr'}, 'IpAddressSpace': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ip Address Space. Allowed values are IPv6 or IPv4.'}, 'dnsServerIps': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Dns Server Ips'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'gateway': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Gateway'}, 'ipPoolName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ip Pool Name'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}}}}}}}, 'id': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'id path parameter. global pool id\n'}}}}}, 'metadata': {'item': {'operation_id': [['CreateGlobalPool', 'UpdateGlobalPool']], 'new_flat_structure': [[{'RequestNetworkSettingsCreateGlobalPool': {'type': 'obj', 'data': [{'name': 'settings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestNetworkSettingsCreateGlobalPoolSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestNetworkSettingsCreateGlobalPoolSettings': {'type': 'obj', 'data': [{'name': 'ippool', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestNetworkSettingsCreateGlobalPoolSettingsIppool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestNetworkSettingsCreateGlobalPoolSettingsIppool': {'type': 'obj', 'data': [{'name': 'ipPoolName', 'description': 'Ip Pool Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipPoolCidr', 'description': 'Ip Pool Cidr', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'gateway', 'description': 'Gateway', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dhcpServerIps', 'description': 'Dhcp Server Ips', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'dnsServerIps', 'description': 'Dns Server Ips', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'IpAddressSpace', 'description': 'Ip Address Space. Allowed values are IPv6 or IPv4.', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestNetworkSettingsUpdateGlobalPool': {'type': 'obj', 'data': [{'name': 'settings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestNetworkSettingsUpdateGlobalPoolSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestNetworkSettingsUpdateGlobalPoolSettings': {'type': 'obj', 'data': [{'name': 'ippool', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestNetworkSettingsUpdateGlobalPoolSettingsIppool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestNetworkSettingsUpdateGlobalPoolSettingsIppool': {'type': 'obj', 'data': [{'name': 'ipPoolName', 'description': 'Ip Pool Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'gateway', 'description': 'Gateway', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dhcpServerIps', 'description': 'Dhcp Server Ips', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'dnsServerIps', 'description': 'Dns Server Ips', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestNetworkSettingsCreateGlobalPool', 'RequestNetworkSettingsUpdateGlobalPool']], 'access_list': [[['settings'], ['settings']]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. global pool id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ippool": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address_space": &schema.Schema{
													Description: `Ip Address Space. Allowed values are IPv6 or IPv4.`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dhcp_server_ips": &schema.Schema{
													Description: `Dhcp Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"dns_server_ips": &schema.Schema{
													Description: `Dns Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"gateway": &schema.Schema{
													Description: `Gateway`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ip_pool_cidr": &schema.Schema{
													Description: `Ip Pool Cidr`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ip_pool_name": &schema.Schema{
													Description: `Ip Pool Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceGlobalPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalPoolCreateGlobalPool(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.NetworkSettings.CreateGlobalPool(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGlobalPool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGlobalPool", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalPool")
		queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetGlobalPool(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGlobalPool", err,
				"Failure at GetGlobalPool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceGlobalPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestGlobalPoolUpdateGlobalPool(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateGlobalPool(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGlobalPool", err, restyResp1.String(),
					"Failure at UpdateGlobalPool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalPool", err,
				"Failure at UpdateGlobalPool, unexpected response", ""))
			return diags
		}
	}

	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestGlobalPoolCreateGlobalPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPool {
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPool{}
	request.Settings = expandRequestGlobalPoolCreateGlobalPoolSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolCreateGlobalPoolSettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool{}
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
		i := expandRequestGlobalPoolCreateGlobalPoolSettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolSettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool {
	request := dnacentersdkgo.RequestNetworkSettingsCreateGlobalPoolSettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_cidr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_cidr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_cidr")))) {
		request.IPPoolCidr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_space")))) {
		request.IPAddressSpace = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPool {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPool{}
	request.Settings = expandRequestGlobalPoolUpdateGlobalPoolSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool {
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool{}
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
		i := expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolSettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
