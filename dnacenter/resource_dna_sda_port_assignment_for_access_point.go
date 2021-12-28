package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaPortAssignmentForAccessPoint() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Add Port assignment for access point in SDA Fabric

- Delete Port assignment for access point in SDA Fabric
`,

		CreateContext: resourceSdaPortAssignmentForAccessPointCreate,
		ReadContext:   resourceSdaPortAssignmentForAccessPointRead,
		UpdateContext: resourceSdaPortAssignmentForAccessPointUpdate,
		DeleteContext: resourceSdaPortAssignmentForAccessPointDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'siteNameHierarchy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Name Hierarchy should be a valid fabric site name hierarchy. e.g Global/USA/San Jose\n'}, 'deviceManagementIpAddress': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Management Ip Address of the edge device \n'}, 'interfaceName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface Name of the edge device \n'}, 'dataIpAddressPoolName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ip Pool Name, that is assigned to INFRA_VN  \n'}, 'authenticateTemplateName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Authenticate TemplateName associated to siteNameHierarchy.\n'}, 'interfaceDescription': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Details or note of interface assignment\n'}}}}}, 'metadata': {'item': {'operation_id': ['AddPortAssignmentForAccessPointInSDAFabric'], 'new_flat_structure': [{'RequestSdaAddPortAssignmentForAccessPointInSDAFabric': {'type': 'obj', 'data': [{'name': 'siteNameHierarchy', 'description': 'Site Name Hierarchy should be a valid fabric site name hierarchy. e.g Global/USA/San Jose\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceManagementIpAddress', 'description': 'Management Ip Address of the edge device \n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceName', 'description': 'Interface Name of the edge device \n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataIpAddressPoolName', 'description': 'Ip Pool Name, that is assigned to INFRA_VN  \n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticateTemplateName', 'description': 'Authenticate TemplateName associated to siteNameHierarchy.\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceDescription', 'description': 'Details or note of interface assignment\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestSdaAddPortAssignmentForAccessPointInSDAFabric'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate TemplateName associated to siteNameHierarchy.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_ip_address_pool_name": &schema.Schema{
							Description: `Ip Pool Name, that is assigned to INFRA_VN  
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the edge device 
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_description": &schema.Schema{
							Description: `Details or note of interface assignment
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name of the edge device 
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy should be a valid fabric site name hierarchy. e.g Global/USA/San Jose
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

func resourceSdaPortAssignmentForAccessPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaPortAssignmentForAccessPointAddPortAssignmentForAccessPointInSdaFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddPortAssignmentForAccessPointInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddPortAssignmentForAccessPointInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddPortAssignmentForAccessPointInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSdaPortAssignmentForAccessPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress, okDeviceManagementIPAddress := resourceMap["device_management_ip_address"]
	vInterfaceName, okInterfaceName := resourceMap["interface_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPortAssignmentForAccessPointInSdaFabric")
		queryParams1 := dnacentersdkgo.GetPortAssignmentForAccessPointInSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		queryParams1.InterfaceName = vInterfaceName.(string)

		response1, restyResp1, err := client.Sda.GetPortAssignmentForAccessPointInSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortAssignmentForAccessPointInSdaFabric", err,
				"Failure at GetPortAssignmentForAccessPointInSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceSdaPortAssignmentForAccessPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaPortAssignmentForAccessPointRead(ctx, d, m)
}

func resourceSdaPortAssignmentForAccessPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSdaPortAssignmentForAccessPointAddPortAssignmentForAccessPointInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddPortAssignmentForAccessPointInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddPortAssignmentForAccessPointInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_ip_address_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_ip_address_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_ip_address_pool_name")))) {
		request.DataIPAddressPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_description")))) {
		request.InterfaceDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
