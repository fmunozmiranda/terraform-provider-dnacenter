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
						"offset": &schema.Schema{
							Description: `offset`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"limit": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": &schema.Schema{
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
	log.Printf("[DEBUG] Beginning ApplicationSets Create")
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters")) TO DO, veririficar que el objeto fue creado realmente, para consultarlo y sacar su id...

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
	vID, okID := resourceMap["id"]

	selectedMethod := 1
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

	queryParams2 := dnacentersdkgo.DeleteApplicationSetQueryParams{}

	if okID {
		queryParams2.ID = vID
	}

	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		item1, err := searchApplicationPolicyGetApplicationSets(m, queryParams1)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationSet(&queryParams2)
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

	//TODO

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

func searchApplicationPolicyGetApplicationSets(m interface{}, queryParams dnacentersdkgo.GetApplicationSetsQueryParams) (*dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSets
	ite, _, err = client.ApplicationPolicy.GetApplicationSets(&queryParams)

	if ite == nil {
		return foundItem, err
	}

	if ite.Response == nil {
		return foundItem, err
	}

	items := ite.Response
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
