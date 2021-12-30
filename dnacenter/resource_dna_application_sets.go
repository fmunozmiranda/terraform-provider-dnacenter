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

func resourceApplicationSets() *schema.Resource {
  return &schema.Resource{
    Description: `It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s
`,

    CreateContext: resourceApplicationSetsCreate,
    ReadContext:   resourceApplicationSetsRead,
    UpdateContext: resourceApplicationSetsUpdate,
    DeleteContext: resourceApplicationSetsDelete,
    Importer: &schema.ResourceImporter{
      StateContext: schema.ImportStatePassthroughContext,
    },

    Schema: map[string]*schema.Schema{
      "last_updated": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
      "parameters": &schema.Schema{
        Description: `Array of RequestApplicationPolicyCreateApplicationSet`,
        Type:        schema.TypeList,
        Optional:    true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
          
            "name": &schema.Schema{
              Description: `Name`,
              Type:        schema.TypeString,
              Optional:    true,
            },
          },
        },
      },
    },
  }
}

func resourceApplicationSetsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  client := m.(*dnacentersdkgo.Client)

  var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestApplicationSetsCreateApplicationSet(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

  
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationSet(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationSet", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
		return diags
	}
				resourceMap := make(map[string]string)
			d.SetId(joinResourceID(resourceMap))
			return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

  resourceID := d.Id()
  resourceMap := separateResourceID(resourceID)
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vName, okName := resourceMap["name"]


	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationSets")
		queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

	  if okOffset {
	    queryParams1.Offset = *stringToFloat64Ptr(vOffset)
	  }
	  if okLimit {
	    queryParams1.Limit = *stringToFloat64Ptr(vLimit)
	  }
	  if okName {
	    queryParams1.Name = vName
	  }

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationSets(&queryParams1)

	
	
		if err != nil || response1 == nil {
		  if restyResp1 != nil {
		    log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		  }
		  diags = append(diags, diagErrorWithAlt(
		    "Failure when executing GetApplicationSets", err,
		    "Failure at GetApplicationSets, unexpected response", ""))
		  return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
  return diags
}

func resourceApplicationSetsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  
  client := m.(*dnacentersdkgo.Client)

  var diags diag.Diagnostics

  resourceID := d.Id()
  resourceMap := separateResourceID(resourceID)
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vName, okName := resourceMap["name"]


	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
	
		getResp1, _, err := client.ApplicationPolicy.GetApplicationSets(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsApplicationPolicyGetApplicationSets(m, getResp1, nil)
		item1, err := searchApplicationPolicyGetApplicationSets(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationSet()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationSet", err, restyResp1.String(),
				"Failure at DeleteApplicationSet, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationSet", err,
			"Failure at DeleteApplicationSet, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

  return diags
}
func expandRequestApplicationSetsCreateApplicationSet(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet{}
	if v := expandRequestApplicationSetsCreateApplicationSetItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestApplicationSetsCreateApplicationSetItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet{}
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
			i := expandRequestApplicationSetsCreateApplicationSetItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
			if i != nil {
				request = append(request, *i)
			}
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}


func expandRequestApplicationSetsCreateApplicationSetItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
        if isEmptyValue(reflect.ValueOf(request)) {
            return nil
        }
    
	return &request
}




func searchApplicationPolicyGetApplicationSets(m interface{}, items []dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse, name string, id string) (, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem 
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *dnacentersdkgo.ResponseApplicationPolicy
			getItem, _, err = client.ApplicationPolicy.(id,name)
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
			var getItem *dnacentersdkgo.ResponseApplicationPolicy
			getItem, _, err = client.ApplicationPolicy.(id,name)
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