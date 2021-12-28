package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePnpDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Onboarding (PnP).

- Adds a device to the PnP database.

- Updates device details specified by device id in PnP database

- Deletes specified device from PnP database
`,

		CreateContext: resourcePnpDeviceCreate,
		ReadContext:   resourcePnpDeviceRead,
		UpdateContext: resourcePnpDeviceUpdate,
		DeleteContext: resourcePnpDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'_id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'deviceInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aaaCredentials': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'password': {'Optional': 'true', 'Type': 'schema.TypeString', 'Sensitive': 'true'}, 'username': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'addedOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'addnMacAddrs': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'agentType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'authStatus': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'authenticatedSudiSerialNo': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'capabilitiesSupported': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'cmState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'deviceSudiSerialNos': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'deviceType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'featuresSupported': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'fileSystemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'freespace': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'readable': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'size': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'writeable': {'Optional': 'true', 'Type': 'schema.TypeBool'}}}}, 'firstContact': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'hostname': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'httpHeaders': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'value': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'imageFile': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'imageVersion': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'ipInterfaces': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ipv4Address': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'ipv6AddressList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {}}}, 'macAddress': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'status': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'lastContact': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'lastSyncTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'lastUpdateOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'location': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'address': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'altitude': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'latitude': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'longitude': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'siteId': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'macAddress': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'mode': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'neighborLinks': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'localInterfaceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'localMacAddress': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'localShortInterfaceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remoteDeviceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remoteInterfaceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remoteMacAddress': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remotePlatform': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remoteShortInterfaceName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'remoteVersion': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'onbState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'pid': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'pnpProfileList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'createdBy': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'discoveryCreated': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'primaryEndpoint': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'certificate': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'fqdn': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'ipv4Address': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'ipv6Address': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'port': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'protocol': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'profileName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'secondaryEndpoint': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'certificate': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'fqdn': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'ipv4Address': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'ipv6Address': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'port': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'protocol': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'populateInventory': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'preWorkflowCliOuputs': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'cli': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'cliOutput': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'projectId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'projectName': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'reloadRequested': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'serialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'smartAccountId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'source': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'stack': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'stackInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'isFullRing': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'stackMemberList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'hardwareVersion': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'licenseLevel': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'licenseType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'macAddress': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'pid': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'priority': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'role': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'serialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'softwareVersion': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'stackNumber': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'sudiSerialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'stackRingProtocol': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'supportsStackWorkflows': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'totalMemberCount': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'validLicenseLevels': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'sudiRequired': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'tags': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'userSudiSerialNos': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'virtualAccountId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workflowId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workflowName': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'runSummaryList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'details': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'errorFlag': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'historyTaskInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'addnDetails': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'value': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workItemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'command': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'outputStr': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'timestamp': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}, 'systemResetWorkflow': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'addToInventory': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'addedOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'configId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'currTaskIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'execTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'imageId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'lastupdateOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'tasks': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'currWorkItemIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'taskSeqNo': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workItemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'command': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'outputStr': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'tenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'useState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}, 'systemWorkflow': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'addToInventory': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'addedOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'configId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'currTaskIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'execTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'imageId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'lastupdateOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'tasks': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'currWorkItemIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'taskSeqNo': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workItemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'command': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'outputStr': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'tenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'useState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}, 'tenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'workflow': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'addToInventory': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'addedOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'configId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'currTaskIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'execTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'imageId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'lastupdateOn': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'tasks': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'currWorkItemIdx': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'taskSeqNo': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'workItemList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'command': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'endTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'outputStr': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'startTime': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'state': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'timeTaken': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'tenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'type': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'useState': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'version': {'Optional': 'true', 'Type': 'schema.TypeInt'}}}}, 'workflowParameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'configList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'configId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'configParameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'value': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'licenseLevel': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'licenseType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'topOfStackSerialNumber': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'_id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}, 'deviceInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'source': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Source'}, 'serialNumber': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Serial Number'}, 'stack': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Stack'}, 'mode': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Mode'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'location': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'siteId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Id'}, 'address': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Address'}, 'latitude': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Latitude'}, 'longitude': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Longitude'}, 'altitude': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Altitude'}}}}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'onbState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Onb State'}, 'authenticatedMicNumber': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Authenticated Mic Number'}, 'authenticatedSudiSerialNo': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Authenticated Sudi Serial No'}, 'capabilitiesSupported': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Capabilities Supported'}, 'featuresSupported': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Features Supported'}, 'cmState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Cm State'}, 'firstContact': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'First Contact'}, 'lastContact': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Last Contact'}, 'macAddress': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Mac Address'}, 'pid': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Pid'}, 'deviceSudiSerialNos': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Device Sudi Serial Nos'}, 'lastUpdateOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Last Update On'}, 'workflowId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Workflow Id'}, 'workflowName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Workflow Name'}, 'projectId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Project Id'}, 'projectName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Project Name'}, 'deviceType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device Type'}, 'agentType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Agent Type'}, 'imageVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image Version'}, 'fileSystemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'writeable': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Writeable'}, 'freespace': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Freespace'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'readable': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Readable'}, 'size': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Size'}}}}, 'pnpProfileList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'profileName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Profile Name'}, 'discoveryCreated': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Discovery Created'}, 'createdBy': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Created By'}, 'primaryEndpoint': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'port': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Port'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Protocol'}, 'ipv4Address': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Ipv4 Address', 'Elem': {'Schema': {}}}, 'ipv6Address': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Ipv6 Address', 'Elem': {'Schema': {}}}, 'fqdn': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Fqdn'}, 'certificate': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Certificate'}}}}, 'secondaryEndpoint': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'port': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Port'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Protocol'}, 'ipv4Address': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Ipv4 Address', 'Elem': {'Schema': {}}}, 'ipv6Address': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Ipv6 Address', 'Elem': {'Schema': {}}}, 'fqdn': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Fqdn'}, 'certificate': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Certificate'}}}}}}}, 'imageFile': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image File'}, 'httpHeaders': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Key'}, 'value': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Value'}}}}, 'neighborLinks': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'localInterfaceName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Local Interface Name'}, 'localShortInterfaceName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Local Short Interface Name'}, 'localMacAddress': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Local Mac Address'}, 'remoteInterfaceName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Interface Name'}, 'remoteShortInterfaceName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Short Interface Name'}, 'remoteMacAddress': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Mac Address'}, 'remoteDeviceName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Device Name'}, 'remotePlatform': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Platform'}, 'remoteVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Remote Version'}}}}, 'lastSyncTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Last Sync Time'}, 'ipInterfaces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'status': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Status'}, 'macAddress': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Mac Address'}, 'ipv4Address': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Ipv4 Address', 'Elem': {'Schema': {}}}, 'ipv6AddressList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Description': 'Ipv6 Address List', 'Elem': {'Schema': {}}}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'hostname': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Hostname'}, 'authStatus': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Auth Status'}, 'stackInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'supportsStackWorkflows': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Supports Stack Workflows'}, 'isFullRing': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is Full Ring'}, 'stackMemberList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'serialNumber': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Serial Number'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'role': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Role'}, 'macAddress': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Mac Address'}, 'pid': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Pid'}, 'licenseLevel': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'License Level'}, 'licenseType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'License Type'}, 'sudiSerialNumber': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Sudi Serial Number'}, 'hardwareVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Hardware Version'}, 'stackNumber': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Stack Number'}, 'softwareVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Software Version'}, 'priority': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Priority'}}}}, 'stackRingProtocol': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Stack Ring Protocol'}, 'validLicenseLevels': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Valid License Levels'}, 'totalMemberCount': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Total Member Count'}}}}, 'reloadRequested': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Reload Requested'}, 'addedOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Added On'}, 'siteId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Id'}, 'aaaCredentials': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'password': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Password', 'Sensitive': 'true'}, 'username': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Username'}}}}, 'userMicNumbers': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'User Mic Numbers'}, 'userSudiSerialNos': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'User Sudi Serial Nos'}, 'addnMacAddrs': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Addn Mac Addrs'}, 'preWorkflowCliOuputs': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'cli': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Cli'}, 'cliOutput': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Cli Output'}}}}, 'tags': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Tags', 'Elem': {'Schema': {}}}, 'sudiRequired': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Sudi Required'}, 'smartAccountId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Smart Account Id'}, 'virtualAccountId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Virtual Account Id'}, 'populateInventory': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Populate Inventory'}, 'siteName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Site Name'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'systemResetWorkflow': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'lastupdateOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Lastupdate On'}, 'imageId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image Id'}, 'currTaskIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Task Idx'}, 'addedOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Added On'}, 'tasks': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'currWorkItemIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Work Item Idx'}, 'taskSeqNo': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Task Seq No'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'workItemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'command': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Command'}, 'outputStr': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Output Str'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}}}}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'addToInventory': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Add To Inventory'}, 'instanceType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instance Type'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'execTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Exec Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'useState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Use State'}, 'configId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config Id'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'version': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Version'}, 'tenantId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Tenant Id'}}}}, 'systemWorkflow': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'lastupdateOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Lastupdate On'}, 'imageId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image Id'}, 'currTaskIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Task Idx'}, 'addedOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Added On'}, 'tasks': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'currWorkItemIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Work Item Idx'}, 'taskSeqNo': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Task Seq No'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'workItemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'command': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Command'}, 'outputStr': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Output Str'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}}}}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'addToInventory': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Add To Inventory'}, 'instanceType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instance Type'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'execTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Exec Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'useState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Use State'}, 'configId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config Id'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'version': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Version'}, 'tenantId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Tenant Id'}}}}, 'workflow': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'_id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Id'}, 'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description'}, 'lastupdateOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Lastupdate On'}, 'imageId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Image Id'}, 'currTaskIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Task Idx'}, 'addedOn': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Added On'}, 'tasks': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'currWorkItemIdx': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Curr Work Item Idx'}, 'taskSeqNo': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Task Seq No'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'workItemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'command': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Command'}, 'outputStr': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Output Str'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}}}}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'addToInventory': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Add To Inventory'}, 'instanceType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instance Type'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'execTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Exec Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'useState': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Use State'}, 'configId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config Id'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}, 'version': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Version'}, 'tenantId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Tenant Id'}}}}, 'runSummaryList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'details': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Details'}, 'historyTaskInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'type': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type'}, 'workItemList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'state': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'State'}, 'command': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Command'}, 'outputStr': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Output Str'}, 'endTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'End Time'}, 'startTime': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Start Time'}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}}}}, 'timeTaken': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Time Taken'}, 'addnDetails': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Key'}, 'value': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Value'}}}}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name'}}}}, 'errorFlag': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Error Flag'}, 'timestamp': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Timestamp'}}}}, 'workflowParameters': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'topOfStackSerialNumber': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Top Of Stack Serial Number'}, 'licenseLevel': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'License Level'}, 'licenseType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'License Type'}, 'configList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'configParameters': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Key'}, 'value': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Value'}}}}, 'configId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config Id'}}}}}}}, 'dayZeroConfig': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'config': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Config'}}}}, 'dayZeroConfigPreview': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Day Zero Config Preview', 'Elem': {'Schema': {}}}, 'version': {'Computed': 'true', 'Type': 'schema.TypeFloat', 'Description': 'Version'}, 'tenantId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Tenant Id'}}}}}, 'metadata': {'item': {'operation_id': [['AddDevice', 'UpdateDevice'], 'GetDeviceById'], 'new_flat_structure': [[{'RequestDeviceOnboardingPnpAddDevice': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfo'}, {'name': 'runSummaryList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceRunSummaryList'}, {'name': 'systemResetWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow'}, {'name': 'systemWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceSystemWorkflow'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'workflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceWorkflow'}, {'name': 'workflowParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceWorkflowParameters'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfo': {'type': 'obj', 'data': [{'name': 'aaaCredentials', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoAaaCredentials'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'addnMacAddrs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'agentType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedSudiSerialNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'capabilitiesSupported', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'cmState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceSudiSerialNos', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'deviceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'featuresSupported', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'fileSystemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList'}, {'name': 'firstContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'hostname', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'httpHeaders', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoHttpHeaders'}, {'name': 'imageFile', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'imageVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipInterfaces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfaces'}, {'name': 'lastContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastSyncTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastUpdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'mode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'neighborLinks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks'}, {'name': 'onbState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pnpProfileList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList'}, {'name': 'populateInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'preWorkflowCliOuputs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs'}, {'name': 'projectId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reloadRequested', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'smartAccountId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'source', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stack', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiRequired', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags'}, {'name': 'userSudiSerialNos', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'virtualAccountId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoAaaCredentials': {'type': 'obj', 'data': [{'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList': {'type': 'obj', 'data': [{'name': 'freespace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'readable', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'size', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'writeable', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoHttpHeaders': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfaces': {'type': 'obj', 'data': [{'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfacesIpv4Address'}, {'name': 'ipv6AddressList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfacesIpv6AddressList'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'status', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfacesIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoIpInterfacesIpv6AddressList': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation': {'type': 'obj', 'data': [{'name': 'address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'altitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'latitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'longitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'siteId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks': {'type': 'obj', 'data': [{'name': 'localInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localMacAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localShortInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteDeviceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteMacAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remotePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteShortInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList': {'type': 'obj', 'data': [{'name': 'createdBy', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCreated', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'primaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint'}, {'name': 'profileName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'secondaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint': {'type': 'obj', 'data': [{'name': 'certificate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fqdn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv6Address'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint': {'type': 'obj', 'data': [{'name': 'certificate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fqdn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv6Address'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs': {'type': 'obj', 'data': [{'name': 'cli', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'cliOutput', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo': {'type': 'obj', 'data': [{'name': 'isFullRing', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackMemberList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList'}, {'name': 'stackRingProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'supportsStackWorkflows', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'totalMemberCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'validLicenseLevels', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList': {'type': 'obj', 'data': [{'name': 'hardwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'priority', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceRunSummaryList': {'type': 'obj', 'data': [{'name': 'details', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorFlag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'historyTaskInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo'}, {'name': 'timestamp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo': {'type': 'obj', 'data': [{'name': 'addnDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflowParameters': {'type': 'obj', 'data': [{'name': 'configList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList'}, {'name': 'licenseLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'topOfStackSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList': {'type': 'obj', 'data': [{'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestDeviceOnboardingPnpUpdateDevice': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo'}, {'name': 'runSummaryList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList'}, {'name': 'systemResetWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow'}, {'name': 'systemWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'workflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceWorkflow'}, {'name': 'workflowParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo': {'type': 'obj', 'data': [{'name': 'aaaCredentials', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAaaCredentials'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'addnMacAddrs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'agentType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedSudiSerialNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'capabilitiesSupported', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'cmState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceSudiSerialNos', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'deviceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'featuresSupported', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'fileSystemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList'}, {'name': 'firstContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'hostname', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'httpHeaders', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHttpHeaders'}, {'name': 'imageFile', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'imageVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipInterfaces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfaces'}, {'name': 'lastContact', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastSyncTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'lastUpdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'mode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'neighborLinks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks'}, {'name': 'onbState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pnpProfileList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList'}, {'name': 'populateInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'preWorkflowCliOuputs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs'}, {'name': 'projectId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'reloadRequested', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'smartAccountId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'source', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stack', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiRequired', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags'}, {'name': 'userSudiSerialNos', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'virtualAccountId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAaaCredentials': {'type': 'obj', 'data': [{'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList': {'type': 'obj', 'data': [{'name': 'freespace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'readable', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'size', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'writeable', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHttpHeaders': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfaces': {'type': 'obj', 'data': [{'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfacesIpv4Address'}, {'name': 'ipv6AddressList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfacesIpv6AddressList'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'status', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfacesIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIpInterfacesIpv6AddressList': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation': {'type': 'obj', 'data': [{'name': 'address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'altitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'latitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'longitude', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'siteId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks': {'type': 'obj', 'data': [{'name': 'localInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localMacAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localShortInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteDeviceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteMacAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remotePlatform', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteShortInterfaceName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList': {'type': 'obj', 'data': [{'name': 'createdBy', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCreated', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'primaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint'}, {'name': 'profileName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'secondaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint': {'type': 'obj', 'data': [{'name': 'certificate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fqdn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv6Address'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint': {'type': 'obj', 'data': [{'name': 'certificate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fqdn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv6Address'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs': {'type': 'obj', 'data': [{'name': 'cli', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'cliOutput', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo': {'type': 'obj', 'data': [{'name': 'isFullRing', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackMemberList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList'}, {'name': 'stackRingProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'supportsStackWorkflows', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'totalMemberCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'validLicenseLevels', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList': {'type': 'obj', 'data': [{'name': 'hardwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'priority', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList': {'type': 'obj', 'data': [{'name': 'details', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'errorFlag', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'historyTaskInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo'}, {'name': 'timestamp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo': {'type': 'obj', 'data': [{'name': 'addnDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'addToInventory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'execTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'imageId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks'}, {'name': 'tenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'useState', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks': {'type': 'obj', 'data': [{'name': 'currWorkItemIdx', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'taskSeqNo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'command', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputStr', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'startTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'state', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeTaken', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters': {'type': 'obj', 'data': [{'name': 'configList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList'}, {'name': 'licenseLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'topOfStackSerialNumber', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList': {'type': 'obj', 'data': [{'name': 'configId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters': {'type': 'obj', 'data': [{'name': 'key', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], {'ResponseDeviceOnboardingPnpGetDeviceById': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfo'}, {'name': 'systemResetWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflow'}, {'name': 'systemWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflow'}, {'name': 'workflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflow'}, {'name': 'runSummaryList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryList'}, {'name': 'workflowParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParameters'}, {'name': 'dayZeroConfig', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfig'}, {'name': 'dayZeroConfigPreview', 'description': 'Day Zero Config Preview', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfigPreview'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfo': {'type': 'obj', 'data': [{'name': 'source', 'description': 'Source', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': 'Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stack', 'description': 'Stack', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'mode', 'description': 'Mode', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoLocation'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'onbState', 'description': 'Onb State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedMicNumber', 'description': 'Authenticated Mic Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedSudiSerialNo', 'description': 'Authenticated Sudi Serial No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'capabilitiesSupported', 'description': 'Capabilities Supported', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'featuresSupported', 'description': 'Features Supported', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'cmState', 'description': 'Cm State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'firstContact', 'description': 'First Contact', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'lastContact', 'description': 'Last Contact', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': 'Pid', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceSudiSerialNos', 'description': 'Device Sudi Serial Nos', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'lastUpdateOn', 'description': 'Last Update On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workflowId', 'description': 'Workflow Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowName', 'description': 'Workflow Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectId', 'description': 'Project Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceType', 'description': 'Device Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'agentType', 'description': 'Agent Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'imageVersion', 'description': 'Image Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fileSystemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoFileSystemList'}, {'name': 'pnpProfileList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileList'}, {'name': 'imageFile', 'description': 'Image File', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'httpHeaders', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoHttpHeaders'}, {'name': 'neighborLinks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoNeighborLinks'}, {'name': 'lastSyncTime', 'description': 'Last Sync Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'ipInterfaces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfaces'}, {'name': 'hostname', 'description': 'Hostname', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authStatus', 'description': 'Auth Status', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfo'}, {'name': 'reloadRequested', 'description': 'Reload Requested', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'siteId', 'description': 'Site Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'aaaCredentials', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoAaaCredentials'}, {'name': 'userMicNumbers', 'description': 'User Mic Numbers', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'userSudiSerialNos', 'description': 'User Sudi Serial Nos', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'addnMacAddrs', 'description': 'Addn Mac Addrs', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'preWorkflowCliOuputs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPreWorkflowCliOuputs'}, {'name': 'tags', 'description': 'Tags', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoTags'}, {'name': 'sudiRequired', 'description': 'Sudi Required', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'smartAccountId', 'description': 'Smart Account Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'virtualAccountId', 'description': 'Virtual Account Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'populateInventory', 'description': 'Populate Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'siteName', 'description': 'Site Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoLocation': {'type': 'obj', 'data': [{'name': 'siteId', 'description': 'Site Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'address', 'description': 'Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'latitude', 'description': 'Latitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'longitude', 'description': 'Longitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'altitude', 'description': 'Altitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoFileSystemList': {'type': 'obj', 'data': [{'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'writeable', 'description': 'Writeable', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'freespace', 'description': 'Freespace', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'readable', 'description': 'Readable', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'size', 'description': 'Size', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileList': {'type': 'obj', 'data': [{'name': 'profileName', 'description': 'Profile Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCreated', 'description': 'Discovery Created', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'createdBy', 'description': 'Created By', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'primaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpoint'}, {'name': 'secondaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpoint'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpoint': {'type': 'obj', 'data': [{'name': 'port', 'description': 'Port', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': 'Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': 'Ipv6 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv6Address'}, {'name': 'fqdn', 'description': 'Fqdn', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'certificate', 'description': 'Certificate', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpoint': {'type': 'obj', 'data': [{'name': 'port', 'description': 'Port', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': 'Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': 'Ipv6 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv6Address'}, {'name': 'fqdn', 'description': 'Fqdn', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'certificate', 'description': 'Certificate', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoHttpHeaders': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoNeighborLinks': {'type': 'obj', 'data': [{'name': 'localInterfaceName', 'description': 'Local Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localShortInterfaceName', 'description': 'Local Short Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localMacAddress', 'description': 'Local Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteInterfaceName', 'description': 'Remote Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteShortInterfaceName', 'description': 'Remote Short Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteMacAddress', 'description': 'Remote Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteDeviceName', 'description': 'Remote Device Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remotePlatform', 'description': 'Remote Platform', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteVersion', 'description': 'Remote Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfaces': {'type': 'obj', 'data': [{'name': 'status', 'description': 'Status', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv4Address'}, {'name': 'ipv6AddressList', 'description': 'Ipv6 Address List', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv6AddressList'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv6AddressList': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfo': {'type': 'obj', 'data': [{'name': 'supportsStackWorkflows', 'description': 'Supports Stack Workflows', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'isFullRing', 'description': 'Is Full Ring', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackMemberList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfoStackMemberList'}, {'name': 'stackRingProtocol', 'description': 'Stack Ring Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'validLicenseLevels', 'description': 'Valid License Levels', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'totalMemberCount', 'description': 'Total Member Count', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfoStackMemberList': {'type': 'obj', 'data': [{'name': 'serialNumber', 'description': 'Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': 'Role', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': 'Pid', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': 'License Level', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': 'License Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiSerialNumber', 'description': 'Sudi Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'hardwareVersion', 'description': 'Hardware Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackNumber', 'description': 'Stack Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'softwareVersion', 'description': 'Software Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'priority', 'description': 'Priority', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoAaaCredentials': {'type': 'obj', 'data': [{'name': 'password', 'description': 'Password', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'username', 'description': 'Username', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPreWorkflowCliOuputs': {'type': 'obj', 'data': [{'name': 'cli', 'description': 'Cli', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'cliOutput', 'description': 'Cli Output', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoTags': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryList': {'type': 'obj', 'data': [{'name': 'details', 'description': 'Details', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'historyTaskInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfo'}, {'name': 'errorFlag', 'description': 'Error Flag', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'timestamp', 'description': 'Timestamp', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfo': {'type': 'obj', 'data': [{'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addnDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoAddnDetails'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoAddnDetails': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParameters': {'type': 'obj', 'data': [{'name': 'topOfStackSerialNumber', 'description': 'Top Of Stack Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': 'License Level', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': 'License Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigList': {'type': 'obj', 'data': [{'name': 'configParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigListConfigParameters'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigListConfigParameters': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfig': {'type': 'obj', 'data': [{'name': 'config', 'description': 'Config', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfigPreview': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByID': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfo'}, {'name': 'systemResetWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflow'}, {'name': 'systemWorkflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflow'}, {'name': 'workflow', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflow'}, {'name': 'runSummaryList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryList'}, {'name': 'workflowParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParameters'}, {'name': 'dayZeroConfig', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfig'}, {'name': 'dayZeroConfigPreview', 'description': 'Day Zero Config Preview', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDayZeroConfigPreview'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfo': {'type': 'obj', 'data': [{'name': 'source', 'description': 'Source', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'serialNumber', 'description': 'Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stack', 'description': 'Stack', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'mode', 'description': 'Mode', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'location', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoLocation'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'onbState', 'description': 'Onb State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedMicNumber', 'description': 'Authenticated Mic Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authenticatedSudiSerialNo', 'description': 'Authenticated Sudi Serial No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'capabilitiesSupported', 'description': 'Capabilities Supported', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'featuresSupported', 'description': 'Features Supported', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'cmState', 'description': 'Cm State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'firstContact', 'description': 'First Contact', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'lastContact', 'description': 'Last Contact', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': 'Pid', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceSudiSerialNos', 'description': 'Device Sudi Serial Nos', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'lastUpdateOn', 'description': 'Last Update On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workflowId', 'description': 'Workflow Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workflowName', 'description': 'Workflow Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectId', 'description': 'Project Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceType', 'description': 'Device Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'agentType', 'description': 'Agent Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'imageVersion', 'description': 'Image Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'fileSystemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoFileSystemList'}, {'name': 'pnpProfileList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileList'}, {'name': 'imageFile', 'description': 'Image File', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'httpHeaders', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoHttpHeaders'}, {'name': 'neighborLinks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoNeighborLinks'}, {'name': 'lastSyncTime', 'description': 'Last Sync Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'ipInterfaces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfaces'}, {'name': 'hostname', 'description': 'Hostname', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'authStatus', 'description': 'Auth Status', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfo'}, {'name': 'reloadRequested', 'description': 'Reload Requested', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'siteId', 'description': 'Site Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'aaaCredentials', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoAaaCredentials'}, {'name': 'userMicNumbers', 'description': 'User Mic Numbers', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'userSudiSerialNos', 'description': 'User Sudi Serial Nos', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'addnMacAddrs', 'description': 'Addn Mac Addrs', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'preWorkflowCliOuputs', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPreWorkflowCliOuputs'}, {'name': 'tags', 'description': 'Tags', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoTags'}, {'name': 'sudiRequired', 'description': 'Sudi Required', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'smartAccountId', 'description': 'Smart Account Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'virtualAccountId', 'description': 'Virtual Account Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'populateInventory', 'description': 'Populate Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'siteName', 'description': 'Site Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoLocation': {'type': 'obj', 'data': [{'name': 'siteId', 'description': 'Site Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'address', 'description': 'Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'latitude', 'description': 'Latitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'longitude', 'description': 'Longitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'altitude', 'description': 'Altitude', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoFileSystemList': {'type': 'obj', 'data': [{'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'writeable', 'description': 'Writeable', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'freespace', 'description': 'Freespace', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'readable', 'description': 'Readable', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'size', 'description': 'Size', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileList': {'type': 'obj', 'data': [{'name': 'profileName', 'description': 'Profile Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCreated', 'description': 'Discovery Created', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'createdBy', 'description': 'Created By', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'primaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpoint'}, {'name': 'secondaryEndpoint', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpoint'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpoint': {'type': 'obj', 'data': [{'name': 'port', 'description': 'Port', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': 'Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': 'Ipv6 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListPrimaryEndpointIpv6Address'}, {'name': 'fqdn', 'description': 'Fqdn', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'certificate', 'description': 'Certificate', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpoint': {'type': 'obj', 'data': [{'name': 'port', 'description': 'Port', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': 'Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv4Address'}, {'name': 'ipv6Address', 'description': 'Ipv6 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoPnpProfileListSecondaryEndpointIpv6Address'}, {'name': 'fqdn', 'description': 'Fqdn', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'certificate', 'description': 'Certificate', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv6Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoHTTPHeaders': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoNeighborLinks': {'type': 'obj', 'data': [{'name': 'localInterfaceName', 'description': 'Local Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localShortInterfaceName', 'description': 'Local Short Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'localMacAddress', 'description': 'Local Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteInterfaceName', 'description': 'Remote Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteShortInterfaceName', 'description': 'Remote Short Interface Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteMacAddress', 'description': 'Remote Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteDeviceName', 'description': 'Remote Device Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remotePlatform', 'description': 'Remote Platform', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'remoteVersion', 'description': 'Remote Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfaces': {'type': 'obj', 'data': [{'name': 'status', 'description': 'Status', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4Address', 'description': 'Ipv4 Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv4Address'}, {'name': 'ipv6AddressList', 'description': 'Ipv6 Address List', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoIpInterfacesIpv6AddressList'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv4Address': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv6AddressList': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfo': {'type': 'obj', 'data': [{'name': 'supportsStackWorkflows', 'description': 'Supports Stack Workflows', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'isFullRing', 'description': 'Is Full Ring', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'stackMemberList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdDeviceInfoStackInfoStackMemberList'}, {'name': 'stackRingProtocol', 'description': 'Stack Ring Protocol', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'validLicenseLevels', 'description': 'Valid License Levels', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'totalMemberCount', 'description': 'Total Member Count', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfoStackMemberList': {'type': 'obj', 'data': [{'name': 'serialNumber', 'description': 'Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'role', 'description': 'Role', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'macAddress', 'description': 'Mac Address', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pid', 'description': 'Pid', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': 'License Level', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': 'License Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sudiSerialNumber', 'description': 'Sudi Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'hardwareVersion', 'description': 'Hardware Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'stackNumber', 'description': 'Stack Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'softwareVersion', 'description': 'Software Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'priority', 'description': 'Priority', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoAAACredentials': {'type': 'obj', 'data': [{'name': 'password', 'description': 'Password', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'username', 'description': 'Username', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPreWorkflowCliOuputs': {'type': 'obj', 'data': [{'name': 'cli', 'description': 'Cli', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'cliOutput', 'description': 'Cli Output', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoTags': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemResetWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdSystemWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflow': {'type': 'obj', 'data': [{'name': '_id', 'description': 'Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastupdateOn', 'description': 'Lastupdate On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'imageId', 'description': 'Image Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currTaskIdx', 'description': 'Curr Task Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addedOn', 'description': 'Added On', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tasks', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasks'}, {'name': 'addToInventory', 'description': 'Add To Inventory', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'instanceType', 'description': 'Instance Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'execTime', 'description': 'Exec Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'useState', 'description': 'Use State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'version', 'description': 'Version', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'tenantId', 'description': 'Tenant Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasks': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'currWorkItemIdx', 'description': 'Curr Work Item Idx', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'taskSeqNo', 'description': 'Task Seq No', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowTasksWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasksWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryList': {'type': 'obj', 'data': [{'name': 'details', 'description': 'Details', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'historyTaskInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfo'}, {'name': 'errorFlag', 'description': 'Error Flag', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'timestamp', 'description': 'Timestamp', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfo': {'type': 'obj', 'data': [{'name': 'type', 'description': 'Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'workItemList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoWorkItemList'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'addnDetails', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdRunSummaryListHistoryTaskInfoAddnDetails'}, {'name': 'name', 'description': 'Name', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoWorkItemList': {'type': 'obj', 'data': [{'name': 'state', 'description': 'State', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'command', 'description': 'Command', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputStr', 'description': 'Output Str', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'endTime', 'description': 'End Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'startTime', 'description': 'Start Time', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'timeTaken', 'description': 'Time Taken', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoAddnDetails': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParameters': {'type': 'obj', 'data': [{'name': 'topOfStackSerialNumber', 'description': 'Top Of Stack Serial Number', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseLevel', 'description': 'License Level', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'licenseType', 'description': 'License Type', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'configList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigList'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigList': {'type': 'obj', 'data': [{'name': 'configParameters', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseDeviceOnboardingPnpGetDeviceByIdWorkflowParametersConfigListConfigParameters'}, {'name': 'configId', 'description': 'Config Id', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigListConfigParameters': {'type': 'obj', 'data': [{'name': 'key', 'description': 'Key', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'value', 'description': 'Value', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfig': {'type': 'obj', 'data': [{'name': 'config', 'description': 'Config', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfigPreview': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': [['RequestDeviceOnboardingPnpAddDevice', 'RequestDeviceOnboardingPnpUpdateDevice'], 'ResponseDeviceOnboardingPnpGetDeviceById'], 'access_list': [[[], []], []]}}}
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
						"day_zero_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config": &schema.Schema{
										Description: `Config`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"day_zero_config_preview": &schema.Schema{
							Description: `Day Zero Config Preview`,
							Type:        schema.TypeList,
							Computed:    true,
						},
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_credentials": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"password": &schema.Schema{
													Description: `Password`,
													Type:        schema.TypeString,
													Sensitive:   true,
													Computed:    true,
												},
												"username": &schema.Schema{
													Description: `Username`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"added_on": &schema.Schema{
										Description: `Added On`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"addn_mac_addrs": &schema.Schema{
										Description: `Addn Mac Addrs`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"agent_type": &schema.Schema{
										Description: `Agent Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"auth_status": &schema.Schema{
										Description: `Auth Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"authenticated_mic_number": &schema.Schema{
										Description: `Authenticated Mic Number`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"authenticated_sudi_serial_no": &schema.Schema{
										Description: `Authenticated Sudi Serial No`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"capabilities_supported": &schema.Schema{
										Description: `Capabilities Supported`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cm_state": &schema.Schema{
										Description: `Cm State`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Description: `Device Sudi Serial Nos`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_type": &schema.Schema{
										Description: `Device Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"features_supported": &schema.Schema{
										Description: `Features Supported`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"file_system_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"freespace": &schema.Schema{
													Description: `Freespace`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"readable": &schema.Schema{
													Description: `Readable`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"size": &schema.Schema{
													Description: `Size`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"writeable": &schema.Schema{
													Description: `Writeable`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"first_contact": &schema.Schema{
										Description: `First Contact`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"hostname": &schema.Schema{
										Description: `Hostname`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"http_headers": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"image_file": &schema.Schema{
										Description: `Image File`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"image_version": &schema.Schema{
										Description: `Image Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ip_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4_address": &schema.Schema{
													Description: `Ipv4 Address`,
													Type:        schema.TypeList,
													Computed:    true,
												},
												"ipv6_address_list": &schema.Schema{
													Description: `Ipv6 Address List`,
													Type:        schema.TypeList,
													Computed:    true,
												},
												"mac_address": &schema.Schema{
													Description: `Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"status": &schema.Schema{
													Description: `Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"last_contact": &schema.Schema{
										Description: `Last Contact`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"last_sync_time": &schema.Schema{
										Description: `Last Sync Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"last_update_on": &schema.Schema{
										Description: `Last Update On`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Description: `Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"altitude": &schema.Schema{
													Description: `Altitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"latitude": &schema.Schema{
													Description: `Latitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"longitude": &schema.Schema{
													Description: `Longitude`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"site_id": &schema.Schema{
													Description: `Site Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Description: `Mac Address`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"mode": &schema.Schema{
										Description: `Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"neighbor_links": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"local_interface_name": &schema.Schema{
													Description: `Local Interface Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"local_mac_address": &schema.Schema{
													Description: `Local Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"local_short_interface_name": &schema.Schema{
													Description: `Local Short Interface Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_device_name": &schema.Schema{
													Description: `Remote Device Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_interface_name": &schema.Schema{
													Description: `Remote Interface Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_mac_address": &schema.Schema{
													Description: `Remote Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_platform": &schema.Schema{
													Description: `Remote Platform`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_short_interface_name": &schema.Schema{
													Description: `Remote Short Interface Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"remote_version": &schema.Schema{
													Description: `Remote Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"onb_state": &schema.Schema{
										Description: `Onb State`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"pid": &schema.Schema{
										Description: `Pid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"pnp_profile_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"created_by": &schema.Schema{
													Description: `Created By`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"discovery_created": &schema.Schema{
													Description: `Discovery Created`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"primary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Description: `Certificate`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"fqdn": &schema.Schema{
																Description: `Fqdn`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"ipv4_address": &schema.Schema{
																Description: `Ipv4 Address`,
																Type:        schema.TypeList,
																Computed:    true,
															},
															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeList,
																Computed:    true,
															},
															"port": &schema.Schema{
																Description: `Port`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"profile_name": &schema.Schema{
													Description: `Profile Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"secondary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Description: `Certificate`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"fqdn": &schema.Schema{
																Description: `Fqdn`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"ipv4_address": &schema.Schema{
																Description: `Ipv4 Address`,
																Type:        schema.TypeList,
																Computed:    true,
															},
															"ipv6_address": &schema.Schema{
																Description: `Ipv6 Address`,
																Type:        schema.TypeList,
																Computed:    true,
															},
															"port": &schema.Schema{
																Description: `Port`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"populate_inventory": &schema.Schema{
										Description: `Populate Inventory`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"pre_workflow_cli_ouputs": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cli": &schema.Schema{
													Description: `Cli`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"cli_output": &schema.Schema{
													Description: `Cli Output`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"project_id": &schema.Schema{
										Description: `Project Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"project_name": &schema.Schema{
										Description: `Project Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"reload_requested": &schema.Schema{
										Description: `Reload Requested`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"serial_number": &schema.Schema{
										Description: `Serial Number`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"site_name": &schema.Schema{
										Description: `Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"smart_account_id": &schema.Schema{
										Description: `Smart Account Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"source": &schema.Schema{
										Description: `Source`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"stack": &schema.Schema{
										Description: `Stack`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"stack_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"is_full_ring": &schema.Schema{
													Description: `Is Full Ring`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hardware_version": &schema.Schema{
																Description: `Hardware Version`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"license_level": &schema.Schema{
																Description: `License Level`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"license_type": &schema.Schema{
																Description: `License Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"pid": &schema.Schema{
																Description: `Pid`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"priority": &schema.Schema{
																Description: `Priority`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"role": &schema.Schema{
																Description: `Role`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"serial_number": &schema.Schema{
																Description: `Serial Number`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"software_version": &schema.Schema{
																Description: `Software Version`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"stack_number": &schema.Schema{
																Description: `Stack Number`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"sudi_serial_number": &schema.Schema{
																Description: `Sudi Serial Number`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Description: `Stack Ring Protocol`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"supports_stack_workflows": &schema.Schema{
													Description: `Supports Stack Workflows`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"total_member_count": &schema.Schema{
													Description: `Total Member Count`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"valid_license_levels": &schema.Schema{
													Description: `Valid License Levels`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"state": &schema.Schema{
										Description: `State`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"sudi_required": &schema.Schema{
										Description: `Sudi Required`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tags": &schema.Schema{
										Description: `Tags`,
										Type:        schema.TypeList,
										Computed:    true,
									},
									"user_mic_numbers": &schema.Schema{
										Description: `User Mic Numbers`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"user_sudi_serial_nos": &schema.Schema{
										Description: `User Sudi Serial Nos`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_account_id": &schema.Schema{
										Description: `Virtual Account Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"workflow_id": &schema.Schema{
										Description: `Workflow Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"workflow_name": &schema.Schema{
										Description: `Workflow Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"run_summary_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"details": &schema.Schema{
										Description: `Details`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"error_flag": &schema.Schema{
										Description: `Error Flag`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"history_task_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addn_details": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Description: `Key`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
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
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},
						"system_reset_workflow": &schema.Schema{
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
						"system_workflow": &schema.Schema{
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
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"workflow": &schema.Schema{
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
						"workflow_parameters": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Description: `Key`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"license_level": &schema.Schema{
										Description: `License Level`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"license_type": &schema.Schema{
										Description: `License Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"top_of_stack_serial_number": &schema.Schema{
										Description: `Top Of Stack Serial Number`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
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
						"device_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_credentials": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"password": &schema.Schema{
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"username": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"added_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"addn_mac_addrs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"agent_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"authenticated_sudi_serial_no": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"capabilities_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cm_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"features_supported": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"file_system_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"freespace": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"readable": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"size": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"writeable": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"first_contact": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"http_headers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"image_file": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"image_version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_interfaces": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4_address": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
												},
												"ipv6_address_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
												},
												"mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"last_contact": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"last_sync_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"last_update_on": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"location": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"altitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"latitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"longitude": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"site_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"neighbor_links": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"local_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"local_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"local_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_device_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_mac_address": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_platform": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_short_interface_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"remote_version": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"onb_state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"pnp_profile_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"created_by": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"discovery_created": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"primary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"profile_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"secondary_endpoint": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"certificate": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"fqdn": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"ipv4_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
															"ipv6_address": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
															"port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"populate_inventory": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"pre_workflow_cli_ouputs": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cli": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"cli_output": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"project_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"project_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"reload_requested": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"smart_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"source": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"stack": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"stack_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"is_full_ring": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"stack_member_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hardware_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"license_level": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"license_type": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"mac_address": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"pid": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"priority": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"role": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"software_version": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"stack_number": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"state": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"sudi_serial_number": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"stack_ring_protocol": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"supports_stack_workflows": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"total_member_count": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"valid_license_levels": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sudi_required": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
									},
									"user_sudi_serial_nos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_account_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"workflow_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"workflow_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"run_summary_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"details": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"error_flag": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"history_task_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addn_details": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
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
									"timestamp": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"system_reset_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
						"system_workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"workflow": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
						"workflow_parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"license_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"license_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"top_of_stack_serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourcePnpDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPnpDeviceAddDevice(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.DeviceOnboardingPnp.GetDeviceList2(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceOnboardingPnpGetDeviceList2(m, response2, nil)
			item2, err := searchDeviceOnboardingPnpGetDeviceList2(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddDevice(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDevice", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourcePnpDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSort, okSort := resourceMap["sort"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vSerialNumber, okSerialNumber := resourceMap["serial_number"]
	vState, okState := resourceMap["state"]
	vOnbState, okOnbState := resourceMap["onb_state"]
	vCmState, okCmState := resourceMap["cm_state"]
	vName, okName := resourceMap["name"]
	vPid, okPid := resourceMap["pid"]
	vSource, okSource := resourceMap["source"]
	vProjectID, okProjectID := resourceMap["project_id"]
	vWorkflowID, okWorkflowID := resourceMap["workflow_id"]
	vProjectName, okProjectName := resourceMap["project_name"]
	vWorkflowName, okWorkflowName := resourceMap["workflow_name"]
	vSmartAccountID, okSmartAccountID := resourceMap["smart_account_id"]
	vVirtualAccountID, okVirtualAccountID := resourceMap["virtual_account_id"]
	vLastContact, okLastContact := resourceMap["last_contact"]
	vMacAddress, okMacAddress := resourceMap["mac_address"]
	vHostname, okHostname := resourceMap["hostname"]
	vSiteName, okSiteName := resourceMap["site_name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okSerialNumber, okState, okOnbState, okCmState, okName, okPid, okSource, okProjectID, okWorkflowID, okProjectName, okWorkflowName, okSmartAccountID, okVirtualAccountID, okLastContact, okMacAddress, okHostname, okSiteName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceList2")
		queryParams1 := dnacentersdkgo.GetDeviceList2QueryParams{}

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
		if okSerialNumber {
			queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
		}
		if okState {
			queryParams1.State = interfaceToSliceString(vState)
		}
		if okOnbState {
			queryParams1.OnbState = interfaceToSliceString(vOnbState)
		}
		if okCmState {
			queryParams1.CmState = interfaceToSliceString(vCmState)
		}
		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}
		if okPid {
			queryParams1.Pid = interfaceToSliceString(vPid)
		}
		if okSource {
			queryParams1.Source = interfaceToSliceString(vSource)
		}
		if okProjectID {
			queryParams1.ProjectID = interfaceToSliceString(vProjectID)
		}
		if okWorkflowID {
			queryParams1.WorkflowID = interfaceToSliceString(vWorkflowID)
		}
		if okProjectName {
			queryParams1.ProjectName = interfaceToSliceString(vProjectName)
		}
		if okWorkflowName {
			queryParams1.WorkflowName = interfaceToSliceString(vWorkflowName)
		}
		if okSmartAccountID {
			queryParams1.SmartAccountID = interfaceToSliceString(vSmartAccountID)
		}
		if okVirtualAccountID {
			queryParams1.VirtualAccountID = interfaceToSliceString(vVirtualAccountID)
		}
		if okLastContact {
			queryParams1.LastContact = vLastContact.(bool)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okHostname {
			queryParams1.Hostname = vHostname.(string)
		}
		if okSiteName {
			queryParams1.SiteName = vSiteName.(string)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetDeviceList2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceList2", err,
				"Failure at GetDeviceList2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.DeviceOnboardingPnp.GetDeviceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceByID", err,
				"Failure at GetDeviceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourcePnpDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vSort, okSort := resourceMap["sort"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vSerialNumber, okSerialNumber := resourceMap["serial_number"]
	vState, okState := resourceMap["state"]
	vOnbState, okOnbState := resourceMap["onb_state"]
	vCmState, okCmState := resourceMap["cm_state"]
	vName, okName := resourceMap["name"]
	vPid, okPid := resourceMap["pid"]
	vSource, okSource := resourceMap["source"]
	vProjectID, okProjectID := resourceMap["project_id"]
	vWorkflowID, okWorkflowID := resourceMap["workflow_id"]
	vProjectName, okProjectName := resourceMap["project_name"]
	vWorkflowName, okWorkflowName := resourceMap["workflow_name"]
	vSmartAccountID, okSmartAccountID := resourceMap["smart_account_id"]
	vVirtualAccountID, okVirtualAccountID := resourceMap["virtual_account_id"]
	vLastContact, okLastContact := resourceMap["last_contact"]
	vMacAddress, okMacAddress := resourceMap["mac_address"]
	vHostname, okHostname := resourceMap["hostname"]
	vSiteName, okSiteName := resourceMap["site_name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okLimit, okOffset, okSort, okSortOrder, okSerialNumber, okState, okOnbState, okCmState, okName, okPid, okSource, okProjectID, okWorkflowID, okProjectName, okWorkflowName, okSmartAccountID, okVirtualAccountID, okLastContact, okMacAddress, okHostname, okSiteName}
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
		request1 := expandRequestPnpDeviceUpdateDevice(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceOnboardingPnp.UpdateDevice(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDevice", err, restyResp1.String(),
					"Failure at UpdateDevice, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDevice", err,
				"Failure at UpdateDevice, unexpected response", ""))
			return diags
		}
	}

	return resourcePnpDeviceRead(ctx, d, m)
}

func resourcePnpDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestPnpDeviceAddDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceAddDeviceDeviceInfo(ctx, key+".device_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_summary_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_summary_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_summary_list")))) {
		request.RunSummaryList = expandRequestPnpDeviceAddDeviceRunSummaryListArray(ctx, key+".run_summary_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_reset_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_reset_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_reset_workflow")))) {
		request.SystemResetWorkflow = expandRequestPnpDeviceAddDeviceSystemResetWorkflow(ctx, key+".system_reset_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_workflow")))) {
		request.SystemWorkflow = expandRequestPnpDeviceAddDeviceSystemWorkflow(ctx, key+".system_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow")))) {
		request.Workflow = expandRequestPnpDeviceAddDeviceWorkflow(ctx, key+".workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_parameters")))) {
		request.WorkflowParameters = expandRequestPnpDeviceAddDeviceWorkflowParameters(ctx, key+".workflow_parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpDeviceAddDeviceDeviceInfoAAACredentials(ctx, key+".aaa_credentials.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_mac_addrs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_mac_addrs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_mac_addrs")))) {
		request.AddnMacAddrs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".agent_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".agent_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".agent_type")))) {
		request.AgentType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_status")))) {
		request.AuthStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticated_sudi_serial_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) {
		request.AuthenticatedSudiSerialNo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".capabilities_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".capabilities_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".capabilities_supported")))) {
		request.CapabilitiesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cm_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cm_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cm_state")))) {
		request.CmState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) {
		request.DeviceSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".features_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".features_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".features_supported")))) {
		request.FeaturesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_system_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_system_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_system_list")))) {
		request.FileSystemList = expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemListArray(ctx, key+".file_system_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_contact")))) {
		request.FirstContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_headers")))) {
		request.HTTPHeaders = expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeadersArray(ctx, key+".http_headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_file")))) {
		request.ImageFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_version")))) {
		request.ImageVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interfaces")))) {
		request.IPInterfaces = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesArray(ctx, key+".ip_interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_contact")))) {
		request.LastContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync_time")))) {
		request.LastSyncTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_on")))) {
		request.LastUpdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = expandRequestPnpDeviceAddDeviceDeviceInfoLocation(ctx, key+".location.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode")))) {
		request.Mode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbor_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbor_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbor_links")))) {
		request.NeighborLinks = expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinksArray(ctx, key+".neighbor_links", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".onb_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".onb_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".onb_state")))) {
		request.OnbState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pnp_profile_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pnp_profile_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pnp_profile_list")))) {
		request.PnpProfileList = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListArray(ctx, key+".pnp_profile_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_workflow_cli_ouputs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) {
		request.PreWorkflowCliOuputs = expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx, key+".pre_workflow_cli_ouputs", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reload_requested")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reload_requested")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reload_requested")))) {
		request.ReloadRequested = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source")))) {
		request.Source = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_info")))) {
		request.StackInfo = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestPnpDeviceAddDeviceDeviceInfoTags(ctx, key+".tags.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_name")))) {
		request.WorkflowName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoFileSystemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".freespace")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".freespace")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".freespace")))) {
		request.Freespace = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".readable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".readable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".readable")))) {
		request.Readable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".size")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".size")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".size")))) {
		request.Size = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".writeable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".writeable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".writeable")))) {
		request.Writeable = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoHTTPHeaders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_list")))) {
		request.IPv6AddressList = expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx, key+".ipv6_address_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoLocation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".altitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".altitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".altitude")))) {
		request.Altitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoNeighborLinks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_interface_name")))) {
		request.LocalInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_mac_address")))) {
		request.LocalMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_short_interface_name")))) {
		request.LocalShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_device_name")))) {
		request.RemoteDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_interface_name")))) {
		request.RemoteInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_mac_address")))) {
		request.RemoteMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_platform")))) {
		request.RemotePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_short_interface_name")))) {
		request.RemoteShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_version")))) {
		request.RemoteVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".created_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".created_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".created_by")))) {
		request.CreatedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_created")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_created")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_created")))) {
		request.DiscoveryCreated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_endpoint")))) {
		request.PrimaryEndpoint = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx, key+".primary_endpoint.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_endpoint")))) {
		request.SecondaryEndpoint = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx, key+".secondary_endpoint.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoPreWorkflowCliOuputs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli")))) {
		request.Cli = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_output")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_output")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_output")))) {
		request.CliOutput = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_ring_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_ring_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_ring_protocol")))) {
		request.StackRingProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supports_stack_workflows")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supports_stack_workflows")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supports_stack_workflows")))) {
		request.SupportsStackWorkflows = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_member_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_member_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_member_count")))) {
		request.TotalMemberCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_license_levels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_license_levels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_license_levels")))) {
		request.ValidLicenseLevels = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_version")))) {
		request.HardwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_number")))) {
		request.StackNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_serial_number")))) {
		request.SudiSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceDeviceInfoTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".details")))) {
		request.Details = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_flag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_flag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_flag")))) {
		request.ErrorFlag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".history_task_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".history_task_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".history_task_info")))) {
		request.HistoryTaskInfo = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfo(ctx, key+".history_task_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp")))) {
		request.Timestamp = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_details")))) {
		request.AddnDetails = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx, key+".addn_details", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceSystemResetWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemResetWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceSystemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceSystemWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceSystemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceSystemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceAddDeviceWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceAddDeviceWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceAddDeviceWorkflowParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListArray(ctx, key+".config_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_of_stack_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) {
		request.TopOfStackSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowParametersConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceAddDeviceWorkflowParametersConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx, key+".device_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_summary_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_summary_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_summary_list")))) {
		request.RunSummaryList = expandRequestPnpDeviceUpdateDeviceRunSummaryListArray(ctx, key+".run_summary_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_reset_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_reset_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_reset_workflow")))) {
		request.SystemResetWorkflow = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflow(ctx, key+".system_reset_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_workflow")))) {
		request.SystemWorkflow = expandRequestPnpDeviceUpdateDeviceSystemWorkflow(ctx, key+".system_workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow")))) {
		request.Workflow = expandRequestPnpDeviceUpdateDeviceWorkflow(ctx, key+".workflow.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_parameters")))) {
		request.WorkflowParameters = expandRequestPnpDeviceUpdateDeviceWorkflowParameters(ctx, key+".workflow_parameters.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpDeviceUpdateDeviceDeviceInfoAAACredentials(ctx, key+".aaa_credentials.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".added_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".added_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".added_on")))) {
		request.AddedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_mac_addrs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_mac_addrs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_mac_addrs")))) {
		request.AddnMacAddrs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".agent_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".agent_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".agent_type")))) {
		request.AgentType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_status")))) {
		request.AuthStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticated_sudi_serial_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticated_sudi_serial_no")))) {
		request.AuthenticatedSudiSerialNo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".capabilities_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".capabilities_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".capabilities_supported")))) {
		request.CapabilitiesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cm_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cm_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cm_state")))) {
		request.CmState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) {
		request.DeviceSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".features_supported")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".features_supported")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".features_supported")))) {
		request.FeaturesSupported = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_system_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_system_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_system_list")))) {
		request.FileSystemList = expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemListArray(ctx, key+".file_system_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_contact")))) {
		request.FirstContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_headers")))) {
		request.HTTPHeaders = expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeadersArray(ctx, key+".http_headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_file")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_file")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_file")))) {
		request.ImageFile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_version")))) {
		request.ImageVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interfaces")))) {
		request.IPInterfaces = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesArray(ctx, key+".ip_interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_contact")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_contact")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_contact")))) {
		request.LastContact = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync_time")))) {
		request.LastSyncTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_on")))) {
		request.LastUpdateOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = expandRequestPnpDeviceUpdateDeviceDeviceInfoLocation(ctx, key+".location.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode")))) {
		request.Mode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbor_links")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbor_links")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbor_links")))) {
		request.NeighborLinks = expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinksArray(ctx, key+".neighbor_links", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".onb_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".onb_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".onb_state")))) {
		request.OnbState = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pnp_profile_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pnp_profile_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pnp_profile_list")))) {
		request.PnpProfileList = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListArray(ctx, key+".pnp_profile_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_workflow_cli_ouputs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_workflow_cli_ouputs")))) {
		request.PreWorkflowCliOuputs = expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx, key+".pre_workflow_cli_ouputs", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".reload_requested")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".reload_requested")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".reload_requested")))) {
		request.ReloadRequested = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source")))) {
		request.Source = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_info")))) {
		request.StackInfo = expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestPnpDeviceUpdateDeviceDeviceInfoTags(ctx, key+".tags.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_name")))) {
		request.WorkflowName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoFileSystemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".freespace")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".freespace")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".freespace")))) {
		request.Freespace = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".readable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".readable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".readable")))) {
		request.Readable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".size")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".size")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".size")))) {
		request.Size = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".writeable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".writeable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".writeable")))) {
		request.Writeable = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoHTTPHeaders(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfaces(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_list")))) {
		request.IPv6AddressList = expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx, key+".ipv6_address_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoLocation(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".altitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".altitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".altitude")))) {
		request.Altitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoNeighborLinks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_interface_name")))) {
		request.LocalInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_mac_address")))) {
		request.LocalMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_short_interface_name")))) {
		request.LocalShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_device_name")))) {
		request.RemoteDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_interface_name")))) {
		request.RemoteInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_mac_address")))) {
		request.RemoteMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_platform")))) {
		request.RemotePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_short_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_short_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_short_interface_name")))) {
		request.RemoteShortInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_version")))) {
		request.RemoteVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".created_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".created_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".created_by")))) {
		request.CreatedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_created")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_created")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_created")))) {
		request.DiscoveryCreated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_endpoint")))) {
		request.PrimaryEndpoint = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx, key+".primary_endpoint.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_endpoint")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_endpoint")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_endpoint")))) {
		request.SecondaryEndpoint = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx, key+".secondary_endpoint.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate")))) {
		request.Certificate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn")))) {
		request.Fqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_address")))) {
		request.IPv4Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx, key+".ipv4_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address")))) {
		request.IPv6Address = expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx, key+".ipv6_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoPreWorkflowCliOuputs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli")))) {
		request.Cli = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_output")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_output")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_output")))) {
		request.CliOutput = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_ring_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_ring_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_ring_protocol")))) {
		request.StackRingProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supports_stack_workflows")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supports_stack_workflows")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supports_stack_workflows")))) {
		request.SupportsStackWorkflows = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_member_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_member_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_member_count")))) {
		request.TotalMemberCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_license_levels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_license_levels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_license_levels")))) {
		request.ValidLicenseLevels = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_version")))) {
		request.HardwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_number")))) {
		request.StackNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_serial_number")))) {
		request.SudiSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceDeviceInfoTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags {
	var request dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".details")))) {
		request.Details = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".error_flag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".error_flag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".error_flag")))) {
		request.ErrorFlag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".history_task_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".history_task_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".history_task_info")))) {
		request.HistoryTaskInfo = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfo(ctx, key+".history_task_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp")))) {
		request.Timestamp = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".addn_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".addn_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".addn_details")))) {
		request.AddnDetails = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx, key+".addn_details", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_taken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_taken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_taken")))) {
		request.TimeTaken = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".work_item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".work_item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".work_item_list")))) {
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemResetWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceSystemWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceSystemWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflow {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflow{}
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
		request.Tasks = expandRequestPnpDeviceUpdateDeviceWorkflowTasksArray(ctx, key+".tasks", d)
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

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowTasks(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks{}
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
		request.WorkItemList = expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemListArray(ctx, key+".work_item_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowTasksWorkItemList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList{}
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

func expandRequestPnpDeviceUpdateDeviceWorkflowParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListArray(ctx, key+".config_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_of_stack_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) {
		request.TopOfStackSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpDeviceUpdateDeviceWorkflowParametersConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
