package dnacenter

import (
	"context"
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
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vVlan_name := resourceItem["vlan_name"]
	vSite_name_hierarchy := resourceItem["site_name_hierarchy"]
	vvVlan_name := interfaceToString(vVlan_name)
	vvSite_name_hierarchy := interfaceToString(vSite_name_hierarchy)

	queryParams1 := dnacentersdkgo.GetSSIDToIPPoolMappingQueryParams{}

	queryParams1.VLANName = vvVlan_name
	queryParams1.SiteNameHierarchy = vvSite_name_hierarchy
	getResponse1, _, err := client.FabricWireless.GetSSIDToIPPoolMapping(&queryParams1)

	if err == nil && getResponse1 != nil {
		resourceMap := make(map[string]string)
		resourceMap["vlan_name"] = vvVlan_name
		resourceMap["site_name_hierarchy"] = vvSite_name_hierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
	}
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
	resourceMap["vlan_name"] = vvVlan_name
	resourceMap["site_name_hierarchy"] = vvSite_name_hierarchy
	return resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx, d, m)
}

func resourceBusinessSdaHostonboardingSSIDIPpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVLANName := resourceMap["vlan_name"]
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

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

		vItem1 := flattenFabricWirelessGetSSIDToIPPoolMappingItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDToIPPoolMapping response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceBusinessSdaHostonboardingSSIDIPpoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVLANName := resourceMap["vlan_name"]
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	//selectedMethod := 1
	queryParams1 := dnacentersdkgo.GetSSIDToIPPoolMappingQueryParams{}
	queryParams1.VLANName = vVLANName
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy

	getResp, _, err := client.FabricWireless.GetSSIDToIPPoolMapping(&queryParams1)
	if err != nil || getResp == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetAllowedProtocolByName", err,
			"Failure at GetAllowedProtocolByName, unexpected response", ""))
		return diags
	}
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", queryParams1)
		request1 := expandRequestBusinessSdaHostonboardingSSIDIPpoolUpdateSSIDToIPPoolMapping(ctx, "parameters.0", d)
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