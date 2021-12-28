package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSensor() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Sensors.

- Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID

- Intent API to delete an existing SENSOR test template
`,

		CreateContext: resourceSensorCreate,
		ReadContext:   resourceSensorRead,
		UpdateContext: resourceSensorUpdate,
		DeleteContext: resourceSensorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ssids': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ssid': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Ssid'}, 'profileName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Profile Name'}, 'authType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Auth Type'}, 'thirdParty': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'selected': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Selected'}}}}, 'psk': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Psk'}, 'tests': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'config': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Config', 'Elem': {'Schema': {}}}}}}, 'categories': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Categories'}, 'qosPolicy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Qos Policy'}}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'connection': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Connection'}, 'apCoverage': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'bands': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Bands'}, 'numberOfApsToTest': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Number Of Aps To Test'}, 'rssiThreshold': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Rssi Threshold'}}}}, 'modelVersion': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Model Version'}}}}}, 'metadata': {'item': {'operation_id': ['CreateSensorTestTemplate'], 'new_flat_structure': [{'RequestSensorsCreateSensorTestTemplate': {'type': 'obj', 'data': [{'name': 'ssids', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestSensorsCreateSensorTestTemplateSsids'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'connection', 'description': 'Connection', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'apCoverage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestSensorsCreateSensorTestTemplateApCoverage'}, {'name': 'modelVersion', 'description': 'Model Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSensorsCreateSensorTestTemplateSsids': {'type': 'obj', 'data': [{'name': 'ssid', 'description': 'Ssid', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'profileName', 'description': 'Profile Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authType', 'description': 'Auth Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'thirdParty', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestSensorsCreateSensorTestTemplateSsidsThirdParty'}, {'name': 'psk', 'description': 'Psk', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tests', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestSensorsCreateSensorTestTemplateSsidsTests'}, {'name': 'categories', 'description': 'Categories', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'qosPolicy', 'description': 'Qos Policy', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSensorsCreateSensorTestTemplateSsidsThirdParty': {'type': 'obj', 'data': [{'name': 'selected', 'description': 'Selected', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSensorsCreateSensorTestTemplateSsidsTests': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'config', 'description': 'Config', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestSensorsCreateSensorTestTemplateSsidsTestsConfig'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSensorsCreateSensorTestTemplateSsidsTestsConfig': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestSensorsCreateSensorTestTemplateApCoverage': {'type': 'obj', 'data': [{'name': 'bands', 'description': 'Bands', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'numberOfApsToTest', 'description': 'Number Of Aps To Test', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rssiThreshold', 'description': 'Rssi Threshold', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestSensorsCreateSensorTestTemplate'], 'access_list': [['ssids', 'name', 'connection', 'apCoverage', 'modelVersion']]}}}
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `Bands`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number Of Aps To Test`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `Rssi Threshold`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `Connection`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"model_version": &schema.Schema{
							Description: `Model Version`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"categories": &schema.Schema{
										Description: `Categories`,
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"profile_name": &schema.Schema{
										Description: `Profile Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"psk": &schema.Schema{
										Description: `Psk`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"qos_policy": &schema.Schema{
										Description: `Qos Policy`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Description: `Config`,
													Type:        schema.TypeList,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `Selected`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
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

func resourceSensorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSensorCreateSensorTestTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sensors.CreateSensorTestTemplate(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSensorTestTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSensorTestTemplate", err))
		return diags
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceSensorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID, okSiteID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Sensors")
		queryParams1 := dnacentersdkgo.SensorsQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.Sensors.Sensors(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Sensors", err,
				"Failure at Sensors, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceSensorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSensorRead(ctx, d, m)
}

func resourceSensorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestSensorCreateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssids")))) {
		request.SSIDs = expandRequestSensorCreateSensorTestTemplateSSIDsArray(ctx, key+".ssids", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection")))) {
		request.Connection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_coverage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_coverage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_coverage")))) {
		request.ApCoverage = expandRequestSensorCreateSensorTestTemplateApCoverageArray(ctx, key+".ap_coverage", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_version")))) {
		request.ModelVersion = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = expandRequestSensorCreateSensorTestTemplateSSIDsThirdParty(ctx, key+".third_party.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorCreateSensorTestTemplateSSIDsTestsArray(ctx, key+".tests", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsThirdParty(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected")))) {
		request.Selected = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDsTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfigArray(ctx, key+".config", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	var request dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateApCoverageArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
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
		i := expandRequestSensorCreateSensorTestTemplateApCoverage(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateSensorTestTemplateApCoverage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_aps_to_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) {
		request.NumberOfApsToTest = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
