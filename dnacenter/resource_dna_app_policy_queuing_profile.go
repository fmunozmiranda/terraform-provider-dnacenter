package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppPolicyQueuingProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing custom application queuing profile

- Create new custom application queuing profile

- Delete existing custom application policy queuing profile by id
`,

		CreateContext: resourceAppPolicyQueuingProfileCreate,
		ReadContext:   resourceAppPolicyQueuingProfileRead,
		UpdateContext: resourceAppPolicyQueuingProfileUpdate,
		DeleteContext: resourceAppPolicyQueuingProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile', 'Elem': {'Schema': {'clause': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'instanceId': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Instance id\n'}, 'interfaceSpeedBandwidthClauses': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'instanceId': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Instance id\n'}, 'interfaceSpeed': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface speed\n'}, 'tcBandwidthSettings': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'trafficClass': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Traffic Class\n'}, 'instanceId': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Instance id\n'}, 'bandwidthPercentage': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Bandwidth percentage\n'}}}}}}}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type\n'}, 'isCommonBetweenAllInterfaceSpeeds': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is common between all interface speeds\n'}, 'tcDscpSettings': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'trafficClass': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Traffic Class\n'}, 'instanceId': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Instance id\n'}, 'dscp': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Dscp value\n'}}}}}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Queueing profile name\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Free test description\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Id of Queueing profile\n'}}}}}, 'metadata': {'item': {'operation_id': [['CreateApplicationPolicyQueuingProfile', 'UpdateApplicationPolicyQueuingProfile']], 'new_flat_structure': [[{'RequestApplicationPolicyCreateApplicationPolicyQueuingProfile': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile', 'description': 'Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile': {'type': 'obj', 'data': [{'name': 'description', 'description': 'Free test description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Queueing profile name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'clause', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause': {'type': 'obj', 'data': [{'name': 'type', 'description': 'Type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'isCommonBetweenAllInterfaceSpeeds', 'description': 'Is common between all interface speeds\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'interfaceSpeedBandwidthClauses', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses'}, {'name': 'tcDscpSettings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses': {'type': 'obj', 'data': [{'name': 'interfaceSpeed', 'description': 'Interface speed\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tcBandwidthSettings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings': {'type': 'obj', 'data': [{'name': 'bandwidthPercentage', 'description': 'Bandwidth percentage\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'trafficClass', 'description': 'Traffic Class\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings': {'type': 'obj', 'data': [{'name': 'dscp', 'description': 'Dscp value\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'trafficClass', 'description': 'Traffic Class\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile', 'description': 'Array of RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile': {'type': 'obj', 'data': [{'name': 'id', 'description': 'Id of Queueing profile\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Free test description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Queueing profile name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'clause', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': 'Type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'isCommonBetweenAllInterfaceSpeeds', 'description': 'Is common between all interface speeds\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'interfaceSpeedBandwidthClauses', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses'}, {'name': 'tcDscpSettings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'interfaceSpeed', 'description': 'Interface speed\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tcBandwidthSettings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'bandwidthPercentage', 'description': 'Bandwidth percentage\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'trafficClass', 'description': 'Traffic Class\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dscp', 'description': 'Dscp value\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'trafficClass', 'description': 'Traffic Class\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestApplicationPolicyCreateApplicationPolicyQueuingProfile', 'RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"clause": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interface_speed_bandwidth_clauses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"traffic_class": &schema.Schema{
																Description: `Traffic Class
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"tc_dscp_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Free test description
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `Id of Queueing profile
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Queueing profile name
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

func resourceAppPolicyQueuingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationPolicyQueuingProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfile", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicyQueuingProfile")
		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationPolicyQueuingProfile", err,
				"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceAppPolicyQueuingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ApplicationPolicy.UpdateApplicationPolicyQueuingProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err, restyResp1.String(),
					"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err,
				"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
	}

	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitle {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitle{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitle {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClause {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClause {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileType {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileType{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileType {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClause {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClause {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItems {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItems{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItems {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClause {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClause {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschema {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschema{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschema {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClause {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClause {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitle {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitle{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitle {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClause {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClause {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTitleClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileType {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileType{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileType {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClause {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClause {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileTypeClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItems {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItems{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItems {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClause {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClause {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileItemsClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschema {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschema{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschema {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClause {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClause {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileSchemaClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileschemaClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
