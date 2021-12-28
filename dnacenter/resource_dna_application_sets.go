package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationSets() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s
`,

		CreateContext: resourceApplicationSetsCreate,
		ReadContext:   resourceApplicationSetsRead,
		UpdateContext: resourceApplicationSetsUpdate,
		DeleteContext: resourceApplicationSetsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestApplicationPolicyCreateApplicationSet', 'Elem': {'Schema': {'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}}, 'metadata': {'item': {'operation_id': ['CreateApplicationSet'], 'new_flat_structure': [{'RequestApplicationPolicyCreateApplicationSet': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemApplicationPolicyCreateApplicationSet', 'description': 'Array of RequestApplicationPolicyCreateApplicationSet'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationSet': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestApplicationPolicyCreateApplicationSet'], 'access_list': [[]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationSet`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceApplicationSetsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestApplicationSetsCreateApplicationSet(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationSet(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationSet", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceApplicationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vName, okName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationSets")
		queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationSets(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationSets", err,
				"Failure at GetApplicationSets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceApplicationSetsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestApplicationSetsCreateApplicationSet(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestApplicationSetsCreateApplicationSetTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestApplicationSetsCreateApplicationSetTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestApplicationSetsCreateApplicationSetItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestApplicationSetsCreateApplicationSetSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetTitle {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetTitle{}
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
		i := expandRequestApplicationSetsCreateApplicationSetTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetTitle {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetType {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetType{}
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
		i := expandRequestApplicationSetsCreateApplicationSetType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetType {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetItems {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetItems{}
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
		i := expandRequestApplicationSetsCreateApplicationSetItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetItems {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetschema {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetschema{}
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
		i := expandRequestApplicationSetsCreateApplicationSetSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetschema {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSetschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
