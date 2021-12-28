package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSNMPProperties() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Discovery.

- Adds SNMP properties
`,

		CreateContext: resourceSNMPPropertiesCreate,
		ReadContext:   resourceSNMPPropertiesRead,
		UpdateContext: resourceSNMPPropertiesUpdate,
		DeleteContext: resourceSNMPPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestDiscoveryCreateUpdateSNMPProperties', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'intValue': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'systemPropertyName': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}, 'metadata': {'item': {'operation_id': ['CreateUpdateSNMPProperties'], 'new_flat_structure': [{'RequestDiscoveryCreateUpdateSNMPProperties': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemDiscoveryCreateUpdateSNMPProperties', 'description': 'Array of RequestDiscoveryCreateUpdateSNMPProperties'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemDiscoveryCreateUpdateSnmpProperties': {'type': 'obj', 'data': [{'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'intValue', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'systemPropertyName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestDiscoveryCreateUpdateSNMPProperties'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateUpdateSNMPProperties`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"int_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"system_property_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSNMPPropertiesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSNMPPropertiesCreateUpdateSNMPProperties(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Discovery.CreateUpdateSNMPProperties(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateUpdateSNMPProperties", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateUpdateSNMPProperties", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSNMPPropertiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSNMPProperties")

		response1, restyResp1, err := client.Discovery.GetSNMPProperties()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSNMPProperties", err,
				"Failure at GetSNMPProperties, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceSNMPPropertiesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSNMPPropertiesRead(ctx, d, m)
}

func resourceSNMPPropertiesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SNMPProperties on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSNMPPropertiesCreateUpdateSNMPProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPProperties {
	request := dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesTitle {
	request := []dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesTitle{}
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
		i := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesTitle {
	request := dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".int_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".int_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".int_value")))) {
		request.IntValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_property_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_property_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_property_name")))) {
		request.SystemPropertyName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesType {
	request := []dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesType{}
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
		i := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesType {
	request := dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".int_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".int_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".int_value")))) {
		request.IntValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_property_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_property_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_property_name")))) {
		request.SystemPropertyName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesItems {
	request := []dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesItems{}
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
		i := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesItems {
	request := dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".int_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".int_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".int_value")))) {
		request.IntValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_property_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_property_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_property_name")))) {
		request.SystemPropertyName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesschema {
	request := []dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesschema{}
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
		i := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesschema {
	request := dnacentersdkgo.RequestDiscoveryCreateUpdateSNMPPropertiesschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".int_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".int_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".int_value")))) {
		request.IntValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_property_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_property_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_property_name")))) {
		request.SystemPropertyName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
