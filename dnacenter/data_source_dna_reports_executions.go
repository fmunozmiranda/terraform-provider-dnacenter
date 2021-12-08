package dnacenter

import (
	"context"
	"fmt"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsExecutions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get details of all executions for a given report

- Returns report content. Save the response to a file by converting the response data as a blob and setting the file
format available from content-disposition response header.
`,

		ReadContext: dataSourceReportsExecutionsRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. executionId of report execution
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_id": &schema.Schema{
				Description: `reportId path parameter. reportId of report
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data_category": &schema.Schema{
							Description: `data category of the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deliveries": &schema.Schema{
							Description: `Array of available delivery channels
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"execution_count": &schema.Schema{
							Description: `Total number of report executions
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"executions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Report execution pipeline end time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"errors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"execution_id": &schema.Schema{
										Description: `Report execution Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"process_status": &schema.Schema{
										Description: `Report execution status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"request_status": &schema.Schema{
										Description: `Report execution acceptance status from scheduler
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Report execution pipeline start time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"warnings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `report dataset name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_id": &schema.Schema{
							Description: `report Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_was_executed": &schema.Schema{
							Description: `true if atleast one execution has started
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"schedule": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"tags": &schema.Schema{
							Description: `array of tags for report
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"view": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `view description
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"field_groups": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"format": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `view name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_id": &schema.Schema{
										Description: `view Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_info": &schema.Schema{
										Description: `view filters info
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"view_group_id": &schema.Schema{
							Description: `viewGroupId of the viewgroup for the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_group_version": &schema.Schema{
							Description: `version of viewgroup for the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceReportsExecutionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID, okReportID := d.GetOk("report_id")
	vExecutionID, okExecutionID := d.GetOk("execution_id")

	method1 := []bool{okReportID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okReportID, okExecutionID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAllExecutionDetailsForAGivenReport")
		vvReportID := vReportID.(string)

		response1, _, err := client.Reports.GetAllExecutionDetailsForAGivenReport(vvReportID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing GetAllExecutionDetailsForAGivenReport", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")
		fmt.Println(response1)
		/*
			vvDirpath := d.Get("dirpath").(string)
			err = response1.SaveDownload(vvDirpath)
			if err != nil {
				diags = append(diags, diagError(
					"Failure when downloading file", err))
				return diags
			}
			log.Printf("[DEBUG] Downloaded file %s", vvDirpath)
		*/

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: DownloadReportContent")
		vvReportID := vReportID.(string)
		vvExecutionID := vExecutionID.(string)

		response2, _, err := client.Reports.DownloadReportContent(vvReportID, vvExecutionID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadReportContent", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response2.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}
