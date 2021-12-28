package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscriptionRest() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Create Rest/Webhook Subscription Endpoint for list of registered events

- Update Rest/Webhook Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionRestCreate,
		ReadContext:   resourceEventSubscriptionRestRead,
		UpdateContext: resourceEventSubscriptionRestUpdate,
		DeleteContext: resourceEventSubscriptionRestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestEventManagementCreateRestWebhookEventSubscription', 'Elem': {'Schema': {'subscriptionId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Subscription Id (Unique UUID)\n'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Version\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description\n'}, 'subscriptionEndpoints': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'instanceId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': '(From \tGet Rest/Webhook Subscription Details --> pick instanceId)\n'}, 'subscriptionDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'connectorType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Connector Type (Must be REST)\n'}}}}}}}, 'filter': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'eventIds': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Event Ids (Comma separated event ids)\n'}}}}}}}}, 'metadata': {'item': {'operation_id': [['CreateRestWebhookEventSubscription', 'UpdateRestWebhookEventSubscription']], 'new_flat_structure': [[{'RequestEventManagementCreateRestWebhookEventSubscription': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementCreateRestWebhookEventSubscription', 'description': 'Array of RequestEventManagementCreateRestWebhookEventSubscription'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateRestWebhookEventSubscription': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id (Unique UUID)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': '(From \tGet Rest/Webhook Subscription Details --> pick instanceId)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'connectorType', 'description': 'Connector Type (Must be REST)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids (Comma separated event ids)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestEventManagementUpdateRestWebhookEventSubscription': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementUpdateRestWebhookEventSubscription', 'description': 'Array of RequestEventManagementUpdateRestWebhookEventSubscription'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateRestWebhookEventSubscription': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id (Unique UUID)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': '(From \tGet Rest/Webhook Subscription Details --> pick instanceId)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'connectorType', 'description': 'Connector Type (Must be REST)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids (Comma separated event ids)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestEventManagementCreateRestWebhookEventSubscription', 'RequestEventManagementUpdateRestWebhookEventSubscription']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateRestWebhookEventSubscription`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"event_ids": &schema.Schema{
										Description: `Event Ids (Comma separated event ids)
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"instance_id": &schema.Schema{
										Description: `(From 	Get Rest/Webhook Subscription Details --> pick instanceId)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connector_type": &schema.Schema{
													Description: `Connector Type (Must be REST)
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
						"subscription_id": &schema.Schema{
							Description: `Subscription Id (Unique UUID)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Description: `Version
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

func resourceEventSubscriptionRestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscription(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.EventManagement.CreateRestWebhookEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRestWebhookEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRestWebhookEventSubscription", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceEventSubscriptionRestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vEventIDs, okEventIDs := resourceMap["event_ids"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRestWebhookEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}

		if okEventIDs {
			queryParams1.EventIDs = vEventIDs.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetRestWebhookEventSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestWebhookEventSubscriptions", err,
				"Failure at GetRestWebhookEventSubscriptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceEventSubscriptionRestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vEventIDs, okEventIDs := resourceMap["event_ids"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscription(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EventManagement.UpdateRestWebhookEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRestWebhookEventSubscription", err, restyResp1.String(),
					"Failure at UpdateRestWebhookEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRestWebhookEventSubscription", err,
				"Failure at UpdateRestWebhookEventSubscription, unexpected response", ""))
			return diags
		}
	}

	return resourceEventSubscriptionRestRead(ctx, d, m)
}

func resourceEventSubscriptionRestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete EventSubscriptionRest on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscription {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitle {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitle{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitle {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleFilter {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionType {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionType{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionType {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeFilter {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItems {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItems{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItems {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsFilter {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschema {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschema{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschema {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaFilter {
	request := dnacentersdkgo.RequestEventManagementCreateRestWebhookEventSubscriptionschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscription {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitle {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitle{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitle {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitle{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionType {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionType{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionType {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionType{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItems {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItems{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItems {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItems{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschema {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschema{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschema {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschema{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscriptionschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
