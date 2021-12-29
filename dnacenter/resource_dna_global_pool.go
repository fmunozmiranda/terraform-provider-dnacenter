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
			queryParams1.Offset = vOffset
		}
		if okLimit {
			queryParams1.Limit = vLimit
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

//TODO
