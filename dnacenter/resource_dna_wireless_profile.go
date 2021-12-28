package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- Delete the Wireless Profile from DNAC whose name is provided.

- Updates the wireless Network Profile with updated details provided. All sites to be present in the network profile
should be provided.

- Creates Wireless Network Profile on DNAC and associates sites and SSIDs to it.
`,

		CreateContext: resourceWirelessProfileCreate,
		ReadContext:   resourceWirelessProfileRead,
		UpdateContext: resourceWirelessProfileUpdate,
		DeleteContext: resourceWirelessProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'profileDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Profile Name\n'}, 'sites': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"]) \n'}, 'ssidDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ssid Name\n'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ssid Type(enum: Enterprise/Guest)\n'}, 'enableFabric': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'true is ssid is fabric else false\n'}, 'flexConnect': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'enableFlexConnect': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'true if flex connect is enabled else false\n'}, 'localToVlan': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Local To Vlan\n'}}}}, 'interfaceName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Interface Name\n'}}}}}}}, 'wirelessProfileName': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'wirelessProfileName path parameter. Wireless Profile Name\n'}}}}}, 'metadata': {'item': {'operation_id': [['CreateWirelessProfile', 'UpdateWirelessProfile']], 'new_flat_structure': [[{'RequestWirelessCreateWirelessProfile': {'type': 'obj', 'data': [{'name': 'profileDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestWirelessCreateWirelessProfileProfileDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessCreateWirelessProfileProfileDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Profile Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sites', 'description': 'array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"]) \n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'ssidDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestWirelessCreateWirelessProfileProfileDetailsSsidDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessCreateWirelessProfileProfileDetailsSsidDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Ssid Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Ssid Type(enum: Enterprise/Guest)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enableFabric', 'description': 'true is ssid is fabric else false\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'flexConnect', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestWirelessCreateWirelessProfileProfileDetailsSsidDetailsFlexConnect'}, {'name': 'interfaceName', 'description': 'Interface Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessCreateWirelessProfileProfileDetailsSsidDetailsFlexConnect': {'type': 'obj', 'data': [{'name': 'enableFlexConnect', 'description': 'true if flex connect is enabled else false\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'localToVlan', 'description': 'Local To Vlan\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestWirelessUpdateWirelessProfile': {'type': 'obj', 'data': [{'name': 'profileDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestWirelessUpdateWirelessProfileProfileDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessUpdateWirelessProfileProfileDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Profile Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sites', 'description': 'array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"]) \n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'ssidDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestWirelessUpdateWirelessProfileProfileDetailsSsidDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessUpdateWirelessProfileProfileDetailsSsidDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Ssid Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Ssid Type(enum: Enterprise/Guest)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enableFabric', 'description': 'true is ssid is fabric else false\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'flexConnect', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestWirelessUpdateWirelessProfileProfileDetailsSsidDetailsFlexConnect'}, {'name': 'interfaceName', 'description': 'Interface Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestWirelessUpdateWirelessProfileProfileDetailsSsidDetailsFlexConnect': {'type': 'obj', 'data': [{'name': 'enableFlexConnect', 'description': 'true if flex connect is enabled else false\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'localToVlan', 'description': 'Local To Vlan\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestWirelessCreateWirelessProfile', 'RequestWirelessUpdateWirelessProfile']], 'access_list': [[['profileDetails'], ['profileDetails']]]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"profile_details": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Profile Name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"sites": &schema.Schema{
										Description: `array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"]) 
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_fabric": &schema.Schema{
													Description: `true is ssid is fabric else false
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"flex_connect": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_flex_connect": &schema.Schema{
																Description: `true if flex connect is enabled else false
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"local_to_vlan": &schema.Schema{
																Description: `Local To Vlan
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"interface_name": &schema.Schema{
													Description: `Interface Name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Ssid Name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Ssid Type(enum: Enterprise/Guest)
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
						"wireless_profile_name": &schema.Schema{
							Description: `wirelessProfileName path parameter. Wireless Profile Name
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessProfileCreateWirelessProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vWirelessProfileName, okWirelessProfileName := resourceItem["wireless_profile_name"]
	vvWirelessProfileName := interfaceToString(vWirelessProfileName)
	resp1, restyResp1, err := client.Wireless.CreateWirelessProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateWirelessProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateWirelessProfile", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["wireless_profile_name"] = vvWirelessProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceWirelessProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vProfileName, okProfileName := resourceMap["profile_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetWirelessProfile")
		queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}

		if okProfileName {
			queryParams1.ProfileName = vProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetWirelessProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWirelessProfile", err,
				"Failure at GetWirelessProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceWirelessProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vProfileName, okProfileName := resourceMap["profile_name"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestWirelessProfileUpdateWirelessProfile(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateWirelessProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateWirelessProfile", err, restyResp1.String(),
					"Failure at UpdateWirelessProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateWirelessProfile", err,
				"Failure at UpdateWirelessProfile, unexpected response", ""))
			return diags
		}
	}

	return resourceWirelessProfileRead(ctx, d, m)
}

func resourceWirelessProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestWirelessProfileCreateWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfile {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfile{}
	request.ProfileDetails = expandRequestWirelessProfileCreateWirelessProfileProfileDetails(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetails {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sites")))) {
		request.Sites = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails {
	request := []dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails{}
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
		i := expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_flex_connect")))) {
		request.EnableFlexConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_to_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_to_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_to_vlan")))) {
		request.LocalToVLAN = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfile {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfile{}
	request.ProfileDetails = expandRequestWirelessProfileUpdateWirelessProfileProfileDetails(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetails {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sites")))) {
		request.Sites = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails {
	request := []dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails{}
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
		i := expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_flex_connect")))) {
		request.EnableFlexConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_to_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_to_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_to_vlan")))) {
		request.LocalToVLAN = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
