package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscription() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Event Management.

- Delete EventSubscriptions

- Update SubscriptionEndpoint to list of registered events

- Subscribe SubscriptionEndpoint to list of registered events
`,

		CreateContext: resourceEventSubscriptionCreate,
		ReadContext:   resourceEventSubscriptionRead,
		UpdateContext: resourceEventSubscriptionUpdate,
		DeleteContext: resourceEventSubscriptionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestEventManagementCreateEventSubscriptions', 'Elem': {'Schema': {'subscriptionId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Subscription Id'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Version'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'subscriptionEndpoints': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'instanceId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Instance Id'}, 'subscriptionDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'url': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Url'}, 'method': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Method'}, 'connectorType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Connector Type'}}}}}}}, 'filter': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'eventIds': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Event Ids'}}}}}}}}, 'metadata': {'item': {'operation_id': [['CreateEventSubscriptions', 'UpdateEventSubscriptions']], 'new_flat_structure': [[{'RequestEventManagementCreateEventSubscriptions': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementCreateEventSubscriptions', 'description': 'Array of RequestEventManagementCreateEventSubscriptions'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEventSubscriptions': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateEventSubscriptionsFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'url', 'description': 'Url', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'method', 'description': 'Method', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'connectorType', 'description': 'Connector Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEventSubscriptionsFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestEventManagementUpdateEventSubscriptions': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementUpdateEventSubscriptions', 'description': 'Array of RequestEventManagementUpdateEventSubscriptions'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEventSubscriptions': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateEventSubscriptionsFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': 'Instance Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'url', 'description': 'Url', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'method', 'description': 'Method', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'connectorType', 'description': 'Connector Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEventSubscriptionsFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestEventManagementCreateEventSubscriptions', 'RequestEventManagementUpdateEventSubscriptions']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateEventSubscriptions`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"event_ids": &schema.Schema{
										Description: `Event Ids`,
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connector_type": &schema.Schema{
													Description: `Connector Type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"method": &schema.Schema{
													Description: `Method`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"subscription_id": &schema.Schema{
							Description: `Subscription Id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceEventSubscriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEventSubscriptionCreateEventSubscriptions(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.EventManagement.CreateEventSubscriptions(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEventSubscriptions", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEventSubscriptions", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceEventSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetEventSubscriptionsQueryParams{}

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

		response1, restyResp1, err := client.EventManagement.GetEventSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEventSubscriptions", err,
				"Failure at GetEventSubscriptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceEventSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		request1 := expandRequestEventSubscriptionUpdateEventSubscriptions(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EventManagement.UpdateEventSubscriptions(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEventSubscriptions", err, restyResp1.String(),
					"Failure at UpdateEventSubscriptions, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEventSubscriptions", err,
				"Failure at UpdateEventSubscriptions, unexpected response", ""))
			return diags
		}
	}

	return resourceEventSubscriptionRead(ctx, d, m)
}

func resourceEventSubscriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestEventSubscriptionCreateEventSubscriptions(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptions {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionCreateEventSubscriptionsTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionCreateEventSubscriptionsTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionCreateEventSubscriptionsItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionCreateEventSubscriptionsSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitle {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitle{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitle {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitle{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionCreateEventSubscriptionsTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsType {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsType{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsType {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsType{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionCreateEventSubscriptionsTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItems {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItems{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItems {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItems{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionCreateEventSubscriptionsItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschema {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschema{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschema {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschema{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionCreateEventSubscriptionsSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEventSubscriptionsschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptions(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptions {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionUpdateEventSubscriptionsTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionUpdateEventSubscriptionsTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionUpdateEventSubscriptionsItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitle {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitle{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitle {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitle{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionUpdateEventSubscriptionsTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsType {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsType{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsType {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsType{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionUpdateEventSubscriptionsTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItems {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItems{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItems {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItems{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionUpdateEventSubscriptionsItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschema {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschema{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschema {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschema{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEventSubscriptionsschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
