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

func resourceBusinessSdaHostonboardingSSIDIPpool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Fabric Wireless.

- Add SSID to IP Pool Mapping.

- Update SSID to IP Pool Mapping.
`,

		CreateContext: resourceBusinessSdaHostonboardingSSIDIPpoolCreate,
		ReadContext:   resourceBusinessSdaHostonboardingSSIDIPpoolRead,
		UpdateContext: resourceBusinessSdaHostonboardingSSIDIPpoolUpdate,
		DeleteContext: resourceBusinessSdaHostonboardingSSIDIPpoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssid_names": &schema.Schema{
							Description: `List of SSIDs
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vlan_name": &schema.Schema{
							Description: `VLAN Name
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

func resourceBusinessSdaHostonboardingSSIDIPpoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestBusinessSdaHostonboardingSSIDIPpoolAddSSIDToIPPoolMapping(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.FabricWireless.AddSSIDToIPPoolMapping(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddSSIDToIPPoolMapping", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddSSIDToIPPoolMapping", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVLANName, okVLANName := resourceMap["vlan_name"]
	vSiteNameHierarchy, okSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSSIDToIPPoolMapping")
		queryParams1 := dnacentersdkgo.GetSSIDToIPPoolMappingQueryParams{}

		queryParams1.VLANName = vVLANName

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		response1, restyResp1, err := client.FabricWireless.GetSSIDToIPPoolMapping(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSSIDToIPPoolMapping", err,
				"Failure at GetSSIDToIPPoolMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO

	}
	return diags
}

func resourceBusinessSdaHostonboardingSSIDIPpoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVLANName, okVLANName := resourceMap["vlan_name"]
	vSiteNameHierarchy, okSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestBusinessSdaHostonboardingSSIDIPpoolUpdateSSIDToIPPoolMapping(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.FabricWireless.UpdateSSIDToIPPoolMapping(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSSIDToIPPoolMapping", err, restyResp1.String(),
					"Failure at UpdateSSIDToIPPoolMapping, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSSIDToIPPoolMapping", err,
				"Failure at UpdateSSIDToIPPoolMapping, unexpected response", ""))
			return diags
		}
	}

	return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
}

func resourceBusinessSdaHostonboardingSSIDIPpoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete BusinessSdaHostonboardingSSIDIPpool on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestBusinessSdaHostonboardingSSIDIPpoolAddSSIDToIPPoolMapping(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessAddSSIDToIPPoolMapping {
	request := dnacentersdkgo.RequestFabricWirelessAddSSIDToIPPoolMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_names")))) {
		request.SSIDNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestBusinessSdaHostonboardingSSIDIPpoolUpdateSSIDToIPPoolMapping(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessUpdateSSIDToIPPoolMapping {
	request := dnacentersdkgo.RequestFabricWirelessUpdateSSIDToIPPoolMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_name")))) {
		request.ScalableGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_names")))) {
		request.SSIDNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
