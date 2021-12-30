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
      "parameters": &schema.Schema{
        Description: `Array of RequestEventManagementCreateEmailEventSubscription`,
        Type:        schema.TypeList,
        Optional:    true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
          
            "description": &schema.Schema{
              Description: `Description
`,
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
                    Description: `Event Ids (Comma separated event ids)
`,
                    Type:        schema.TypeList,
                    Optional:    true,
                    Elem:        &schema.Schema{
                      Type:      schema.TypeString,
                    },
                  },
                },
              },
            },
            "name": &schema.Schema{
              Description: `Name
`,
              Type:        schema.TypeString,
              Optional:    true,
            },
            "subscription_endpoints": &schema.Schema{
              Type:     schema.TypeList,
              Optional: true,
              Elem: &schema.Resource{
                Schema: map[string]*schema.Schema{
                
                  "instance_id": &schema.Schema{
                    Description: `(From Get Email Subscription Details --> pick InstanceId)
`,
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
                          Description: `Connector Type (Must be EMAIL)
`,
                          Type:        schema.TypeString,
                          Optional:    true,
                        },
                        "from_email_address": &schema.Schema{
                          Description: `Senders Email Address
`,
                          Type:        schema.TypeString,
                          Optional:    true,
                        },
                        "subject": &schema.Schema{
                          Description: `Email Subject
`,
                          Type:        schema.TypeString,
                          Optional:    true,
                        },
                        "to_email_addresses": &schema.Schema{
                          Description: `Recipient's Email Addresses (Comma separated)
`,
                          Type:        schema.TypeList,
                          Optional:    true,
                          Elem:        &schema.Schema{
                            Type:      schema.TypeString,
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
              Type:        schema.TypeString,
              Optional:    true,
            },
            "version": &schema.Schema{
              Description: `Version
`,
              Type:        schema.TypeString,
              Optional:    true,
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
			return resourceEventSubscriptionEmailRead(ctx, d, m)
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
	if v := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription{}
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
			i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
			if i != nil {
				request = append(request, *i)
			}
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints{}
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
			i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
			if i != nil {
				request = append(request, *i)
			}
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
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


func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilter{}
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
	if v := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription{}
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
			i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
			if i != nil {
				request = append(request, *i)
			}
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints{}
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
			i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
			if i != nil {
				request = append(request, *i)
			}
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
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


func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}




func searchEventManagementGetEmailEventSubscriptions(m interface{}, items []dnacentersdkgo.ResponseEventManagementGetEmailEventSubscriptions, name string, id string) (, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem 
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *dnacentersdkgo.ResponseEventManagement
			getItem, _, err = client.EventManagement.(id,name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "")
			}
			foundItem = getItem
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *dnacentersdkgo.ResponseEventManagement
			getItem, _, err = client.EventManagement.(id,name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "")
			}
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}