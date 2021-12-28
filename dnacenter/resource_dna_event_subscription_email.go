package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscriptionEmail() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Create Email Subscription Endpoint for list of registered events.

- Update Email Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionEmailCreate,
		ReadContext:   resourceEventSubscriptionEmailRead,
		UpdateContext: resourceEventSubscriptionEmailUpdate,
		DeleteContext: resourceEventSubscriptionEmailDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of RequestEventManagementCreateEmailEventSubscription', 'Elem': {'Schema': {'subscriptionId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Subscription Id (Unique UUID)\n'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Version\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description\n'}, 'subscriptionEndpoints': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'instanceId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': '(From Get Email Subscription Details --> pick InstanceId)\n'}, 'subscriptionDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'connectorType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Connector Type (Must be EMAIL)\n'}, 'fromEmailAddress': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Senders Email Address\n'}, 'toEmailAddresses': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': "Recipient's Email Addresses (Comma separated)\n"}, 'subject': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Email Subject\n'}}}}}}}, 'filter': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'eventIds': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Event Ids (Comma separated event ids)\n'}}}}}}}}, 'metadata': {'item': {'operation_id': [['CreateEmailEventSubscription', 'UpdateEmailEventSubscription']], 'new_flat_structure': [[{'RequestEventManagementCreateEmailEventSubscription': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementCreateEmailEventSubscription', 'description': 'Array of RequestEventManagementCreateEmailEventSubscription'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEmailEventSubscription': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id (Unique UUID)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateEmailEventSubscriptionFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': '(From Get Email Subscription Details --> pick InstanceId)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'connectorType', 'description': 'Connector Type (Must be EMAIL)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fromEmailAddress', 'description': 'Senders Email Address\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'toEmailAddresses', 'description': "Recipient's Email Addresses (Comma separated)\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'subject', 'description': 'Email Subject\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementCreateEmailEventSubscriptionFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids (Comma separated event ids)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestEventManagementUpdateEmailEventSubscription': {'type': 'array', 'data': [{'name': '', 'type': '[]RequestItemEventManagementUpdateEmailEventSubscription', 'description': 'Array of RequestEventManagementUpdateEmailEventSubscription'}], 'epType': 'json', 'isRef': True, 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEmailEventSubscription': {'type': 'obj', 'data': [{'name': 'subscriptionId', 'description': 'Subscription Id (Unique UUID)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionEndpoints', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints'}, {'name': 'filter', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateEmailEventSubscriptionFilter'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints': {'type': 'obj', 'data': [{'name': 'instanceId', 'description': '(From Get Email Subscription Details --> pick InstanceId)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'subscriptionDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails': {'type': 'obj', 'data': [{'name': 'connectorType', 'description': 'Connector Type (Must be EMAIL)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fromEmailAddress', 'description': 'Senders Email Address\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'toEmailAddresses', 'description': "Recipient's Email Addresses (Comma separated)\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'subject', 'description': 'Email Subject\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestItemEventManagementUpdateEmailEventSubscriptionFilter': {'type': 'obj', 'data': [{'name': 'eventIds', 'description': 'Event Ids (Comma separated event ids)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}]], 'flatten_structure_key': [['RequestEventManagementCreateEmailEventSubscription', 'RequestEventManagementUpdateEmailEventSubscription']], 'access_list': [[[], []]]}}}
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateEmailEventSubscription`,
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
										Description: `(From Get Email Subscription Details --> pick InstanceId)
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
													Description: `Connector Type (Must be EMAIL)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"from_email_address": &schema.Schema{
													Description: `Senders Email Address
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"subject": &schema.Schema{
													Description: `Email Subject
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"to_email_addresses": &schema.Schema{
													Description: `Recipient's Email Addresses (Comma separated)
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

func resourceEventSubscriptionEmailCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEventSubscriptionEmailCreateEmailEventSubscription(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.EventManagement.CreateEmailEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEmailEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEmailEventSubscription", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceEventSubscriptionEmailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetEmailEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetEmailEventSubscriptionsQueryParams{}

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

		response1, restyResp1, err := client.EventManagement.GetEmailEventSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEmailEventSubscriptions", err,
				"Failure at GetEmailEventSubscriptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceEventSubscriptionEmailUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		request1 := expandRequestEventSubscriptionEmailUpdateEmailEventSubscription(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EventManagement.UpdateEmailEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEmailEventSubscription", err, restyResp1.String(),
					"Failure at UpdateEmailEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEmailEventSubscription", err,
				"Failure at UpdateEmailEventSubscription, unexpected response", ""))
			return diags
		}
	}

	return resourceEventSubscriptionEmailRead(ctx, d, m)
}

func resourceEventSubscriptionEmailDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete EventSubscriptionEmail on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestEventSubscriptionEmailCreateEmailEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscription {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitle {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitle{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitle {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitle{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionType {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionType{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionType {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionType{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItems {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItems{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItems {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItems{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschema {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschema{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschema {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschema{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaFilter {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscriptionschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscription {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".title")))) {
		request.Title = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleArray(ctx, key+".title", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeArray(ctx, key+".type", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".$schema")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".$schema")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".$schema")))) {
		request.Schema = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaArray(ctx, key+".$schema", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitle {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitle{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitle(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitle(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitle {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitle{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTitleFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTitleFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionType {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionType{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionType(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionType(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionType {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionType{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionTypeFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionTypeFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItems {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItems{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItems {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItems{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemsFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionItemsFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschema {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschema{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchema(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchema(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschema {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschema{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpoints {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionSchemaFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaFilter {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscriptionschemaFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
