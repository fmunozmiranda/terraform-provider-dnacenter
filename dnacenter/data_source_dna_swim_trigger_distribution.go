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
func dataSourceSwimTriggerDistribution() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it
can be distributed
`,

		ReadContext: dataSourceSwimTriggerDistributionRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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

func dataSourceSwimTriggerDistributionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: TriggerSoftwareImageDistribution")
		request1 := expandRequestSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx, "", d)

		response1, _, err := client.SoftwareImageManagementSwim.TriggerSoftwareImageDistribution(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing TriggerSoftwareImageDistribution", err,
				"Failure at TriggerSoftwareImageDistribution, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting TriggerSoftwareImageDistribution response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
	if v := expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSwimTriggerDistributionTriggerSoftwareImageDistributionArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := []dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
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
		i := expandRequestItemSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemSwimTriggerDistributionTriggerSoftwareImageDistribution(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution {
	request := dnacentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get("device_uuid"))) {
		request.DeviceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get("image_uuid"))) {
		request.ImageUUID = interfaceToString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse) []map[string]interface{} {
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
