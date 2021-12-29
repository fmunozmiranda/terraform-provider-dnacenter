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

func resourceNfvProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Site Design.

- API to create network profile for different NFV topologies

- API to update a NFV Network profile

- API to delete nfv network profile.
`,

		CreateContext: resourceNfvProfileCreate,
		ReadContext:   resourceNfvProfileRead,
		UpdateContext: resourceNfvProfileUpdate,
		DeleteContext: resourceNfvProfileDelete,
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

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"current_device_tag": &schema.Schema{
										Description: `Existing device tag name saved in the nfv profiles (eg: dev1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_networks": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connection_type": &schema.Schema{
													Description: `Type of network connection from custom network (eg: lan)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"network_name": &schema.Schema{
													Description: `Name of custom network (eg: cust-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"services_to_connect": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"service_name": &schema.Schema{
																Description: `Name of service to be connected to the custom network (eg: router-1)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"vlan_id": &schema.Schema{
													Description: `Vlan id for the custom network(eg: 4000)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"vlan_mode": &schema.Schema{
													Description: `Network mode (eg Access or Trunk)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"custom_template": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_type": &schema.Schema{
													Description: `Type of the device. Allowed values are 'Cisco 5400 Enterprise Network Compute System', 'Cisco Integrated Services Virtual Router', 'Cisco Adaptive Security Virtual Appliance (ASAv)', 'NFVIS', 'ASAV'.
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"template": &schema.Schema{
													Description: `Name of the template(eg NFVIS template)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"template_type": &schema.Schema{
													Description: `Name of the template type to which template is associated (eg: Cloud DayN Templates). Allowed values are 'Onboarding Template(s)' and 'Day-N-Template(s)'.
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"device_tag": &schema.Schema{
										Description: `Device Tag name(eg: dev1)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_type": &schema.Schema{
										Description: `Name of the device used in creating nfv profile. Allowed values are 'Cisco 5400 Enterprise Network Compute System', 'Cisco 5100 Enterprise Network Compute System'.
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"direct_internet_access_for_firewall": &schema.Schema{
										Description: `Direct internet access value should be boolean (eg: false or true)
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"service_provider_profile": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connect": &schema.Schema{
													Description: `Connection of service provider and device value should be boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"connect_default_gateway_on_wan": &schema.Schema{
													Description: `Connect default gateway connect value as boolean (eg: true)
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"link_type": &schema.Schema{
													Description: `Name of connection type(eg: GigabitEthernet) 
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"service_provider": &schema.Schema{
													Description: `Name of the service provider(eg: Airtel)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"services": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"firewall_mode": &schema.Schema{
													Description: `Firewall mode details example (routed, transparent)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"image_name": &schema.Schema{
													Description: `Service image name (eg: isrv-universalk9.16.12.01a.tar.gz)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"profile_type": &schema.Schema{
													Description: `Profile type of service (eg: ISRv-mini)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"service_name": &schema.Schema{
													Description: `Name of the service (eg: Router-1)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"service_type": &schema.Schema{
													Description: `Service type (eg: ISRV)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"v_nic_mapping": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"assign_ip_address_to_network": &schema.Schema{
																Description: `Assign ip address to network (eg: true or false)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"network_type": &schema.Schema{
																Description: `Type of connection (eg:  wan, lan or internal)
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
									"vlan_for_l2": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"vlan_description": &schema.Schema{
													Description: `Vlan description(eg: Access 4018)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"vlan_id": &schema.Schema{
													Description: `Vlan id (eg: 4018)
`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"vlan_type": &schema.Schema{
													Description: `Vlan type(eg: Access or Trunk)
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
						"id": &schema.Schema{
							Description: `id path parameter. Id of the NFV profile to be updated
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"profile_name": &schema.Schema{
							Description: `Name of the profile to create NFV profile
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

func resourceNfvProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNfvProfileCreateNfvProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse1, _, err := client.SiteDesign.GetNfvProfile(vvID, nil)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceNfvProfileRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.SiteDesign.CreateNfvProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNfvProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNfvProfile", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceNfvProfileRead(ctx, d, m)
}

func resourceNfvProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vName, okName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNfvProfile")
		vvID := vID
		queryParams1 := dnacentersdkgo.GetNfvProfileQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset
		}
		if okLimit {
			queryParams1.Limit = vLimit
		}
		if okName {
			queryParams1.Name = vName
		}

		response1, restyResp1, err := client.SiteDesign.GetNfvProfile(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNfvProfile", err,
				"Failure at GetNfvProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceNfvProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vName, okName := resourceMap["name"]

	selectedMethod := 1
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestNfvProfileUpdateNfvProfile(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SiteDesign.UpdateNfvProfile(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNfvProfile", err, restyResp1.String(),
					"Failure at UpdateNfvProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNfvProfile", err,
				"Failure at UpdateNfvProfile, unexpected response", ""))
			return diags
		}
	}

	return resourceNfvProfileRead(ctx, d, m)
}

func resourceNfvProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestNfvProfileCreateNfvProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfile {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestNfvProfileCreateNfvProfileDeviceArray(ctx, key+".device", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDevice {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDevice{}
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
		i := expandRequestNfvProfileCreateNfvProfileDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDevice {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_tag")))) {
		request.DeviceTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_provider_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_provider_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_provider_profile")))) {
		request.ServiceProviderProfile = expandRequestNfvProfileCreateNfvProfileDeviceServiceProviderProfileArray(ctx, key+".service_provider_profile", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direct_internet_access_for_firewall")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direct_internet_access_for_firewall")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direct_internet_access_for_firewall")))) {
		request.DirectInternetAccessForFirewall = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = expandRequestNfvProfileCreateNfvProfileDeviceServicesArray(ctx, key+".services", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_networks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_networks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_networks")))) {
		request.CustomNetworks = expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksArray(ctx, key+".custom_networks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_for_l2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_for_l2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_for_l2")))) {
		request.VLANForL2 = expandRequestNfvProfileCreateNfvProfileDeviceVLANForL2Array(ctx, key+".vlan_for_l2", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_template")))) {
		request.CustomTemplate = expandRequestNfvProfileCreateNfvProfileDeviceCustomTemplateArray(ctx, key+".custom_template", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServiceProviderProfileArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceServiceProviderProfile(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServiceProviderProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_provider")))) {
		request.ServiceProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link_type")))) {
		request.LinkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connect")))) {
		request.Connect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connect_default_gateway_on_wan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connect_default_gateway_on_wan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connect_default_gateway_on_wan")))) {
		request.ConnectDefaultGatewayOnWan = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServices {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServices{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceServices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServices {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_type")))) {
		request.ServiceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_type")))) {
		request.ProfileType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_name")))) {
		request.ServiceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_name")))) {
		request.ImageName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".v_nic_mapping")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".v_nic_mapping")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".v_nic_mapping")))) {
		request.VNicMapping = expandRequestNfvProfileCreateNfvProfileDeviceServicesVNicMappingArray(ctx, key+".v_nic_mapping", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".firewall_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".firewall_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".firewall_mode")))) {
		request.FirewallMode = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServicesVNicMappingArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceServicesVNicMapping(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceServicesVNicMapping(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_type")))) {
		request.NetworkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_ip_address_to_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_ip_address_to_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_ip_address_to_network")))) {
		request.AssignIPAddressToNetwork = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworks {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworks{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworks {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_name")))) {
		request.NetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services_to_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services_to_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services_to_connect")))) {
		request.ServicesToConnect = expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksServicesToConnectArray(ctx, key+".services_to_connect", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_type")))) {
		request.ConnectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_mode")))) {
		request.VLANMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksServicesToConnectArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksServicesToConnect(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomNetworksServicesToConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_name")))) {
		request.ServiceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceVLANForL2Array(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceVLANForL2 {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceVLANForL2{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceVLANForL2(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceVLANForL2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceVLANForL2 {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceVLANForL2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_type")))) {
		request.VLANType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_description")))) {
		request.VLANDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomTemplateArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomTemplate {
	request := []dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomTemplate{}
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
		i := expandRequestNfvProfileCreateNfvProfileDeviceCustomTemplate(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileCreateNfvProfileDeviceCustomTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomTemplate {
	request := dnacentersdkgo.RequestSiteDesignCreateNfvProfileDeviceCustomTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template")))) {
		request.Template = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_type")))) {
		request.TemplateType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfile {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device")))) {
		request.Device = expandRequestNfvProfileUpdateNfvProfileDeviceArray(ctx, key+".device", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDevice {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDevice{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDevice(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDevice {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_tag")))) {
		request.DeviceTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direct_internet_access_for_firewall")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direct_internet_access_for_firewall")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direct_internet_access_for_firewall")))) {
		request.DirectInternetAccessForFirewall = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services")))) {
		request.Services = expandRequestNfvProfileUpdateNfvProfileDeviceServicesArray(ctx, key+".services", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_networks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_networks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_networks")))) {
		request.CustomNetworks = expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksArray(ctx, key+".custom_networks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_for_l2")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_for_l2")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_for_l2")))) {
		request.VLANForL2 = expandRequestNfvProfileUpdateNfvProfileDeviceVLANForL2Array(ctx, key+".vlan_for_l2", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_template")))) {
		request.CustomTemplate = expandRequestNfvProfileUpdateNfvProfileDeviceCustomTemplateArray(ctx, key+".custom_template", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".current_device_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".current_device_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".current_device_tag")))) {
		request.CurrentDeviceTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceServicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServices {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServices{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceServices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceServices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServices {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_type")))) {
		request.ServiceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_type")))) {
		request.ProfileType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_name")))) {
		request.ServiceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_name")))) {
		request.ImageName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".v_nic_mapping")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".v_nic_mapping")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".v_nic_mapping")))) {
		request.VNicMapping = expandRequestNfvProfileUpdateNfvProfileDeviceServicesVNicMappingArray(ctx, key+".v_nic_mapping", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".firewall_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".firewall_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".firewall_mode")))) {
		request.FirewallMode = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceServicesVNicMappingArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceServicesVNicMapping(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceServicesVNicMapping(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_type")))) {
		request.NetworkType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_ip_address_to_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_ip_address_to_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_ip_address_to_network")))) {
		request.AssignIPAddressToNetwork = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_name")))) {
		request.NetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".services_to_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".services_to_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".services_to_connect")))) {
		request.ServicesToConnect = expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksServicesToConnectArray(ctx, key+".services_to_connect", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_type")))) {
		request.ConnectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_mode")))) {
		request.VLANMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksServicesToConnectArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksServicesToConnect(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomNetworksServicesToConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_name")))) {
		request.ServiceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceVLANForL2Array(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceVLANForL2 {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceVLANForL2{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceVLANForL2(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceVLANForL2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceVLANForL2 {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceVLANForL2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_type")))) {
		request.VLANType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_description")))) {
		request.VLANDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomTemplateArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate {
	request := []dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate{}
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
		i := expandRequestNfvProfileUpdateNfvProfileDeviceCustomTemplate(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNfvProfileUpdateNfvProfileDeviceCustomTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate {
	request := dnacentersdkgo.RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template")))) {
		request.Template = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_type")))) {
		request.TemplateType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

//TODO
