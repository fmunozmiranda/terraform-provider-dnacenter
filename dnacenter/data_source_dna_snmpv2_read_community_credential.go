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

// dataSourceAction
func dataSourceSNMPv2ReadCommunityCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create and update operations on Discovery.

- Updates global SNMP read community

- Adds global SNMP read community
`,

		ReadContext: dataSourceSNMPv2ReadCommunityCredentialRead,
		Schema: map[string]*schema.Schema{
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"payload": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Description: `Comments to identify the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"credential_type": &schema.Schema{
							Description: `Credential type to identify the application that uses the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Name/Description of the credential
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_community": &schema.Schema{
							Description: `SNMP read community. NO!$DATA!$ for no value change
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

func dataSourceSNMPv2ReadCommunityCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateSNMPReadCommunity")
		request1 := expandRequestSNMPv2ReadCommunityCredentialUpdateSNMPReadCommunity(ctx, "", d)

		response1, restyResp1, err := client.Discovery.UpdateSNMPReadCommunity(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSNMPReadCommunity", err,
				"Failure at UpdateSNMPReadCommunity, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: CreateSNMPReadCommunity")
		request2 := expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunity(ctx, "", d)

		response2, restyResp2, err := client.Discovery.CreateSNMPReadCommunity(request2)

		if request2 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request2))
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSNMPReadCommunity", err,
				"Failure at CreateSNMPReadCommunity, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenDiscoveryCreateSNMPReadCommunityItemName(response2.Response)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSNMPReadCommunity response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags
		vItemID2 := flattenDiscoveryCreateSNMPReadCommunityItemID(response2.Response)
		if err := d.Set("item_id", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSNMPReadCommunity response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSNMPv2ReadCommunityCredentialUpdateSNMPReadCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateSNMPReadCommunity {
	request := dnacentersdkgo.RequestDiscoveryUpdateSNMPReadCommunity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	return &request
}

func expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateSNMPReadCommunity {
	request := dnacentersdkgo.RequestDiscoveryCreateSNMPReadCommunity{}
	if v := expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunityItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunityItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity {
	request := []dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity{}
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
		i := expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunityItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSNMPv2ReadCommunityCredentialCreateSNMPReadCommunityItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity {
	request := dnacentersdkgo.RequestItemDiscoveryCreateSNMPReadCommunity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	return &request
}

func flattenDiscoveryCreateSNMPReadCommunityItemName(item *dnacentersdkgo.ResponseDiscoveryCreateSNMPReadCommunityResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDiscoveryCreateSNMPReadCommunityItemID(item *dnacentersdkgo.ResponseDiscoveryCreateSNMPReadCommunityResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
