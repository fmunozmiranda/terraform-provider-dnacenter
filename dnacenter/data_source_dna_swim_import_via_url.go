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
func dataSourceSwimImportViaURL() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image
files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
`,

		ReadContext: dataSourceSwimImportViaURLRead,
		Schema: map[string]*schema.Schema{
			"schedule_at": &schema.Schema{
				Description: `scheduleAt query parameter. Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional) 
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_desc": &schema.Schema{
				Description: `scheduleDesc query parameter. Custom Description (Optional)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_origin": &schema.Schema{
				Description: `scheduleOrigin query parameter. Originator of this call (Optional)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_family": &schema.Schema{
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
			"source_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"third_party": &schema.Schema{
				// Type:     schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceSwimImportViaURLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vScheduleAt, okScheduleAt := d.GetOk("schedule_at")
	vScheduleDesc, okScheduleDesc := d.GetOk("schedule_desc")
	vScheduleOrigin, okScheduleOrigin := d.GetOk("schedule_origin")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportSoftwareImageViaURL")
		queryParams1 := dnacentersdkgo.ImportSoftwareImageViaURLQueryParams{}
		request1 := expandRequestSwimImportViaURLImportSoftwareImageViaURL(ctx, "", d)
		if okScheduleAt {
			queryParams1.ScheduleAt = vScheduleAt.(string)
		}
		if okScheduleDesc {
			queryParams1.ScheduleDesc = vScheduleDesc.(string)
		}
		if okScheduleOrigin {
			queryParams1.ScheduleOrigin = vScheduleOrigin.(string)
		}

		response1, _, err := client.SoftwareImageManagementSwim.ImportSoftwareImageViaURL(request1, &queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportSoftwareImageViaURL", err,
				"Failure at ImportSoftwareImageViaURL, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSoftwareImageManagementSwimImportSoftwareImageViaURLItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportSoftwareImageViaURL response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSwimImportViaURLImportSoftwareImageViaURL(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURL{}
	if v := expandRequestSwimImportViaURLImportSoftwareImageViaURLArray(ctx, key+".", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSwimImportViaURLImportSoftwareImageViaURLArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := []dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL{}
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
		i := expandRequestItemSwimImportViaURLImportSoftwareImageViaURL(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestItemSwimImportViaURLImportSoftwareImageViaURL(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get("application_type"))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_family")))) && (ok || !reflect.DeepEqual(v, d.Get("image_family"))) {
		request.ImageFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_url")))) && (ok || !reflect.DeepEqual(v, d.Get("source_url"))) {
		request.SourceURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get("third_party"))) {
		request.ThirdParty = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vendor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vendor")))) && (ok || !reflect.DeepEqual(v, d.Get("vendor"))) {
		request.Vendor = interfaceToString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimImportSoftwareImageViaURLItem(item *dnacentersdkgo.ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLResponse) []map[string]interface{} {
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
