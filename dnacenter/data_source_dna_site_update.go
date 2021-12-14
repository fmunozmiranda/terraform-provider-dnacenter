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
func dataSourceSiteUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sites.

- Update site area/building/floor with specified hierarchy and new values
`,

		ReadContext: dataSourceSiteUpdateRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site id to which site details to be updated.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"parent_name": &schema.Schema{
							Description: `Parent Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"building": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"address": &schema.Schema{
							Description: `Address`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"latitude": &schema.Schema{
							Description: `Latitude`,
							Type:        schema.TypeFloat,
							Optional:    true,
						},
						"longitude": &schema.Schema{
							Description: `Longitude`,
							Type:        schema.TypeFloat,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"parent_name": &schema.Schema{
							Description: `Parent Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"floor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"height": &schema.Schema{
							Description: `Height`,
							Type:        schema.TypeFloat,
							Optional:    true,
						},
						"length": &schema.Schema{
							Description: `Length`,
							Type:        schema.TypeFloat,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rf_model": &schema.Schema{
							Description: `Rf Model. Allowed values are 'Cubes And Walled Offices', 'Drywall Office Only', 'Indoor High Ceiling', 'Outdoor Open Space'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"width": &schema.Schema{
							Description: `Width`,
							Type:        schema.TypeFloat,
							Optional:    true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Description: `Data`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_error": &schema.Schema{
							Description: `Is Error`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"operation_id_list": &schema.Schema{
							Description: `Operation Id List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"progress": &schema.Schema{
							Description: `Progress`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"root_id": &schema.Schema{
							Description: `Root Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"service_type": &schema.Schema{
							Description: `Service Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vRunsync, okRunsync := d.GetOk("runsync")
	vTimeout, okTimeout := d.GetOk("timeout")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	method1 := []bool{okSiteID, okRunsync, okTimeout, okPersistbapioutput}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okSiteID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateSite")
		vvSiteID := vSiteID.(string)
		headerParams1 := dnacentersdkgo.UpdateSiteHeaderParams{}
		request1 := expandRequestSiteUpdateUpdateSite(ctx, "", d)
		if okRunsync {
			headerParams1.Runsync = vRunsync.(string)
		}
		if okTimeout {
			headerParams1.Timeout = vTimeout.(string)
		}
		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, _, err := client.Sites.UpdateSite(vvSiteID, request1, &headerParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSite", err,
				"Failure at UpdateSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSitesUpdateSiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateSite response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: DeleteSite")
		vvSiteID := vSiteID.(string)

		response2, _, err := client.Sites.DeleteSite(vvSiteID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteSite", err,
				"Failure at DeleteSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

	}
	return diags
}

func expandRequestSiteUpdateUpdateSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSite {
	request := dnacentersdkgo.RequestSitesUpdateSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site")))) {
		request.Site = expandRequestSiteUpdateUpdateSiteSite(ctx, key+".site.0", d)
	}
	return &request
}

func expandRequestSiteUpdateUpdateSiteSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSite {
	request := dnacentersdkgo.RequestSitesUpdateSiteSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".area")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".area")))) && (ok || !reflect.DeepEqual(v, d.Get("area"))) {
		request.Area = expandRequestSiteUpdateUpdateSiteSiteArea(ctx, key+".area.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".building")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".building")))) && (ok || !reflect.DeepEqual(v, d.Get("building"))) {
		request.Building = expandRequestSiteUpdateUpdateSiteSiteBuilding(ctx, key+".building.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor")))) && (ok || !reflect.DeepEqual(v, d.Get("floor"))) {
		request.Floor = expandRequestSiteUpdateUpdateSiteSiteFloor(ctx, key+".floor.0", d)
	}
	return &request
}

func expandRequestSiteUpdateUpdateSiteSiteArea(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteArea {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteArea{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get("parent_name"))) {
		request.ParentName = interfaceToString(v)
	}
	return &request
}

func expandRequestSiteUpdateUpdateSiteSiteBuilding(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteBuilding {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteBuilding{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get("address"))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_name")))) && (ok || !reflect.DeepEqual(v, d.Get("parent_name"))) {
		request.ParentName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get("latitude"))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get("longitude"))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	return &request
}

func expandRequestSiteUpdateUpdateSiteSiteFloor(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesUpdateSiteSiteFloor {
	request := dnacentersdkgo.RequestSitesUpdateSiteSiteFloor{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get("name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get("rf_model"))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get("width"))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get("length"))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get("height"))) {
		request.Height = interfaceToFloat64Ptr(v)
	}
	return &request
}

func flattenSitesUpdateSiteItem(item *dnacentersdkgo.ResponseSitesUpdateSiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["version"] = item.Version
	respItem["start_time"] = item.StartTime
	respItem["progress"] = item.Progress
	respItem["data"] = item.Data
	respItem["service_type"] = item.ServiceType
	respItem["operation_id_list"] = item.OperationIDList
	respItem["is_error"] = item.IsError
	respItem["root_id"] = item.RootID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
