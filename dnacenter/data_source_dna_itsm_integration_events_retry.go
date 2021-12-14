package dnacenter

import (
	"context"

	"reflect"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceItsmIntegrationEventsRetry() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on ITSM.

- Allows retry of multiple failed ITSM event instances. The retry request payload can be given as a list of strings:
["instance1","instance2","instance3",..] A minimum of one instance Id is mandatory. The list of failed event instance
Ids can be retrieved using the 'Get Failed ITSM Events' API in the 'instanceId' attribute.
`,

		ReadContext: dataSourceItsmIntegrationEventsRetryRead,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func dataSourceItsmIntegrationEventsRetryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInstanceID, okInstanceID := d.GetOk("instance_id")

	method1 := []bool{okInstanceID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetFailedItsmEvents")
		queryParams1 := dnacentersdkgo.GetFailedItsmEventsQueryParams{}

		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
		}

		response1, _, err := client.Itsm.GetFailedItsmEvents(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetFailedItsmEvents", err,
				"Failure at GetFailedItsmEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: RetryIntegrationEvents")
		request2 := expandRequestItsmIntegrationEventsRetryRetryIntegrationEvents(ctx, "", d)

		response2, _, err := client.Itsm.RetryIntegrationEvents(request2)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetryIntegrationEvents", err,
				"Failure at RetryIntegrationEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenItsmRetryIntegrationEventsItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetryIntegrationEvents response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestItsmIntegrationEventsRetryRetryIntegrationEvents(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItsmRetryIntegrationEvents {
	request := dnacentersdkgo.RequestItsmRetryIntegrationEvents{}
	if v, ok := d.GetOkExists(fixKeyAccess(key)); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key)))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key)))) {
		request = interfaceToSliceString(v)
	}
	return &request
}

func flattenItsmRetryIntegrationEventsItem(item *dnacentersdkgo.ResponseItsmRetryIntegrationEvents) []map[string]interface{} {
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
