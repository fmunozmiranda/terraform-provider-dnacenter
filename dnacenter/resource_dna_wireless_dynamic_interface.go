package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessDynamicInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Wireless.

- API to create or update an dynamic interface

- Delete a dynamic interface
`,

		CreateContext: resourceWirelessDynamicInterfaceCreate,
		ReadContext:   resourceWirelessDynamicInterfaceRead,
		UpdateContext: resourceWirelessDynamicInterfaceUpdate,
		DeleteContext: resourceWirelessDynamicInterfaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'interfaceName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'dynamic-interface name\n'}, 'vlanId': {'Optional': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Vlan Id\n'}}}}}, 'metadata': {'item': {'operation_id': ['CreateUpdateDynamicInterface'], 'new_flat_structure': [{'RequestWirelessCreateUpdateDynamicInterface': {'type': 'obj', 'data': [{'name': 'interfaceName', 'description': 'dynamic-interface name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vlanId', 'description': 'Vlan Id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestWirelessCreateUpdateDynamicInterface'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_name": &schema.Schema{
							Description: `dynamic-interface name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id
`,
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessDynamicInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessDynamicInterfaceCreateUpdateDynamicInterface(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vInterfaceName, okInterfaceName := resourceItem["interface_name"]
	vvInterfaceName := interfaceToString(vInterfaceName)
	resp1, restyResp1, err := client.Wireless.CreateUpdateDynamicInterface(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateUpdateDynamicInterface", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateUpdateDynamicInterface", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["interface_name"] = vvInterfaceName
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceWirelessDynamicInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vInterfaceName, okInterfaceName := resourceMap["interface_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDynamicInterface")
		queryParams1 := dnacentersdkgo.GetDynamicInterfaceQueryParams{}

		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetDynamicInterface(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDynamicInterface", err,
				"Failure at GetDynamicInterface, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceWirelessDynamicInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessDynamicInterfaceRead(ctx, d, m)
}

func resourceWirelessDynamicInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestWirelessDynamicInterfaceCreateUpdateDynamicInterface(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateUpdateDynamicInterface {
	request := dnacentersdkgo.RequestWirelessCreateUpdateDynamicInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
