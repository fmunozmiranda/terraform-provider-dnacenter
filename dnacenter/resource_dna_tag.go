package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Tag.

- Updates a tag specified by id

- Creates tag with specified tag attributes

- Deletes a tag specified by id
`,

		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'systemTag': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'dynamicRules': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'memberType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'rules': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'values': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'items': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {}}}, 'operation': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'value': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'systemTag': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dynamicRules': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'memberType': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'rules': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'values': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'items': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {}}}, 'operation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'value': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}, 'metadata': {'item': {'operation_id': [['CreateTag', 'UpdateTag'], 'GetTagById'], 'new_flat_structure': [[{'RequestTagCreateTag': {'type': 'obj', 'data': [{'name': 'systemTag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dynamicRules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestTagCreateTagDynamicRules'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagCreateTagDynamicRules': {'type': 'obj', 'data': [{'name': 'memberType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestTagCreateTagDynamicRulesRules'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagCreateTagDynamicRulesRules': {'type': 'obj', 'data': [{'name': 'values', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'items', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestTagCreateTagDynamicRulesRulesItems'}, {'name': 'operation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagCreateTagDynamicRulesRulesItems': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestTagUpdateTag': {'type': 'obj', 'data': [{'name': 'systemTag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dynamicRules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestTagUpdateTagDynamicRules'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagUpdateTagDynamicRules': {'type': 'obj', 'data': [{'name': 'memberType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestTagUpdateTagDynamicRulesRules'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagUpdateTagDynamicRulesRules': {'type': 'obj', 'data': [{'name': 'values', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'items', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestTagUpdateTagDynamicRulesRulesItems'}, {'name': 'operation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestTagUpdateTagDynamicRulesRulesItems': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], {'ResponseTagGetTagByIdResponse': {'type': 'obj', 'data': [{'name': 'systemTag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dynamicRules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseTagGetTagByIdResponseDynamicRules'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIdResponseDynamicRules': {'type': 'obj', 'data': [{'name': 'memberType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseTagGetTagByIdResponseDynamicRulesRules'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIdResponseDynamicRulesRules': {'type': 'obj', 'data': [{'name': 'values', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'items', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseTagGetTagByIdResponseDynamicRulesRulesItems'}, {'name': 'operation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIdResponseDynamicRulesRulesItems': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIDResponse': {'type': 'obj', 'data': [{'name': 'systemTag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dynamicRules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseTagGetTagByIdResponseDynamicRules'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIDResponseDynamicRules': {'type': 'obj', 'data': [{'name': 'memberType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rules', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseTagGetTagByIdResponseDynamicRulesRules'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIDResponseDynamicRulesRules': {'type': 'obj', 'data': [{'name': 'values', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'items', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseTagGetTagByIdResponseDynamicRulesRulesItems'}, {'name': 'operation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseTagGetTagByIDResponseDynamicRulesRulesItems': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': [['RequestTagCreateTag', 'RequestTagUpdateTag'], 'ResponseTagGetTagByIdResponse'], 'access_list': [[[], []], ['response']]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
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
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"values": &schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTagCreateTag(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.Tag.GetTagByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.Tag.GetTag(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsTagGetTag(m, response2, nil)
			item2, err := searchTagGetTag(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.Tag.CreateTag(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTag", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTag", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vAdditionalInfonameSpace, okAdditionalInfonameSpace := resourceMap["additional_info_name_space"]
	vAdditionalInfoattributes, okAdditionalInfoattributes := resourceMap["additional_info_attributes"]
	vLevel, okLevel := resourceMap["level"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vSize, okSize := resourceMap["size"]
	vField, okField := resourceMap["field"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]
	vSystemTag, okSystemTag := resourceMap["system_tag"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName, okAdditionalInfonameSpace, okAdditionalInfoattributes, okLevel, okOffset, okLimit, okSize, okField, okSortBy, okOrder, okSystemTag}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTag")
		queryParams1 := dnacentersdkgo.GetTagQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okAdditionalInfonameSpace {
			queryParams1.AdditionalInfonameSpace = vAdditionalInfonameSpace.(string)
		}
		if okAdditionalInfoattributes {
			queryParams1.AdditionalInfoattributes = vAdditionalInfoattributes.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okSize {
			queryParams1.Size = vSize.(string)
		}
		if okField {
			queryParams1.Field = vField.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSystemTag {
			queryParams1.SystemTag = vSystemTag.(string)
		}

		response1, restyResp1, err := client.Tag.GetTag(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTag", err,
				"Failure at GetTag, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetTagByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Tag.GetTagByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vAdditionalInfonameSpace, okAdditionalInfonameSpace := resourceMap["additional_info_name_space"]
	vAdditionalInfoattributes, okAdditionalInfoattributes := resourceMap["additional_info_attributes"]
	vLevel, okLevel := resourceMap["level"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vSize, okSize := resourceMap["size"]
	vField, okField := resourceMap["field"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]
	vSystemTag, okSystemTag := resourceMap["system_tag"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName, okAdditionalInfonameSpace, okAdditionalInfoattributes, okLevel, okOffset, okLimit, okSize, okField, okSortBy, okOrder, okSystemTag}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.Tag.GetTagByID(vvID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagByID", err,
				"Failure at GetTagByID, unexpected response", ""))
			return diags
		}
		//Set value vvName = getResp.
		if getResp != nil {
			vvName = getResp.Name
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestTagUpdateTag(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Tag.UpdateTag(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTag", err, restyResp1.String(),
					"Failure at UpdateTag, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTag", err,
				"Failure at UpdateTag, unexpected response", ""))
			return diags
		}
	}

	return resourceTagRead(ctx, d, m)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestTagCreateTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTag {
	request := dnacentersdkgo.RequestTagCreateTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_tag")))) {
		request.SystemTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_rules")))) {
		request.DynamicRules = expandRequestTagCreateTagDynamicRulesArray(ctx, key+".dynamic_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagCreateTagDynamicRules {
	request := []dnacentersdkgo.RequestTagCreateTagDynamicRules{}
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
		i := expandRequestTagCreateTagDynamicRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRules {
	request := dnacentersdkgo.RequestTagCreateTagDynamicRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestTagCreateTagDynamicRulesRules(ctx, key+".rules.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRulesRules {
	request := dnacentersdkgo.RequestTagCreateTagDynamicRulesRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".values")))) {
		request.Values = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestTagCreateTagDynamicRulesRulesItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRulesItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems {
	request := []dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems{}
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
		i := expandRequestTagCreateTagDynamicRulesRulesItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagCreateTagDynamicRulesRulesItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems {
	var request dnacentersdkgo.RequestTagCreateTagDynamicRulesRulesItems
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTag {
	request := dnacentersdkgo.RequestTagUpdateTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_tag")))) {
		request.SystemTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dynamic_rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dynamic_rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dynamic_rules")))) {
		request.DynamicRules = expandRequestTagUpdateTagDynamicRulesArray(ctx, key+".dynamic_rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagUpdateTagDynamicRules {
	request := []dnacentersdkgo.RequestTagUpdateTagDynamicRules{}
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
		i := expandRequestTagUpdateTagDynamicRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRules {
	request := dnacentersdkgo.RequestTagUpdateTagDynamicRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_type")))) {
		request.MemberType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestTagUpdateTagDynamicRulesRules(ctx, key+".rules.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRulesRules {
	request := dnacentersdkgo.RequestTagUpdateTagDynamicRulesRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".values")))) {
		request.Values = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".items")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".items")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".items")))) {
		request.Items = expandRequestTagUpdateTagDynamicRulesRulesItemsArray(ctx, key+".items", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRulesItemsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems {
	request := []dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems{}
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
		i := expandRequestTagUpdateTagDynamicRulesRulesItems(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTagUpdateTagDynamicRulesRulesItems(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems {
	var request dnacentersdkgo.RequestTagUpdateTagDynamicRulesRulesItems
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
