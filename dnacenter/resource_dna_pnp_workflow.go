package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePnpWorkflow() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Onboarding (PnP).

- Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database

- Deletes a workflow specified by id

- Updates an existing workflow
`,

		CreateContext: resourcePnpWorkflowCreate,
		ReadContext:   resourcePnpWorkflowRead,
		UpdateContext: resourcePnpWorkflowUpdate,
		DeleteContext: resourcePnpWorkflowDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'_id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'addToInventory': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'addedOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'configId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'currTaskIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'execTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'imageId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'lastupdateOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'tasks': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'currWorkItemIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'taskSeqNo': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workItemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'command': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'outputStr': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'tenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'useState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'_id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'lastupdateOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Lastupdate On'}, 'imageId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image Id'}, 'currTaskIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Task Idx'}, 'addedOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Added On'}, 'tasks': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'currWorkItemIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Work Item Idx'}, 'taskSeqNo': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Task Seq No'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'workItemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'command': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Command'}, 'outputStr': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Output Str'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}}}}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'addToInventory': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Add To Inventory'}, 'instanceType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instance Type'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'execTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Exec Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'useState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Use State'}, 'configId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config Id'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'version': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Version'}, 'tenantId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Tenant Id'}}}}}, 'metadata': {'item': {'operation_id': [['AddAWorkflow', 'UpdateWorkflow'], 'GetWorkflowById'], 'new_flat_structure': [[{'RequestDeviceOnboardingPnpAddAWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddAWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddAWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestDeviceOnboardingPnpUpdateWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], {'ResponseDeviceOnboardingPnpGetWorkflowById': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetWorkflowByIdTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetWorkflowByIdTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetWorkflowByIdTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetWorkflowByIdTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetWorkflowByID': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetWorkflowByIdTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetWorkflowByIDTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetWorkflowByIdTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': [['RequestDeviceOnboardingPnpAddAWorkflow', 'RequestDeviceOnboardingPnpUpdateWorkflow'], 'ResponseDeviceOnboardingPnpGetWorkflowById'], 'access_list': [[[], []], []]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"add_to_inventory": &schema.Schema{
							Description: `Add To Inventory`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"added_on": &schema.Schema{
							Description: `Added On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"config_id": &schema.Schema{
							Description: `Config Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"curr_task_idx": &schema.Schema{
							Description: `Curr Task Idx`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"exec_time": &schema.Schema{
							Description: `Exec Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"image_id": &schema.Schema{
							Description: `Image Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"lastupdate_on": &schema.Schema{
							Description: `Lastupdate On`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"state": &schema.Schema{
							Description: `State`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"curr_work_item_idx": &schema.Schema{
										Description: `Curr Work Item Idx`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"end_time": &schema.Schema{
										Description: `End Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"state": &schema.Schema{
										Description: `State`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"task_seq_no": &schema.Schema{
										Description: `Task Seq No`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"time_taken": &schema.Schema{
										Description: `Time Taken`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Description: `Command`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"output_str": &schema.Schema{
													Description: `Output Str`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"time_taken": &schema.Schema{
													Description: `Time Taken`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"use_state": &schema.Schema{
							Description: `Use State`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"add_to_inventory": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"added_on": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"curr_task_idx": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"exec_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"image_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lastupdate_on": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"curr_work_item_idx": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"task_seq_no": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"time_taken": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"end_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"output_str": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_time": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"state": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"time_taken": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPnpWorkflowAddAWorkflow(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.DeviceOnboardingPnp.GetWorkflows(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceOnboardingPnpGetWorkflows(m, response2, nil)
			item2, err := searchDeviceOnboardingPnpGetWorkflows(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddAWorkflow(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddAWorkflow", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddAWorkflow", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourcePnpWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSort, okSort := resourceMap["sort"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vType, okType := resourceMap["type"]
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okType, okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetWorkflows")
		queryParams1 := dnacentersdkgo.GetWorkflowsQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okSort {
			queryParams1.Sort = interfaceToSliceString(vSort)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okType {
			queryParams1.Type = interfaceToSliceString(vType)
		}
		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetWorkflows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflows", err,
				"Failure at GetWorkflows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetWorkflowByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetWorkflowByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetWorkflowByID", err,
				"Failure at GetWorkflowByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourcePnpWorkflowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSort, okSort := resourceMap["sort"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vType, okType := resourceMap["type"]
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okType, okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestPnpWorkflowUpdateWorkflow(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceOnboardingPnp.UpdateWorkflow(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateWorkflow", err, restyResp1.String(),
					"Failure at UpdateWorkflow, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateWorkflow", err,
				"Failure at UpdateWorkflow, unexpected response", ""))
			return diags
		}
	}

	return resourcePnpWorkflowRead(ctx, d, m)
}

func resourcePnpWorkflowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestPnpWorkflowAddAWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpWorkflowAddAWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks{}
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
		i := expandRequestPnpWorkflowAddAWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpWorkflowAddAWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpWorkflowAddAWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowAddAWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".add_to_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".add_to_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".add_to_inventory")))) {
		request.AddToInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_task_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_task_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_task_idx")))) {
		request.CurrTaskIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exec_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exec_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exec_time")))) {
		request.ExecTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_type")))) {
		request.InstanceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lastupdate_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lastupdate_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lastupdate_on")))) {
		request.LastupdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tasks")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tasks")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tasks")))) {
		request.Tasks = expandRequestPnpWorkflowUpdateWorkflowTasksArray(ctx, key+".tasks", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_state")))) {
		request.UseState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks{}
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
		i := expandRequestPnpWorkflowUpdateWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".curr_work_item_idx")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".curr_work_item_idx")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".curr_work_item_idx")))) {
		request.CurrWorkItemIDx = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_seq_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_seq_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_seq_no")))) {
		request.TaskSeqNo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpWorkflowUpdateWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".output_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".output_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".output_str")))) {
		request.OutputStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
