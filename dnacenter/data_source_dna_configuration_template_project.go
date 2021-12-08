package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplateProject() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- List the projects

- Get the details of the given project by its id.
`,

		ReadContext: dataSourceConfigurationTemplateProjectRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name query parameter. Name of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId path parameter. projectId(UUID) of project to get project details
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceConfigurationTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vProjectID, okProjectID := d.GetOk("project_id")

	method1 := []bool{okName, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsAListOfProjects")
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, _, err := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsAListOfProjects", err,
				"Failure at GetsAListOfProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenConfigurationTemplatesGetsAListOfProjectsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAListOfProjects response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetsTheDetailsOfAGivenProject")
		vvProjectID := vProjectID.(string)

		response2, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheDetailsOfAGivenProject response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())

	}
	return diags
}

func flattenConfigurationTemplatesGetsAListOfProjectsItems(items *dnacentersdkgo.ResponseConfigurationTemplatesGetsAListOfProjects) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(item.Tags)
		respItem["create_time"] = item.CreateTime
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["templates"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(item.Templates)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(items *[]dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(item *dnacentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates) interface{} {
	if item == nil {
		return nil
	}
	respItem := item

	return respItem

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(item.Tags)
	respItem["create_time"] = item.CreateTime
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["templates"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(item.Templates)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(items *[]dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(item *dnacentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates) interface{} {
	if item == nil {
		return nil
	}
	respItem := item

	return respItem

}
