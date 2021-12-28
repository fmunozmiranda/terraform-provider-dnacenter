package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReports() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Reports.

- Create/Schedule a report configuration. Use "Get view details for a given view group & view" API to get the metadata
required to configure a report.

- Delete a scheduled report configuration. Deletes the report executions also.
`,

		CreateContext: resourceReportsCreate,
		ReadContext:   resourceReportsRead,
		UpdateContext: resourceReportsUpdate,
		DeleteContext: resourceReportsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'array of tags for report\n'}, 'deliveries': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of available delivery channels\n', 'Elem': {'Schema': {}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'report name\n'}, 'schedule': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'view': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'fieldGroups': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'fieldGroupDisplayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Field group label/displayname for user\n'}, 'fieldGroupName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Field group name\n'}, 'fields': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'field label/displayname\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'field name\n'}}}}}}}, 'filters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'filter label/displayname\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'filter name\n'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'filter type\n'}, 'value': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.\n', 'Elem': {'Schema': {}}}}}}, 'format': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'formatType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'format type of report\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'format name of report\n'}}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'view name\n'}, 'viewId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'view Id\n'}}}}, 'viewGroupId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'viewGroupId of the viewgroup for the report\n'}, 'viewGroupVersion': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'version of viewgroup for the report\n'}, 'reportId': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'reportId path parameter. reportId of report\n'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'array of tags for report\n'}, 'dataCategory': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'data category of the report\n'}, 'deliveries': {'Computed': 'true', 'Type': 'schema.TypeList', 'Description': 'Array of available delivery channels\n', 'Elem': {'Schema': {}}}, 'executionCount': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Total number of report executions\n'}, 'executions': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'endTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Report execution pipeline end time\n'}, 'errors': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'executionId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Report execution Id.\n'}, 'processStatus': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Report execution status\n'}, 'requestStatus': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Report execution acceptance status from scheduler\n'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Report execution pipeline start time\n'}, 'warnings': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'report name\n'}, 'reportId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'report Id\n'}, 'reportWasExecuted': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'true if atleast one execution has started\n'}, 'schedule': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'view': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'fieldGroups': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'fieldGroupDisplayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Field group label/displayname for user\n'}, 'fieldGroupName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Field group name\n'}, 'fields': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'field label/displayname\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'field name\n'}}}}}}}, 'filters': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'filter label/displayname\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'filter name\n'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'filter type\n'}, 'value': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'value of filter. data type is based on the filter type.\n', 'Elem': {'Schema': {}}}}}}, 'format': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'formatType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'format type of report\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'format name of report\n'}, 'default': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'true, if the format type is considered default\n'}}}}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'view name\n'}, 'viewId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'view Id\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'view description\n'}, 'viewInfo': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'view filters info\n'}}}}, 'viewGroupId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'viewGroupId of the viewgroup for the report\n'}, 'viewGroupVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'version of viewgroup for the report\n'}}}}}, 'metadata': {'item': {'operation_id': ['CreateOrScheduleAReport', 'GetAScheduledReport'], 'new_flat_structure': [{'RequestReportsCreateOrScheduleAReport': {'type': 'obj', 'data': [{'name': 'tags', 'description': 'array of tags for report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'deliveries', 'description': 'Array of available delivery channels\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestReportsCreateOrScheduleAReportDeliveries'}, {'name': 'name', 'description': 'report name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'schedule', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestReportsCreateOrScheduleAReportSchedule'}, {'name': 'view', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestReportsCreateOrScheduleAReportView'}, {'name': 'viewGroupId', 'description': 'viewGroupId of the viewgroup for the report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'viewGroupVersion', 'description': 'version of viewgroup for the report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportDeliveries': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportSchedule': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportView': {'type': 'obj', 'data': [{'name': 'fieldGroups', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestReportsCreateOrScheduleAReportViewFieldGroups'}, {'name': 'filters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestReportsCreateOrScheduleAReportViewFilters'}, {'name': 'format', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestReportsCreateOrScheduleAReportViewFormat'}, {'name': 'name', 'description': 'view name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'viewId', 'description': 'view Id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportViewFieldGroups': {'type': 'obj', 'data': [{'name': 'fieldGroupDisplayName', 'description': 'Field group label/displayname for user\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fieldGroupName', 'description': 'Field group name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fields', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestReportsCreateOrScheduleAReportViewFieldGroupsFields'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportViewFieldGroupsFields': {'type': 'obj', 'data': [{'name': 'displayName', 'description': 'field label/displayname\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'field name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportViewFilters': {'type': 'obj', 'data': [{'name': 'displayName', 'description': 'filter label/displayname\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'filter name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'filter type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestReportsCreateOrScheduleAReportViewFiltersValue'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportViewFiltersValue': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestReportsCreateOrScheduleAReportViewFormat': {'type': 'obj', 'data': [{'name': 'formatType', 'description': 'format type of report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'format name of report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'ResponseReportsGetAScheduledReport': {'type': 'obj', 'data': [{'name': 'tags', 'description': 'array of tags for report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'dataCategory', 'description': 'data category of the report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deliveries', 'description': 'Array of available delivery channels\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseReportsGetAScheduledReportDeliveries'}, {'name': 'executionCount', 'description': 'Total number of report executions\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'executions', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseReportsGetAScheduledReportExecutions'}, {'name': 'name', 'description': 'report name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reportId', 'description': 'report Id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reportWasExecuted', 'description': 'true if atleast one execution has started\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'schedule', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseReportsGetAScheduledReportSchedule'}, {'name': 'view', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseReportsGetAScheduledReportView'}, {'name': 'viewGroupId', 'description': 'viewGroupId of the viewgroup for the report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'viewGroupVersion', 'description': 'version of viewgroup for the report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportDeliveries': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportExecutions': {'type': 'obj', 'data': [{'name': 'endTime', 'description': 'Report execution pipeline end time\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'errors', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'executionId', 'description': 'Report execution Id.\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'processStatus', 'description': 'Report execution status\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'requestStatus', 'description': 'Report execution acceptance status from scheduler\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': 'Report execution pipeline start time\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'warnings', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportSchedule': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportView': {'type': 'obj', 'data': [{'name': 'fieldGroups', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseReportsGetAScheduledReportViewFieldGroups'}, {'name': 'filters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseReportsGetAScheduledReportViewFilters'}, {'name': 'format', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseReportsGetAScheduledReportViewFormat'}, {'name': 'name', 'description': 'view name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'viewId', 'description': 'view Id\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'view description\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'viewInfo', 'description': 'view filters info\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportViewFieldGroups': {'type': 'obj', 'data': [{'name': 'fieldGroupDisplayName', 'description': 'Field group label/displayname for user\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fieldGroupName', 'description': 'Field group name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fields', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseReportsGetAScheduledReportViewFieldGroupsFields'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportViewFieldGroupsFields': {'type': 'obj', 'data': [{'name': 'displayName', 'description': 'field label/displayname\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'field name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportViewFilters': {'type': 'obj', 'data': [{'name': 'displayName', 'description': 'filter label/displayname\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'filter name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'filter type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'value of filter. data type is based on the filter type.\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseReportsGetAScheduledReportViewFiltersValue'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportViewFiltersValue': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseReportsGetAScheduledReportViewFormat': {'type': 'obj', 'data': [{'name': 'formatType', 'description': 'format type of report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'format name of report\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'default', 'description': 'true, if the format type is considered default\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestReportsCreateOrScheduleAReport', 'ResponseReportsGetAScheduledReport'], 'access_list': [['tags', 'deliveries', 'name', 'schedule', 'view', 'viewGroupId', 'viewGroupVersion'], []]}}}
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
							Description: `report name
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type.
`,
													Type:     schema.TypeList,
													Computed: true,
												},
											},
										},
									},
									"format": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default": &schema.Schema{
													Description: `true, if the format type is considered default
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deliveries": &schema.Schema{
							Description: `Array of available delivery channels
`,
							Type:     schema.TypeList,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `report name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"report_id": &schema.Schema{
							Description: `reportId path parameter. reportId of report
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
						},
						"tags": &schema.Schema{
							Description: `array of tags for report
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"view": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"field_groups": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.
`,
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
												},
											},
										},
									},
									"format": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `view name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"view_id": &schema.Schema{
										Description: `view Id
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"view_group_id": &schema.Schema{
							Description: `viewGroupId of the viewgroup for the report
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"view_group_version": &schema.Schema{
							Description: `version of viewgroup for the report
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

func resourceReportsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestReportsCreateOrScheduleAReport(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vReportID, okReportID := resourceItem["report_id"]
	vvReportID := interfaceToString(vReportID)
	if okReportID && vvReportID != "" {
		getResponse2, _, err := client.Reports.GetAScheduledReport(vvReportID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["report_id"] = vvReportID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.Reports.GetListOfScheduledReports(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsReportsGetListOfScheduledReports(m, response2, nil)
			item2, err := searchReportsGetListOfScheduledReports(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["report_id"] = vvReportID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.Reports.CreateOrScheduleAReport(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateOrScheduleAReport", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateOrScheduleAReport", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["report_id"] = vvReportID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceReportsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vViewGroupID, okViewGroupID := resourceMap["view_group_id"]
	vViewID, okViewID := resourceMap["view_id"]
	vReportID, okReportID := resourceMap["report_id"]

	method1 := []bool{okViewGroupID, okViewID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okReportID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetListOfScheduledReports")
		queryParams1 := dnacentersdkgo.GetListOfScheduledReportsQueryParams{}

		if okViewGroupID {
			queryParams1.ViewGroupID = vViewGroupID.(string)
		}
		if okViewID {
			queryParams1.ViewID = vViewID.(string)
		}

		response1, restyResp1, err := client.Reports.GetListOfScheduledReports(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetListOfScheduledReports", err,
				"Failure at GetListOfScheduledReports, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetAScheduledReport")
		vvReportID := vReportID.(string)

		response2, restyResp2, err := client.Reports.GetAScheduledReport(vvReportID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAScheduledReport", err,
				"Failure at GetAScheduledReport, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourceReportsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceReportsRead(ctx, d, m)
}

func resourceReportsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestReportsCreateOrScheduleAReport(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReport {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReport{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deliveries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deliveries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deliveries")))) {
		request.Deliveries = expandRequestReportsCreateOrScheduleAReportDeliveriesArray(ctx, key+".deliveries", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schedule")))) {
		request.Schedule = expandRequestReportsCreateOrScheduleAReportSchedule(ctx, key+".schedule.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view")))) {
		request.View = expandRequestReportsCreateOrScheduleAReportView(ctx, key+".view.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_group_id")))) {
		request.ViewGroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_group_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_group_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_group_version")))) {
		request.ViewGroupVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportDeliveriesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries{}
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
		i := expandRequestReportsCreateOrScheduleAReportDeliveries(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportDeliveries(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportDeliveries
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportSchedule {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportSchedule
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportView(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportView {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportView{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_groups")))) {
		request.FieldGroups = expandRequestReportsCreateOrScheduleAReportViewFieldGroupsArray(ctx, key+".field_groups", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestReportsCreateOrScheduleAReportViewFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".format")))) {
		request.Format = expandRequestReportsCreateOrScheduleAReportViewFormat(ctx, key+".format.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_id")))) {
		request.ViewID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFieldGroups(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_group_display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_group_display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_group_display_name")))) {
		request.FieldGroupDisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_group_name")))) {
		request.FieldGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fields")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fields")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fields")))) {
		request.Fields = expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFieldsArray(ctx, key+".fields", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFieldsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFields(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFieldGroupsFields(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFieldGroupsFields{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters {
	request := []dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters{}
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
		i := expandRequestReportsCreateOrScheduleAReportViewFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestReportsCreateOrScheduleAReportViewFiltersValue(ctx, key+".value.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFiltersValue {
	var request dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFiltersValue
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestReportsCreateOrScheduleAReportViewFormat(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFormat {
	request := dnacentersdkgo.RequestReportsCreateOrScheduleAReportViewFormat{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".format_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".format_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".format_type")))) {
		request.FormatType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
