package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceReplacement() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Device Replacement.

- UnMarks device for replacement

- Marks device for replacement
`,

		CreateContext: resourceDeviceReplacementCreate,
		ReadContext:   resourceDeviceReplacementRead,
		UpdateContext: resourceDeviceReplacementUpdate,
		DeleteContext: resourceDeviceReplacementDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestDeviceReplacementMarkDeviceForReplacement', 'Elem': {'Schema': {'creationTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'family': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'faultyDeviceId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'faultyDeviceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'faultyDevicePlatform': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'faultyDeviceSerialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'neighbourDeviceId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'networkReadinessTaskId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'replacementDevicePlatform': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'replacementDeviceSerialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'replacementStatus': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'replacementTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'workflowId': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}, 'metadata': {'item': {'operation_id': [['MarkDeviceForReplacement', 'UnMarkDeviceForReplacement']], 'new_flat_structure': [[{'RequestDeviceReplacementMarkDeviceForReplacement': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemDeviceReplacementMarkDeviceForReplacement', 'description': 'Array of RequestDeviceReplacementMarkDeviceForReplacement'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemDeviceReplacementMarkDeviceForReplacement': {'type': 'obj', 'data': [{'name': 'creationTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'family', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDevicePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'neighbourDeviceId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'networkReadinessTaskId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementDevicePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementDeviceSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'workflowId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestDeviceReplacementUnMarkDeviceForReplacement': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemDeviceReplacementUnMarkDeviceForReplacement', 'description': 'Array of RequestDeviceReplacementUnMarkDeviceForReplacement'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemDeviceReplacementUnmarkDeviceForReplacement': {'type': 'obj', 'data': [{'name': 'creationTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'family', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDevicePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'faultyDeviceSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'neighbourDeviceId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'networkReadinessTaskId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementDevicePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementDeviceSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'replacementTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'workflowId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestDeviceReplacementMarkDeviceForReplacement', 'RequestDeviceReplacementUnMarkDeviceForReplacement']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestDeviceReplacementMarkDeviceForReplacement`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_platform": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"neighbour_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_readiness_task_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_device_platform": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_device_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"workflow_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceDeviceReplacementCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceReplacementMarkDeviceForReplacement(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.DeviceReplacement.MarkDeviceForReplacement(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing MarkDeviceForReplacement", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing MarkDeviceForReplacement", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceDeviceReplacementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFaultyDeviceName, okFaultyDeviceName := resourceMap["faulty_device_name"]
	vFaultyDevicePlatform, okFaultyDevicePlatform := resourceMap["faulty_device_platform"]
	vReplacementDevicePlatform, okReplacementDevicePlatform := resourceMap["replacement_device_platform"]
	vFaultyDeviceSerialNumber, okFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber, okReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]
	vReplacementStatus, okReplacementStatus := resourceMap["replacement_status"]
	vFamily, okFamily := resourceMap["family"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ReturnListOfReplacementDevicesWithReplacementDetails")
		queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

		if okFaultyDeviceName {
			queryParams1.FaultyDeviceName = vFaultyDeviceName.(string)
		}
		if okFaultyDevicePlatform {
			queryParams1.FaultyDevicePlatform = vFaultyDevicePlatform.(string)
		}
		if okReplacementDevicePlatform {
			queryParams1.ReplacementDevicePlatform = vReplacementDevicePlatform.(string)
		}
		if okFaultyDeviceSerialNumber {
			queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber.(string)
		}
		if okReplacementDeviceSerialNumber {
			queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber.(string)
		}
		if okReplacementStatus {
			queryParams1.ReplacementStatus = interfaceToSliceString(vReplacementStatus)
		}
		if okFamily {
			queryParams1.Family = interfaceToSliceString(vFamily)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		response1, restyResp1, err := client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReturnListOfReplacementDevicesWithReplacementDetails", err,
				"Failure at ReturnListOfReplacementDevicesWithReplacementDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceDeviceReplacementUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFaultyDeviceName, okFaultyDeviceName := resourceMap["faulty_device_name"]
	vFaultyDevicePlatform, okFaultyDevicePlatform := resourceMap["faulty_device_platform"]
	vReplacementDevicePlatform, okReplacementDevicePlatform := resourceMap["replacement_device_platform"]
	vFaultyDeviceSerialNumber, okFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber, okReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]
	vReplacementStatus, okReplacementStatus := resourceMap["replacement_status"]
	vFamily, okFamily := resourceMap["family"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestDeviceReplacementUnmarkDeviceForReplacement(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceReplacement.UnmarkDeviceForReplacement(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UnmarkDeviceForReplacement", err, restyResp1.String(),
					"Failure at UnmarkDeviceForReplacement, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UnmarkDeviceForReplacement", err,
				"Failure at UnmarkDeviceForReplacement, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceReplacementRead(ctx, d, m)
}

func resourceDeviceReplacementDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete DeviceReplacement on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestDeviceReplacementMarkDeviceForReplacement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacement {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestDeviceReplacementMarkDeviceForReplacementTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestDeviceReplacementMarkDeviceForReplacementTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestDeviceReplacementMarkDeviceForReplacementItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestDeviceReplacementMarkDeviceForReplacementSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementTitle {
	request := []dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementTitle{}
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
		i := expandRequestDeviceReplacementMarkDeviceForReplacementTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementTitle {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementType {
	request := []dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementType{}
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
		i := expandRequestDeviceReplacementMarkDeviceForReplacementType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementType {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementItems {
	request := []dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementItems{}
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
		i := expandRequestDeviceReplacementMarkDeviceForReplacementItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementItems {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementschema {
	request := []dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementschema{}
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
		i := expandRequestDeviceReplacementMarkDeviceForReplacementSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementschema {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacementschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacement {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestDeviceReplacementUnmarkDeviceForReplacementTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestDeviceReplacementUnmarkDeviceForReplacementTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestDeviceReplacementUnmarkDeviceForReplacementItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestDeviceReplacementUnmarkDeviceForReplacementSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementTitle {
	request := []dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementTitle{}
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
		i := expandRequestDeviceReplacementUnmarkDeviceForReplacementTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementTitle {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementType {
	request := []dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementType{}
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
		i := expandRequestDeviceReplacementUnmarkDeviceForReplacementType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementType {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementItems {
	request := []dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementItems{}
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
		i := expandRequestDeviceReplacementUnmarkDeviceForReplacementItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementItems {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementschema {
	request := []dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementschema{}
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
		i := expandRequestDeviceReplacementUnmarkDeviceForReplacementSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementschema {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacementschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
