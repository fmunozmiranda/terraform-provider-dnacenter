package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_options": &schema.Schema{
							Description: `Client Options`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"configure_external_dhcp": &schema.Schema{
							Description: `Configure External Dhcp`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"context_key": &schema.Schema{
										Description: `Context Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"context_value": &schema.Schema{
										Description: `Context Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"owner": &schema.Schema{
										Description: `Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"create_time": &schema.Schema{
							Description: `Create Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dhcp_server_ips": &schema.Schema{
							Description: `Dhcp Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"dns_server_ips": &schema.Schema{
							Description: `Dns Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"gateways": &schema.Schema{
							Description: `Gateways`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_cidr": &schema.Schema{
							Description: `Ip Pool Cidr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"overlapping": &schema.Schema{
							Description: `Overlapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"owner": &schema.Schema{
							Description: `Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_uuid": &schema.Schema{
							Description: `Parent Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"shared": &schema.Schema{
							Description: `Shared`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"total_ip_address_count": &schema.Schema{
							Description: `Total Ip Address Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"used_ip_address_count": &schema.Schema{
							Description: `Used Ip Address Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"used_percentage": &schema.Schema{
							Description: `Used Percentage`,
							Type:        schema.TypeString,
							Computed:    true,
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

	vIpPoolName := resourceItem["ip_pool_name"]
	vvIpPoolName := interfaceToString(vIpPoolName)
	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

	response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vvIpPoolName)

	if err != nil || response1 != nil {
		resourceMap := make(map[string]string)
		resourceMap["ip_pool_name"] = vvIpPoolName
		d.SetId(joinResourceID(resourceMap))
	}

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
	resourceMap["ip_pool_name"] = vvIpPoolName
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalPool")
		queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

		response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vIpPoolName)

		if err != nil || response1 == nil {

			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGlobalPool", err,
				"Failure at GetGlobalPool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenNetworkSettingsGetGlobalPoolItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalPool search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["ip_pool_name"]

	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalPool", err,
			"Failure at GetGlobalPool, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", queryParams1)
		request1 := expandRequestGlobalPoolUpdateGlobalPool(ctx, "parameters.0", d)
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
	vID := resourceMap["ip_pool_name"]

	queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalPool", err,
			"Failure at GetGlobalPool, unexpected response", ""))
		return diags
	}

	response1, restyResp1, err := client.NetworkSettings.DeleteGlobalIPPool(vID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGlobalIPPool", err, restyResp1.String(),
				"Failure at DeleteGlobalIPPool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGlobalIPPool", err,
			"Failure at DeleteGlobalIPPool, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

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

func searchNetworkSettingsGetGlobalPool(m interface{}, queryParams dnacentersdkgo.GetGlobalPoolQueryParams, vID string) (*dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse
	offset := 1
	queryParams.Offset = strconv.Itoa(offset)

	//var allItems []*dnacenterskgo.ResponseItemApplicationPolicyGetApplications
	nResponse, _, err := client.NetworkSettings.GetGlobalPool(&queryParams)
	maxPageSize := len(*nResponse.Response)
	//maxPageSize := 10
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.IPPoolName {
				foundItem = &item
				return foundItem, err
			}
			//allItems = append(allItems, &item)
		}

		queryParams.Limit = strconv.Itoa(maxPageSize)
		offset += maxPageSize
		queryParams.Offset = strconv.Itoa(offset)
		nResponse, _, err = client.NetworkSettings.GetGlobalPool(&queryParams)
	}
	return foundItem, err
}
