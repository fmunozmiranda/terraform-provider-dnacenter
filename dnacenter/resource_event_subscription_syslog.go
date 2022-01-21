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

func resourceEventSubscriptionSyslog() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Update Syslog Subscription Endpoint for list of registered events

- Create Syslog Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionSyslogCreate,
		ReadContext:   resourceEventSubscriptionSyslogRead,
		UpdateContext: resourceEventSubscriptionSyslogUpdate,
		DeleteContext: resourceEventSubscriptionSyslogDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"categories": &schema.Schema{
										Description: `Categories`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"domains_subdomains": &schema.Schema{
										Description: `Domains Subdomains`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"event_ids": &schema.Schema{
										Description: `Event Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"others": &schema.Schema{
										Description: `Others`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"severities": &schema.Schema{
										Description: `Severities`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"sources": &schema.Schema{
										Description: `Sources`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"types": &schema.Schema{
										Description: `Types`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"is_private": &schema.Schema{
							Description: `Is Private`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connector_type": &schema.Schema{
										Description: `Connector Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connector_type": &schema.Schema{
													Description: `Connector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"syslog_config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"config_id": &schema.Schema{
																Description: `Config Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"description": &schema.Schema{
																Description: `Description`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"host": &schema.Schema{
																Description: `Host`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"port": &schema.Schema{
																Description: `Port`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"tenant_id": &schema.Schema{
																Description: `Tenant Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"version": &schema.Schema{
																Description: `Version`,
																Type:        schema.TypeString,
																Computed:    true,
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

						"subscription_id": &schema.Schema{
							Description: `Subscription Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateSyslogEventSubscription`,
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
										Description: `(From Get Syslog Subscription Details --> pick instanceId)
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
													Description: `Connector Type (Must be SYSLOG)
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

func resourceEventSubscriptionSyslogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscription(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.EventManagement.CreateSyslogEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSyslogEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSyslogEventSubscription", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceEventSubscriptionSyslogRead(ctx, d, m)
}

func resourceEventSubscriptionSyslogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vEventIDs, okEventIDs := resourceMap["event_ids"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSyslogEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetSyslogEventSubscriptionsQueryParams{}

		if okEventIDs {
			queryParams1.EventIDs = vEventIDs
		}

		response1, restyResp1, err := client.EventManagement.GetSyslogEventSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSyslogEventSubscriptions", err,
				"Failure at GetSyslogEventSubscriptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenEventManagementGetSyslogEventSubscriptionsItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSyslogEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionSyslogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vEventIDs := resourceMap["event_ids"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]
	vSortBy := resourceMap["sort_by"]
	vOrder := resourceMap["order"]

	queryParams1 := dnacentersdkgo.GetSyslogEventSubscriptionsQueryParams{}
	queryParams1.EventIDs = vEventIDs
	queryParams1.Offset = *stringToFloat64Ptr(vOffset)
	queryParams1.Limit = *stringToFloat64Ptr(vLimit)
	queryParams1.SortBy = vSortBy
	queryParams1.Order = vOrder
	item, err := searchEventManagementGetSyslogEventSubscriptions(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetSyslogEventSubscriptions", err,
			"Failure at GetSyslogEventSubscriptions, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscription(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EventManagement.UpdateSyslogEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSyslogEventSubscription", err, restyResp1.String(),
					"Failure at UpdateSyslogEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSyslogEventSubscription", err,
				"Failure at UpdateSyslogEventSubscription, unexpected response", ""))
			return diags
		}
	}

	return resourceEventSubscriptionSyslogRead(ctx, d, m)
}

func resourceEventSubscriptionSyslogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete EventSubscriptionSyslog on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateSyslogEventSubscription {
	request := dnacentersdkgo.RequestEventManagementCreateSyslogEventSubscription{}
	if v := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscription{}
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
		i := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateSyslogEventSubscription {
	request := dnacentersdkgo.RequestEventManagementUpdateSyslogEventSubscription{}
	if v := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscription{}
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
		i := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchEventManagementGetSyslogEventSubscriptions(m interface{}, queryParams dnacentersdkgo.GetSyslogEventSubscriptionsQueryParams) (*dnacentersdkgo.ResponseItemEventManagementGetSyslogEventSubscriptions, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemEventManagementGetSyslogEventSubscriptions
	var ite *dnacentersdkgo.ResponseEventManagementGetSyslogEventSubscriptions
	ite, _, err = client.EventManagement.GetSyslogEventSubscriptions(&queryParams)
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
		if item.Name == queryParams.EventIDs {
			var getItem *dnacentersdkgo.ResponseItemEventManagementGetSyslogEventSubscriptions
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}