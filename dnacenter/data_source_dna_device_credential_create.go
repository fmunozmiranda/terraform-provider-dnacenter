package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceDeviceCredentialCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to create device credentials.
`,

		ReadContext: dataSourceDeviceCredentialCreateRead,
		Schema: map[string]*schema.Schema{
			"cli_credential": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Name or description for CLI credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_password": &schema.Schema{
							Description: `Enable password for CLI credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Description: `Password for CLI credential
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"username": &schema.Schema{
							Description: `User name for CLI credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"https_read": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name or description of http read credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Description: `Password for http read credential
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"port": &schema.Schema{
							Description: `Port for http read credential
`,
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"username": &schema.Schema{
							Description: `User name of the http read credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"https_write": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name or description of http write credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Description: `Password for http write credential
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"port": &schema.Schema{
							Description: `Port for http write credential
`,
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"username": &schema.Schema{
							Description: `User name of the http write credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"snmp_v2c_read": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description for snmp v2 read
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_community": &schema.Schema{
							Description: `Ready community for snmp v2 read credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"snmp_v2c_write": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description for snmp v2 write
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"write_community": &schema.Schema{
							Description: `Write community for snmp v2 write credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"snmp_v3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_password": &schema.Schema{
							Description: `Authentication password for snmpv3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_type": &schema.Schema{
							Description: `Authentication type for snmpv3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Name or description for SNMPV3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"privacy_password": &schema.Schema{
							Description: `Privacy password for snmpv3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"privacy_type": &schema.Schema{
							Description: `Privacy type for snmpv3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Description: `Mode for snmpv3 credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"username": &schema.Schema{
							Description: `User name for SNMPv3 credential
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

func dataSourceDeviceCredentialCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okSiteID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateDeviceCredentials")
		request1 := expandRequestDeviceCredentialCreateCreateDeviceCredentials(ctx, "", d)

		response1, _, err := client.NetworkSettings.CreateDeviceCredentials(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateDeviceCredentials", err,
				"Failure at CreateDeviceCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenNetworkSettingsCreateDeviceCredentialsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateDeviceCredentials response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: UpdateDeviceCredentials")
		request2 := expandRequestDeviceCredentialCreateUpdateDeviceCredentials(ctx, "", d)

		response2, _, err := client.NetworkSettings.UpdateDeviceCredentials(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceCredentials", err,
				"Failure at UpdateDeviceCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetDeviceCredentialDetails")
		queryParams3 := dnacentersdkgo.GetDeviceCredentialDetailsQueryParams{}

		if okSiteID {
			queryParams3.SiteID = vSiteID.(string)
		}

		response3, _, err := client.NetworkSettings.GetDeviceCredentialDetails(&queryParams3)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceCredentialDetails", err,
				"Failure at GetDeviceCredentialDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

	}
	return diags
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentials {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentials{}
	request.Settings = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettings(ctx, key, d)
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get("cli_credential"))) {
		request.CliCredential = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsCliCredentialArray(ctx, key+".cli_credential", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v2c_read"))) {
		request.SNMPV2CRead = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CReadArray(ctx, key+".snmp_v2c_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v2c_write"))) {
		request.SNMPV2CWrite = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CWriteArray(ctx, key+".snmp_v2c_write", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v3"))) {
		request.SNMPV3 = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV3Array(ctx, key+".snmp_v3", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get("https_read"))) {
		request.HTTPSRead = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSReadArray(ctx, key+".https_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get("https_write"))) {
		request.HTTPSWrite = expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSWriteArray(ctx, key+".https_write", d)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsCliCredentialArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsCliCredential(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsCliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get("enable_password"))) {
		request.EnablePassword = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get("read_community"))) {
		request.ReadCommunity = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get("write_community"))) {
		request.WriteCommunity = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV3Array(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3 {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV3(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsSNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3 {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get("privacy_type"))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get("privacy_password"))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get("auth_type"))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get("auth_password"))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_mode"))) {
		request.SNMPMode = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get("port"))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite{}
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
		i := expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceCredentialCreateCreateDeviceCredentialsSettingsHTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get("port"))) {
		request.Port = interfaceToFloat64Ptr(v)
	}
	return &request
}

func flattenNetworkSettingsCreateDeviceCredentialsItem(item *dnacentersdkgo.ResponseNetworkSettingsCreateDeviceCredentials) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials{}
	request.Settings = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettings(ctx, key, d)
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get("cli_credential"))) {
		request.CliCredential = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsCliCredential(ctx, key+".cli_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v2c_read"))) {
		request.SNMPV2CRead = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx, key+".snmp_v2c_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v2c_write"))) {
		request.SNMPV2CWrite = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx, key+".snmp_v2c_write.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_v3"))) {
		request.SNMPV3 = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV3(ctx, key+".snmp_v3.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get("https_read"))) {
		request.HTTPSRead = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsHTTPSRead(ctx, key+".https_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get("https_write"))) {
		request.HTTPSWrite = expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsHTTPSWrite(ctx, key+".https_write.0", d)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsCliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get("enable_password"))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get("read_community"))) {
		request.ReadCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get("write_community"))) {
		request.WriteCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsSNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3 {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get("auth_password"))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get("auth_type"))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get("snmp_mode"))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get("privacy_password"))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get("privacy_type"))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get("description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsHTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get("port"))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceCredentialCreateUpdateDeviceCredentialsSettingsHTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get("username"))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get("password"))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get("port"))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get("id"))) {
		request.ID = interfaceToString(v)
	}
	return &request
}
