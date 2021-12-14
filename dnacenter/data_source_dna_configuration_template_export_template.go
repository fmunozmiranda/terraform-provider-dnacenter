package dnacenter

import (
	"context"

	"fmt"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceConfigurationTemplateExportTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- Exports the templates for given templateIds.
`,

		ReadContext: dataSourceConfigurationTemplateExportTemplateRead,
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

func dataSourceConfigurationTemplateExportTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportsTheTemplatesForAGivenCriteria")
		request1 := expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx, "", d)

		response1, _, err := client.ConfigurationTemplates.ExportsTheTemplatesForAGivenCriteria(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ExportsTheTemplatesForAGivenCriteria", err,
				"Failure at ExportsTheTemplatesForAGivenCriteria, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ExportsTheTemplatesForAGivenCriteria response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	request := dnacentersdkgo.RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria{}
	if v := expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria{}
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
		i := expandRequestItemConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria
	request = d.Get(key)
	return &request
}

func flattenConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaItem(item *dnacentersdkgo.ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaResponse) []map[string]interface{} {
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
