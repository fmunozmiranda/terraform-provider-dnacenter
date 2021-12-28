package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDiscovery() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Discovery.

- Stops all the discoveries and removes them

- Stops or starts an existing discovery

- Initiates discovery with the given parameters

- Stops the discovery for the given Discovery ID and removes it. Discovery ID can be obtained using the "Get Discoveries
by range" API.
`,

		CreateContext: resourceDiscoveryCreate,
		ReadContext:   resourceDiscoveryRead,
		UpdateContext: resourceDiscoveryUpdate,
		DeleteContext: resourceDiscoveryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'snmpAuthProtocol': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "SNMP auth protocol. SHA' or 'MD5'\n"}, 'snmpRWCommunity': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Snmp RW community of the devices to be discovered\n'}, 'cdpLevel': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'CDP level to which neighbor devices to be discovered\n'}, 'snmpRwCommunity': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'userNameList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Username of the devices to be discovered\n\nERROR: Different types for param userNameList schema.TypeList schema.TypeString'}, 'discoveryCondition': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'snmpRoCommunityDesc': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'enablePasswordList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Enable Password of the devices to be discovered\n\nERROR: Different types for param enablePasswordList schema.TypeList schema.TypeString'}, 'retry': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Number of times to try establishing connection to device\n'}, 'snmpRoCommunity': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'isAutoCdp': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'httpWriteCredential': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'secure': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Flag for HTTPS\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'password': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'HTTP(S) password\n', 'Sensitive': 'true'}, 'instanceTenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'username': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'HTTP(S) username\n'}, 'comments': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'port': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'HTTP(S) port\n'}, 'credentialType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'updateMgmtIp': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'snmpRwCommunityDesc': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'snmpVersion': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Version of SNMP. v2 or v3\n'}, 'numDevices': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'passwordList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Password of the devices to be discovered\n\nERROR: Different types for param passwordList schema.TypeList schema.TypeString', 'Sensitive': 'true'}, 'snmpROCommunityDesc': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description for Snmp RO community\n'}, 'protocolOrder': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet\n"}, 'parentDiscoveryId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'snmpPrivPassphrase': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Pass phrase for SNMP privacy\n'}, 'discoveryType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "Type of Discovery. 'SINGLE', 'RANGE', 'MULTI RANGE', 'CDP', 'LLDP'\n"}, 'ipFilterList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'IP Addresses of the devices to be filtered out during discovery\n\nERROR: Different types for param ipFilterList schema.TypeList schema.TypeString'}, 'deviceIds': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of the discovery\n'}, 'snmpRWCommunityDesc': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description for Snmp RW community\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'lldpLevel': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'LLDP level to which neighbor devices to be discovered\n'}, 'snmpMode': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'\n"}, 'snmpROCommunity': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Snmp RO community of the devices to be discovered\n'}, 'httpReadCredential': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'secure': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Flag for HTTPS\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'password': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'HTTP(S) password\n', 'Sensitive': 'true'}, 'instanceTenantId': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'username': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'HTTP(S) username\n'}, 'comments': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'port': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'HTTP(S) port\n'}, 'credentialType': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString'}}}}, 'timeOut': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'discoveryStatus': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'attributeInfo': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'netconfPort': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Netconf Port. It will need valid SSH credentials to work\n'}, 'snmpUserName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'SNMP username of the device\n'}, 'snmpPrivProtocol': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "SNMP privacy protocol. 'DES' or 'AES128'\n"}, 'snmpAuthPassphrase': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Auth Pass phrase for SNMP\n'}, 'globalCredentialIdList': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Global Credential Ids to be used for discovery\n'}, 'preferredMgmtIPMethod': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "Preferred Management IP Method.'None' or 'UseLoopBack'. Default is 'None'\n"}, 'ipAddressList': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': "IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE\n"}, 'retryCount': {'Optional': 'true', 'Type': 'schema.TypeInt'}, 'timeout': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Time to wait for device response in seconds\n'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'attributeInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {}}}, 'cdpLevel': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'deviceIds': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'discoveryCondition': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'discoveryStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'discoveryType': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'enablePasswordList': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'globalCredentialIdList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'httpReadCredential': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'comments': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'credentialType': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'password': {'Computed': 'true', 'Type': 'schema.TypeString', 'Sensitive': 'true'}, 'port': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'secure': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'username': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'httpWriteCredential': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'comments': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'credentialType': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceTenantId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'instanceUuid': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'password': {'Computed': 'true', 'Type': 'schema.TypeString', 'Sensitive': 'true'}, 'port': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'secure': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'username': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipAddressList': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipFilterList': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'isAutoCdp': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'lldpLevel': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'netconfPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'numDevices': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'parentDiscoveryId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'passwordList': {'Computed': 'true', 'Type': 'schema.TypeString', 'Sensitive': 'true'}, 'preferredMgmtIPMethod': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocolOrder': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'retryCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'snmpAuthPassphrase': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpAuthProtocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpMode': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpPrivPassphrase': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpPrivProtocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpRoCommunity': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpRoCommunityDesc': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpRwCommunity': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpRwCommunityDesc': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'snmpUserName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'timeOut': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'updateMgmtIp': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'userNameList': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}, 'metadata': {'item': {'operation_id': [['StartDiscovery', 'UpdatesAnExistingDiscoveryBySpecifiedId'], 'GetDiscoveryById'], 'new_flat_structure': [[{'RequestDiscoveryStartDiscovery': {'type': 'obj', 'data': [{'name': 'cdpLevel', 'description': 'CDP level to which neighbor devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'discoveryType', 'description': "Type of Discovery. 'SINGLE', 'RANGE', 'MULTI RANGE', 'CDP', 'LLDP'\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enablePasswordList', 'description': 'Enable Password of the devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'globalCredentialIdList', 'description': 'Global Credential Ids to be used for discovery\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'httpReadCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDiscoveryStartDiscoveryHttpReadCredential'}, {'name': 'httpWriteCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDiscoveryStartDiscoveryHttpWriteCredential'}, {'name': 'ipAddressList', 'description': "IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipFilterList', 'description': 'IP Addresses of the devices to be filtered out during discovery\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'lldpLevel', 'description': 'LLDP level to which neighbor devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of the discovery\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'netconfPort', 'description': 'Netconf Port. It will need valid SSH credentials to work\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'passwordList', 'description': 'Password of the devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'preferredMgmtIPMethod', 'description': "Preferred Management IP Method.'None' or 'UseLoopBack'. Default is 'None'\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocolOrder', 'description': "Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'retry', 'description': 'Number of times to try establishing connection to device\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'snmpAuthPassphrase', 'description': 'Auth Pass phrase for SNMP\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpAuthProtocol', 'description': "SNMP auth protocol. SHA' or 'MD5'\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpMode', 'description': "Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivPassphrase', 'description': 'Pass phrase for SNMP privacy\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivProtocol', 'description': "SNMP privacy protocol. 'DES' or 'AES128'\n", 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpROCommunity', 'description': 'Snmp RO community of the devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpROCommunityDesc', 'description': 'Description for Snmp RO community\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRWCommunity', 'description': 'Snmp RW community of the devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRWCommunityDesc', 'description': 'Description for Snmp RW community\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpUserName', 'description': 'SNMP username of the device\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpVersion', 'description': 'Version of SNMP. v2 or v3\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeout', 'description': 'Time to wait for device response in seconds\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'userNameList', 'description': 'Username of the devices to be discovered\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDiscoveryStartDiscoveryHttpReadCredential': {'type': 'obj', 'data': [{'name': 'password', 'description': 'HTTP(S) password\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': 'HTTP(S) port\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': 'Flag for HTTPS\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': 'HTTP(S) username\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDiscoveryStartDiscoveryHttpWriteCredential': {'type': 'obj', 'data': [{'name': 'password', 'description': 'HTTP(S) password\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': 'HTTP(S) port\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': 'Flag for HTTPS\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': 'HTTP(S) username\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedId': {'type': 'obj', 'data': [{'name': 'attributeInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdAttributeInfo'}, {'name': 'cdpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'deviceIds', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCondition', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enablePasswordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'globalCredentialIdList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'httpReadCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdHttpReadCredential'}, {'name': 'httpWriteCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdHttpWriteCredential'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipAddressList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipFilterList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'isAutoCdp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'lldpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'netconfPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'numDevices', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'parentDiscoveryId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'passwordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'preferredMgmtIPMethod', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocolOrder', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'retryCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'snmpAuthPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpAuthProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpMode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpUserName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeOut', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'updateMgmtIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'userNameList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdAttributeInfo': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdHttpReadCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIdHttpWriteCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], {'ResponseDiscoveryGetDiscoveryByIdResponse': {'type': 'obj', 'data': [{'name': 'attributeInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseAttributeInfo'}, {'name': 'cdpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'deviceIds', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCondition', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enablePasswordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'globalCredentialIdList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'httpReadCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseHttpReadCredential'}, {'name': 'httpWriteCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseHttpWriteCredential'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipAddressList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipFilterList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'isAutoCdp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'lldpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'netconfPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'numDevices', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'parentDiscoveryId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'passwordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'preferredMgmtIPMethod', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocolOrder', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'retryCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'snmpAuthPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpAuthProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpMode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpUserName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeOut', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'updateMgmtIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'userNameList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIdResponseAttributeInfo': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIdResponseHttpReadCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIdResponseHttpWriteCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIDResponse': {'type': 'obj', 'data': [{'name': 'attributeInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseAttributeInfo'}, {'name': 'cdpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'deviceIds', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryCondition', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'discoveryType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'enablePasswordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'globalCredentialIdList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'httpReadCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseHttpReadCredential'}, {'name': 'httpWriteCredential', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseDiscoveryGetDiscoveryByIdResponseHttpWriteCredential'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipAddressList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipFilterList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'isAutoCdp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'lldpLevel', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'netconfPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'numDevices', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'parentDiscoveryId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'passwordList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'preferredMgmtIPMethod', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocolOrder', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'retryCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'snmpAuthPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpAuthProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpMode', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivPassphrase', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpPrivProtocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRoCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunity', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpRwCommunityDesc', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'snmpUserName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'timeOut', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'updateMgmtIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'userNameList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIDResponseAttributeInfo': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIDResponseHTTPReadCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseDiscoveryGetDiscoveryByIDResponseHTTPWriteCredential': {'type': 'obj', 'data': [{'name': 'comments', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'credentialType', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceTenantId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instanceUuid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'password', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'port', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'secure', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'username', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': [['RequestDiscoveryStartDiscovery', 'RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedId'], 'ResponseDiscoveryGetDiscoveryByIdResponse'], 'access_list': [[[], ['attributeInfo', 'cdpLevel', 'deviceIds', 'discoveryCondition', 'discoveryStatus', 'discoveryType', 'enablePasswordList', 'globalCredentialIdList', 'httpReadCredential', 'httpWriteCredential', 'id', 'ipAddressList', 'ipFilterList', 'isAutoCdp', 'lldpLevel', 'name', 'netconfPort', 'numDevices', 'parentDiscoveryId', 'passwordList', 'preferredMgmtIPMethod', 'protocolOrder', 'retryCount', 'snmpAuthPassphrase', 'snmpAuthProtocol', 'snmpMode', 'snmpPrivPassphrase', 'snmpPrivProtocol', 'snmpRoCommunity', 'snmpRoCommunityDesc', 'snmpRwCommunity', 'snmpRwCommunityDesc', 'snmpUserName', 'timeOut', 'updateMgmtIp', 'userNameList']], ['response']]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
						},
						"cdp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"discovery_condition": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"discovery_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_password_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"global_credential_id_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"secure": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"secure": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_filter_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_auto_cdp": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"lldp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_list": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"preferred_mgmt_ipmethod": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol_order": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"update_mgmt_ip": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name_list": &schema.Schema{
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

						"attribute_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
						},
						"cdp_level": &schema.Schema{
							Description: `CDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"discovery_condition": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"discovery_type": &schema.Schema{
							Description: `Type of Discovery. 'SINGLE', 'RANGE', 'MULTI RANGE', 'CDP', 'LLDP'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_password_list": &schema.Schema{
							Description: `Enable Password of the devices to be discovered

ERROR: Different types for param enablePasswordList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"global_credential_id_list": &schema.Schema{
							Description: `Global Credential Ids to be used for discovery
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"password": &schema.Schema{
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"credential_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"password": &schema.Schema{
										Description: `HTTP(S) password
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"port": &schema.Schema{
										Description: `HTTP(S) port
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"secure": &schema.Schema{
										Description: `Flag for HTTPS
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"username": &schema.Schema{
										Description: `HTTP(S) username
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address_list": &schema.Schema{
							Description: `IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_filter_list": &schema.Schema{
							Description: `IP Addresses of the devices to be filtered out during discovery

ERROR: Different types for param ipFilterList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"is_auto_cdp": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"lldp_level": &schema.Schema{
							Description: `LLDP level to which neighbor devices to be discovered
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of the discovery
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"netconf_port": &schema.Schema{
							Description: `Netconf Port. It will need valid SSH credentials to work
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_list": &schema.Schema{
							Description: `Password of the devices to be discovered

ERROR: Different types for param passwordList schema.TypeList schema.TypeString`,
							Type:      schema.TypeList,
							Optional:  true,
							Sensitive: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preferred_mgmt_ipmethod": &schema.Schema{
							Description: `Preferred Management IP Method.'None' or 'UseLoopBack'. Default is 'None'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol_order": &schema.Schema{
							Description: `Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"retry": &schema.Schema{
							Description: `Number of times to try establishing connection to device
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Description: `Auth Pass phrase for SNMP
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Description: `SNMP auth protocol. SHA' or 'MD5'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Description: `Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Description: `Pass phrase for SNMP privacy
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Description: `SNMP privacy protocol. 'DES' or 'AES128'
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Description: `Snmp RO community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Description: `Description for Snmp RO community
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Description: `Snmp RW community of the devices to be discovered
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Description: `Description for Snmp RW community
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_user_name": &schema.Schema{
							Description: `SNMP username of the device
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_version": &schema.Schema{
							Description: `Version of SNMP. v2 or v3
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"timeout": &schema.Schema{
							Description: `Time to wait for device response in seconds
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"update_mgmt_ip": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"user_name_list": &schema.Schema{
							Description: `Username of the devices to be discovered

ERROR: Different types for param userNameList schema.TypeList schema.TypeString`,
							Type:     schema.TypeList,
							Optional: true,
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

func resourceDiscoveryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDiscoveryStartDiscovery(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okID && vvID != "" {
		getResponse1, _, err := client.Discovery.GetDiscoveryByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Discovery.StartDiscovery(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing StartDiscovery", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing StartDiscovery", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceDiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDiscoveryByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Discovery.GetDiscoveryByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	return diags
}

func resourceDiscoveryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]

	selectedMethod := 1
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.Discovery.GetDiscoveryByID(vvID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryByID", err,
				"Failure at GetDiscoveryByID, unexpected response", ""))
			return diags
		}
		//Set value vvName = getResp.
		if getResp.response != nil {
			vvName = getResp.response.Name
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdatesAnExistingDiscoveryBySpecifiedID(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAnExistingDiscoveryBySpecifiedID", err, restyResp1.String(),
					"Failure at UpdatesAnExistingDiscoveryBySpecifiedID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAnExistingDiscoveryBySpecifiedID", err,
				"Failure at UpdatesAnExistingDiscoveryBySpecifiedID, unexpected response", ""))
			return diags
		}
	}

	return resourceDiscoveryRead(ctx, d, m)
}

func resourceDiscoveryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestDiscoveryStartDiscovery(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscovery {
	request := dnacentersdkgo.RequestDiscoveryStartDiscovery{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cdp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cdp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cdp_level")))) {
		request.CdpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_type")))) {
		request.DiscoveryType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_list")))) {
		request.EnablePasswordList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_credential_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_credential_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_credential_id_list")))) {
		request.GlobalCredentialIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read_credential")))) {
		request.HTTPReadCredential = expandRequestDiscoveryStartDiscoveryHTTPReadCredential(ctx, key+".http_read_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write_credential")))) {
		request.HTTPWriteCredential = expandRequestDiscoveryStartDiscoveryHTTPWriteCredential(ctx, key+".http_write_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_list")))) {
		request.IPAddressList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_filter_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_filter_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_filter_list")))) {
		request.IPFilterList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lldp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lldp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lldp_level")))) {
		request.LldpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_list")))) {
		request.PasswordList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_mgmt_ipmethod")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) {
		request.PreferredMgmtIPMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol_order")))) {
		request.ProtocolOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retry")))) {
		request.Retry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) {
		request.SNMPROCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) {
		request.SNMPRWCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_list")))) {
		request.UserNameList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryStartDiscoveryHTTPReadCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPReadCredential {
	request := dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPReadCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryStartDiscoveryHTTPWriteCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPWriteCredential {
	request := dnacentersdkgo.RequestDiscoveryStartDiscoveryHTTPWriteCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_info")))) {
		request.AttributeInfo = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo(ctx, key+".attribute_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cdp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cdp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cdp_level")))) {
		request.CdpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ids")))) {
		request.DeviceIDs = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_condition")))) {
		request.DiscoveryCondition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_status")))) {
		request.DiscoveryStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_type")))) {
		request.DiscoveryType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password_list")))) {
		request.EnablePasswordList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_credential_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_credential_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_credential_id_list")))) {
		request.GlobalCredentialIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read_credential")))) {
		request.HTTPReadCredential = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential(ctx, key+".http_read_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write_credential")))) {
		request.HTTPWriteCredential = expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential(ctx, key+".http_write_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_list")))) {
		request.IPAddressList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_filter_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_filter_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_filter_list")))) {
		request.IPFilterList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_auto_cdp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_auto_cdp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_auto_cdp")))) {
		request.IsAutoCdp = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lldp_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lldp_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lldp_level")))) {
		request.LldpLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_devices")))) {
		request.NumDevices = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_discovery_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_discovery_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_discovery_id")))) {
		request.ParentDiscoveryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_list")))) {
		request.PasswordList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preferred_mgmt_ipmethod")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preferred_mgmt_ipmethod")))) {
		request.PreferredMgmtIPMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol_order")))) {
		request.ProtocolOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retry_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retry_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retry_count")))) {
		request.RetryCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPRoCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community_desc")))) {
		request.SNMPRoCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRwCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community_desc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community_desc")))) {
		request.SNMPRwCommunityDesc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_out")))) {
		request.TimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_mgmt_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_mgmt_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_mgmt_ip")))) {
		request.UpdateMgmtIP = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_list")))) {
		request.UserNameList = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo {
	var request dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential {
	request := dnacentersdkgo.RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
