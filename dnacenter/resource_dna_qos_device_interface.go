package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceQosDeviceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing qos device interface infos associate with network device id

- Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as
WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces

- Delete all qos device interface infos associate with network device id
`,

		CreateContext: resourceQosDeviceInterfaceCreate,
		ReadContext:   resourceQosDeviceInterfaceRead,
		UpdateContext: resourceQosDeviceInterfaceUpdate,
		DeleteContext: resourceQosDeviceInterfaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo', 'Elem': {'Schema': {'qosDeviceInterfaceInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'interfaceId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface id\n'}, 'instanceId': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Instance id\n'}, 'label': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'SP Profile name\n'}, 'dmvpnRemoteSitesBw': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Dmvpn remote sites bandwidth\n'}, 'role': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface role\n'}, 'uploadBW': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Upload bandwidth\n'}, 'interfaceName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface name\n'}}}}, 'networkDeviceId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Network device id\n'}, 'excludedInterfaces': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Excluded interfaces ids\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device name\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Id of Qos device info\n'}}}}}, 'metadata': {'item': {'operation_id': [['CreateQosDeviceInterfaceInfo', 'UpdateQosDeviceInterfaceInfo']], 'new_flat_structure': [[{'RequestApplicationPolicyCreateQosDeviceInterfaceInfo': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo', 'description': 'Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Device name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'excludedInterfaces', 'description': 'Excluded interfaces ids\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'networkDeviceId', 'description': 'Network device id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosDeviceInterfaceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo': {'type': 'obj', 'data': [{'name': 'dmvpnRemoteSitesBw', 'description': 'Dmvpn remote sites bandwidth\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]int'}, {'name': 'interfaceId', 'description': 'Interface id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceName', 'description': 'Interface name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'label', 'description': 'SP Profile name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': 'Interface role\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'uploadBW', 'description': 'Upload bandwidth\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestApplicationPolicyUpdateQosDeviceInterfaceInfo': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo', 'description': 'Array of RequestApplicationPolicyUpdateQosDeviceInterfaceInfo'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo': {'type': 'obj', 'data': [{'name': 'id', 'description': 'Id of Qos device info\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Device name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'excludedInterfaces', 'description': 'Excluded interfaces ids\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'networkDeviceId', 'description': 'Network device id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosDeviceInterfaceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dmvpnRemoteSitesBw', 'description': 'Dmvpn remote sites bandwidth\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]int'}, {'name': 'interfaceId', 'description': 'Interface id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceName', 'description': 'Interface name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'label', 'description': 'SP Profile name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': 'Interface role\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'uploadBW', 'description': 'Upload bandwidth\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestApplicationPolicyCreateQosDeviceInterfaceInfo', 'RequestApplicationPolicyUpdateQosDeviceInterfaceInfo']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"excluded_interfaces": &schema.Schema{
							Description: `Excluded interfaces ids
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `Id of Qos device info
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device id
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"qos_device_interface_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dmvpn_remote_sites_bw": &schema.Schema{
										Description: `Dmvpn remote sites bandwidth
`,
										Type:     schema.TypeList,
										Optional: true,
									},
									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interface_id": &schema.Schema{
										Description: `Interface id
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"label": &schema.Schema{
										Description: `SP Profile name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"role": &schema.Schema{
										Description: `Interface role
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"upload_bw": &schema.Schema{
										Description: `Upload bandwidth
`,
										Type:     schema.TypeInt,
										Optional: true,
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

func resourceQosDeviceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.ApplicationPolicy.CreateQosDeviceInterfaceInfo(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateQosDeviceInterfaceInfo", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateQosDeviceInterfaceInfo", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetworkDeviceID, okNetworkDeviceID := resourceMap["network_device_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetQosDeviceInterfaceInfo")
		queryParams1 := dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetQosDeviceInterfaceInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetQosDeviceInterfaceInfo", err,
				"Failure at GetQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceQosDeviceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetworkDeviceID, okNetworkDeviceID := resourceMap["network_device_id"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ApplicationPolicy.UpdateQosDeviceInterfaceInfo(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateQosDeviceInterfaceInfo", err, restyResp1.String(),
					"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateQosDeviceInterfaceInfo", err,
				"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}
	}

	return resourceQosDeviceInterfaceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitle {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitle{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitle {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoType {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoType{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoType {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItems {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItems{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItems {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschema {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschema{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschema {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitle {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitle{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitle {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTitleQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoType {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoType{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoType {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoTypeQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItems {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItems{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItems {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschema {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschema{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschema {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoSchemaQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfoschemaQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = v
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
