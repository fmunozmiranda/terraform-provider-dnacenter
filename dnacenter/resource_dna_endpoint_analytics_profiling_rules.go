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

func resourceEndpointAnalyticsProfilingRules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Policy.

- Creates profiling rule from the request body.

- Updates the profiling rule for the given 'ruleId'.

- Deletes the profiling rule for the given 'ruleId'.
`,

		CreateContext: resourceEndpointAnalyticsProfilingRulesCreate,
		ReadContext:   resourceEndpointAnalyticsProfilingRulesRead,
		UpdateContext: resourceEndpointAnalyticsProfilingRulesUpdate,
		DeleteContext: resourceEndpointAnalyticsProfilingRulesDelete,
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

						"cluster_id": &schema.Schema{
							Description: `Unique identifier for ML cluster. Only applicable for 'ML Rule'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"condition_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_dictionary": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"operator": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"condition_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_deleted": &schema.Schema{
							Description: `Flag to indicate whether the rule was deleted.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_by": &schema.Schema{
							Description: `User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_on": &schema.Schema{
							Description: `Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"plugin_id": &schema.Schema{
							Description: `Plugin for the rule. Only applicable for 'Cisco Default' rules.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rejected": &schema.Schema{
							Description: `Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"result": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_type": &schema.Schema{
										Description: `List of device types determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_manufacturer": &schema.Schema{
										Description: `List of hardware manufacturers determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_model": &schema.Schema{
										Description: `List of hardware models determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operating_system": &schema.Schema{
										Description: `List of operating systems determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"rule_id": &schema.Schema{
							Description: `Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_name": &schema.Schema{
							Description: `Human readable name for the rule.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_priority": &schema.Schema{
							Description: `Priority for the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rule_type": &schema.Schema{
							Description: `Type of the rule. Allowed values are 'Cisco Default - Static', 'Cisco Default - Dynamic', 'Custom Rule', 'ML Rule'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_version": &schema.Schema{
							Description: `Version of the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"source_priority": &schema.Schema{
							Description: `Source priority for the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_attributes": &schema.Schema{
							Description: `List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cluster_id": &schema.Schema{
							Description: `Unique identifier for ML cluster. Only applicable for 'ML Rule'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"condition_groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"attribute_dictionary": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"operator": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"condition_group": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"is_deleted": &schema.Schema{
							Description: `Flag to indicate whether the rule was deleted.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"last_modified_by": &schema.Schema{
							Description: `User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_modified_on": &schema.Schema{
							Description: `Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"plugin_id": &schema.Schema{
							Description: `Plugin for the rule. Only applicable for 'Cisco Default' rules.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rejected": &schema.Schema{
							Description: `Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"result": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_type": &schema.Schema{
										Description: `List of device types determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_manufacturer": &schema.Schema{
										Description: `List of hardware manufacturers determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_model": &schema.Schema{
										Description: `List of hardware models determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operating_system": &schema.Schema{
										Description: `List of operating systems determined by the current rule.
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
						"rule_id": &schema.Schema{
							Description: `Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule_name": &schema.Schema{
							Description: `Human readable name for the rule.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule_priority": &schema.Schema{
							Description: `Priority for the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rule_type": &schema.Schema{
							Description: `Type of the rule. Allowed values are 'Cisco Default - Static', 'Cisco Default - Dynamic', 'Custom Rule', 'ML Rule'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule_version": &schema.Schema{
							Description: `Version of the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source_priority": &schema.Schema{
							Description: `Source priority for the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"used_attributes": &schema.Schema{
							Description: `List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
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
	}
}

func resourceEndpointAnalyticsProfilingRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRule(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vRuleID, okRuleID := resourceItem["rule_id"]
	vvRuleID := interfaceToString(vRuleID)
	if okRuleID && vvRuleID != "" {
		getResponse2, _, err := client.Policy.GetDetailsOfASingleProfilingRule(vvRuleID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["rule_id"] = vvRuleID
			d.SetId(joinResourceID(resourceMap))
			return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.Policy.GetListOfProfilingRules(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsPolicyGetListOfProfilingRules(m, response2, nil)
			item2, err := searchPolicyGetListOfProfilingRules(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["rule_id"] = vvRuleID
				d.SetId(joinResourceID(resourceMap))
				return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.Policy.CreateAProfilingRule(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAProfilingRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAProfilingRule", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["rule_id"] = vvRuleID
	d.SetId(joinResourceID(resourceMap))
	return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
}

func resourceEndpointAnalyticsProfilingRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vRuleType, okRuleType := resourceMap["rule_type"]
	vIncludeDeleted, okIncludeDeleted := resourceMap["include_deleted"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]
	vRuleID, okRuleID := resourceMap["rule_id"]

	method1 := []bool{okRuleType, okIncludeDeleted, okLimit, okOffset, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okRuleID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetListOfProfilingRules")
		queryParams1 := dnacentersdkgo.GetListOfProfilingRulesQueryParams{}

		if okRuleType {
			queryParams1.RuleType = vRuleType
		}
		if okIncludeDeleted {
			queryParams1.IncludeDeleted = *stringToBooleanPtr(vIncludeDeleted)
		}
		if okLimit {
			queryParams1.Limit = *stringToFloat64Ptr(vLimit)
		}
		if okOffset {
			queryParams1.Offset = *stringToFloat64Ptr(vOffset)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy
		}
		if okOrder {
			queryParams1.Order = vOrder
		}

		response1, restyResp1, err := client.Policy.GetListOfProfilingRules(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetListOfProfilingRules", err,
				"Failure at GetListOfProfilingRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDetailsOfASingleProfilingRule")
		vvRuleID := vRuleID

		response2, restyResp2, err := client.Policy.GetDetailsOfASingleProfilingRule(vvRuleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDetailsOfASingleProfilingRule", err,
				"Failure at GetDetailsOfASingleProfilingRule, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPolicyGetDetailsOfASingleProfilingRuleItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDetailsOfASingleProfilingRule response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceEndpointAnalyticsProfilingRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vRuleType, okRuleType := resourceMap["rule_type"]
	vIncludeDeleted, okIncludeDeleted := resourceMap["include_deleted"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]
	vRuleID, okRuleID := resourceMap["rule_id"]

	method1 := []bool{okRuleType, okIncludeDeleted, okLimit, okOffset, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okRuleID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRule(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		restyResp1, err := client.Policy.UpdateAnExistingProfilingRule(vvRuleID, request1)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAnExistingProfilingRule", err, restyResp1.String(),
					"Failure at UpdateAnExistingProfilingRule, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAnExistingProfilingRule", err,
				"Failure at UpdateAnExistingProfilingRule, unexpected response", ""))
			return diags
		}
	}

	return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
}

func resourceEndpointAnalyticsProfilingRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vRuleType, okRuleType := resourceMap["rule_type"]
	vIncludeDeleted, okIncludeDeleted := resourceMap["include_deleted"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vOrder, okOrder := resourceMap["order"]
	vRuleID, okRuleID := resourceMap["rule_id"]

	method1 := []bool{okRuleType, okIncludeDeleted, okLimit, okOffset, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okRuleID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Policy.GetListOfProfilingRules(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsPolicyGetListOfProfilingRules(m, getResp1, nil)
		item1, err := searchPolicyGetListOfProfilingRules(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.Policy.GetDetailsOfASingleProfilingRule(vvRuleID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.Policy.DeleteAnExistingProfilingRule(vvRuleID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAnExistingProfilingRule", err, restyResp1.String(),
				"Failure at DeleteAnExistingProfilingRule, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAnExistingProfilingRule", err,
			"Failure at DeleteAnExistingProfilingRule, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyCreateAProfilingRule {
	request := dnacentersdkgo.RequestPolicyCreateAProfilingRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_id")))) {
		request.RuleID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_name")))) {
		request.RuleName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_type")))) {
		request.RuleType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_version")))) {
		request.RuleVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_priority")))) {
		request.RulePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_priority")))) {
		request.SourcePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_deleted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_deleted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_deleted")))) {
		request.IsDeleted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_by")))) {
		request.LastModifiedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_on")))) {
		request.LastModifiedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plugin_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plugin_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plugin_id")))) {
		request.PluginID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cluster_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cluster_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cluster_id")))) {
		request.ClusterID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rejected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rejected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rejected")))) {
		request.Rejected = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".result")))) {
		request.Result = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleResult(ctx, key+".result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_groups")))) {
		request.ConditionGroups = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroups(ctx, key+".condition_groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".used_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".used_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".used_attributes")))) {
		request.UsedAttributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleResult(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyCreateAProfilingRuleResult {
	request := dnacentersdkgo.RequestPolicyCreateAProfilingRuleResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operating_system")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operating_system")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operating_system")))) {
		request.OperatingSystem = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroups {
	request := dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_group")))) {
		request.ConditionGroup = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsConditionGroupArray(ctx, key+".condition_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsCondition(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsCondition {
	request := dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute")))) {
		request.Attribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_dictionary")))) {
		request.AttributeDictionary = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsConditionGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsConditionGroup {
	request := []dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsConditionGroup{}
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
		i := expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsConditionGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsConditionGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsConditionGroup {
	var request dnacentersdkgo.RequestPolicyCreateAProfilingRuleConditionGroupsConditionGroup
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRule {
	request := dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_id")))) {
		request.RuleID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_name")))) {
		request.RuleName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_type")))) {
		request.RuleType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_version")))) {
		request.RuleVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_priority")))) {
		request.RulePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_priority")))) {
		request.SourcePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_deleted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_deleted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_deleted")))) {
		request.IsDeleted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_by")))) {
		request.LastModifiedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_on")))) {
		request.LastModifiedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plugin_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plugin_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plugin_id")))) {
		request.PluginID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cluster_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cluster_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cluster_id")))) {
		request.ClusterID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rejected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rejected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rejected")))) {
		request.Rejected = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".result")))) {
		request.Result = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleResult(ctx, key+".result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_groups")))) {
		request.ConditionGroups = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroups(ctx, key+".condition_groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".used_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".used_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".used_attributes")))) {
		request.UsedAttributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleResult(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleResult {
	request := dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operating_system")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operating_system")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operating_system")))) {
		request.OperatingSystem = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroups {
	request := dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_group")))) {
		request.ConditionGroup = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsConditionGroupArray(ctx, key+".condition_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsCondition(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsCondition {
	request := dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute")))) {
		request.Attribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_dictionary")))) {
		request.AttributeDictionary = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsConditionGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsConditionGroup {
	request := []dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsConditionGroup{}
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
		i := expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsConditionGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsConditionGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsConditionGroup {
	var request dnacentersdkgo.RequestPolicyUpdateAnExistingProfilingRuleConditionGroupsConditionGroup
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchPolicyGetListOfProfilingRules(m interface{}, items []dnacentersdkgo.ResponsePolicyGetListOfProfilingRulesProfilingRules, name string, id string) (*dnacentersdkgo.ResponsePolicyGetDetailsOfASingleProfilingRule, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponsePolicyGetDetailsOfASingleProfilingRule
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *dnacentersdkgo.ResponsePolicyGetDetailsOfASingleProfilingRule
			getItem, _, err = client.Policy.GetDetailsOfASingleProfilingRule(id, name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDetailsOfASingleProfilingRule")
			}
			foundItem = getItem
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *dnacentersdkgo.ResponsePolicyGetDetailsOfASingleProfilingRule
			getItem, _, err = client.Policy.GetDetailsOfASingleProfilingRule(id, name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDetailsOfASingleProfilingRule")
			}
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}