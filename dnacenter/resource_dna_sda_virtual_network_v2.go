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

func resourceSdaVirtualNetworkV2() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Add virtual network with scalable groups at global level

- Delete virtual network with scalable groups

- Update virtual network with scalable groups
`,

		CreateContext: resourceSdaVirtualNetworkV2Create,
		ReadContext:   resourceSdaVirtualNetworkV2Read,
		UpdateContext: resourceSdaVirtualNetworkV2Update,
		DeleteContext: resourceSdaVirtualNetworkV2Delete,
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

						"is_guest_virtual_network": &schema.Schema{
							Description: `To create guest virtual network
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"scalable_group_names": &schema.Schema{
							Description: `Scalable Group to be associated to virtual network
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name to be assigned  global level
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network_type": &schema.Schema{
							Description: `Virtual Network Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaVirtualNetworkV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkV2AddVirtualNetworkWithScalableGroups(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddVirtualNetworkWithScalableGroups(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddVirtualNetworkWithScalableGroups", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddVirtualNetworkWithScalableGroups", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaVirtualNetworkV2Read(ctx, d, m)
}

func resourceSdaVirtualNetworkV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName, okVirtualNetworkName := resourceMap["virtual_network_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVirtualNetworkWithScalableGroups")
		queryParams1 := dnacentersdkgo.GetVirtualNetworkWithScalableGroupsQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName

		response1, restyResp1, err := client.Sda.GetVirtualNetworkWithScalableGroups(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualNetworkWithScalableGroups", err,
				"Failure at GetVirtualNetworkWithScalableGroups, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetVirtualNetworkWithScalableGroupsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkWithScalableGroups response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaVirtualNetworkV2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName, okVirtualNetworkName := resourceMap["virtual_network_name"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestSdaVirtualNetworkV2UpdateVirtualNetworkWithScalableGroups(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.UpdateVirtualNetworkWithScalableGroups(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateVirtualNetworkWithScalableGroups", err, restyResp1.String(),
					"Failure at UpdateVirtualNetworkWithScalableGroups, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateVirtualNetworkWithScalableGroups", err,
				"Failure at UpdateVirtualNetworkWithScalableGroups, unexpected response", ""))
			return diags
		}
	}

	return resourceSdaVirtualNetworkV2Read(ctx, d, m)
}

func resourceSdaVirtualNetworkV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName, okVirtualNetworkName := resourceMap["virtual_network_name"]

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Sda.GetVirtualNetworkWithScalableGroups(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSdaGetVirtualNetworkWithScalableGroups(m, getResp1, nil)
		item1, err := searchSdaGetVirtualNetworkWithScalableGroups(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Sda.DeleteVirtualNetworkWithScalableGroups()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteVirtualNetworkWithScalableGroups", err, restyResp1.String(),
				"Failure at DeleteVirtualNetworkWithScalableGroups, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteVirtualNetworkWithScalableGroups", err,
			"Failure at DeleteVirtualNetworkWithScalableGroups, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaVirtualNetworkV2AddVirtualNetworkWithScalableGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddVirtualNetworkWithScalableGroups {
	request := dnacentersdkgo.RequestSdaAddVirtualNetworkWithScalableGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_guest_virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) {
		request.IsGuestVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_names")))) {
		request.ScalableGroupNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_type")))) {
		request.VirtualNetworkType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaVirtualNetworkV2UpdateVirtualNetworkWithScalableGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateVirtualNetworkWithScalableGroups {
	request := dnacentersdkgo.RequestSdaUpdateVirtualNetworkWithScalableGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_type")))) {
		request.VirtualNetworkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_guest_virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) {
		request.IsGuestVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_names")))) {
		request.ScalableGroupNames = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
