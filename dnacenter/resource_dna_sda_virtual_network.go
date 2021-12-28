package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete virtual network (VN) from SDA Fabric

- Add virtual network (VN) in SDA Fabric
`,

		CreateContext: resourceSdaVirtualNetworkCreate,
		ReadContext:   resourceSdaVirtualNetworkRead,
		UpdateContext: resourceSdaVirtualNetworkUpdate,
		DeleteContext: resourceSdaVirtualNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'virtualNetworkName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Virtual Network Name, that is created in Global level\n'}, 'siteNameHierarchy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Name Hierarchy should be a valid fabric site name hierarchy.( e.g. Global/USA/San Jose)\n'}}}}}, 'metadata': {'item': {'operation_id': ['AddVNInSDAFabric'], 'new_flat_structure': [{'RequestSdaAddVNInSDAFabric': {'type': 'obj', 'data': [{'name': 'virtualNetworkName', 'description': 'Virtual Network Name, that is created in Global level\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'siteNameHierarchy', 'description': 'Site Name Hierarchy should be a valid fabric site name hierarchy.( e.g. Global/USA/San Jose)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestSdaAddVNInSDAFabric'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy should be a valid fabric site name hierarchy.( e.g. Global/USA/San Jose)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name, that is created in Global level
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

func resourceSdaVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkAddVnInSdaFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddVnInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddVnInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddVnInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSdaVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName, okVirtualNetworkName := resourceMap["virtual_network_name"]
	vSiteNameHierarchy, okSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVnFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVnFromSdaFabric", err,
				"Failure at GetVnFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceSdaVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaVirtualNetworkRead(ctx, d, m)
}

func resourceSdaVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSdaVirtualNetworkAddVnInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddVnInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddVnInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
