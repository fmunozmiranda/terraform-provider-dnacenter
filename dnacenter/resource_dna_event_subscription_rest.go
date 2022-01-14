package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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

	//resourceItem := *getResourceItem(d.Get("parameters"))
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
	return resourceEventSubscriptionRestRead(ctx, d, m)
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
			queryParams1.EventIDs = vEventIDs
		}
		if okOffset {
			queryParams1.Offset = *stringToFloat64Ptr(vOffset)
		}
		if okLimit {
			queryParams1.Limit = *stringToFloat64Ptr(vLimit)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy
		}
		if okOrder {
			queryParams1.Order = vOrder
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

		//TODO FOR DNAC

		vItem1 := flattenEventManagementGetRestWebhookEventSubscriptionsItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestWebhookEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionRestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vEventIDs := resourceMap["event_ids"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]
	vSortBy := resourceMap["sort_by"]
	vOrder := resourceMap["order"]

	queryParams1 := dnacentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}
	queryParams1.EventIDs = vEventIDs
	queryParams1.Offset = *stringToFloat64Ptr(vOffset)
	queryParams1.Limit = *stringToFloat64Ptr(vLimit)
	queryParams1.SortBy = vSortBy
	queryParams1.Order = vOrder
	item, err := searchEventManagementGetRestWebhookEventSubscriptions(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetRestWebhookEventSubscriptions", err,
			"Failure at GetRestWebhookEventSubscriptions, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscription(ctx, "parameters.0", d)
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
	if v := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter{}
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
	if v := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchEventManagementGetRestWebhookEventSubscriptions(m interface{}, queryParams dnacentersdkgo.GetRestWebhookEventSubscriptionsQueryParams) (*dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptions, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptions
	var ite *dnacentersdkgo.ResponseEventManagementGetRestWebhookEventSubscriptions
	ite, _, err = client.EventManagement.GetRestWebhookEventSubscriptions(&queryParams)
	if err != nil {
		return foundItem, err
	}
	if ite == nil {
		return foundItem, err
	}

	items := ite
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.SubscriptionID == queryParams.EventIDs {
			var getItem *dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptions
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
