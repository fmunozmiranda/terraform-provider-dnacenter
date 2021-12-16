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
func dataSourceConfigurationTemplateDeployV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- V2 API to deploy a template.
`,

		ReadContext: dataSourceConfigurationTemplateDeployV2Read,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
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
		},
	}
}

func dataSourceConfigurationTemplateDeployV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeployTemplateV2")
		request1 := expandRequestConfigurationTemplateDeployV2DeployTemplateV2(ctx, "", d)

		response1, restyResp1, err := client.ConfigurationTemplates.DeployTemplateV2(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployTemplateV2", err,
				"Failure at DeployTemplateV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesDeployTemplateV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployTemplateV2 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2 {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_push_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_push_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_push_template")))) {
		request.ForcePushTemplate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_composite")))) {
		request.IsComposite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".main_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".main_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".main_template_id")))) {
		request.MainTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_template_deployment_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_template_deployment_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_template_deployment_info")))) {
		request.MemberTemplateDeploymentInfo = expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx, key+".member_template_deployment_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_info")))) {
		request.TargetInfo = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoArray(ctx, key+".target_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoParams(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParamsArray(ctx, key+".resource_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".versioned_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".versioned_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".versioned_template_id")))) {
		request.VersionedTemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams{}
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
		i := expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestConfigurationTemplateDeployV2DeployTemplateV2TargetInfoResourceParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenConfigurationTemplatesDeployTemplateV2Item(item *dnacentersdkgo.ResponseConfigurationTemplatesDeployTemplateV2Response) []map[string]interface{} {
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
