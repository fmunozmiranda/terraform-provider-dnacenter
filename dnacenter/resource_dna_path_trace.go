package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePathTrace() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Path Trace.

- Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to
get results and follow progress.

- Deletes a flow analysis request by its id
`,

		CreateContext: resourcePathTraceCreate,
		ReadContext:   resourcePathTraceRead,
		UpdateContext: resourcePathTraceUpdate,
		DeleteContext: resourcePathTraceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPath': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'destIP': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'inclusions': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'periodicRefresh': {'Optional': 'true', 'Type': 'schema.TypeBool'}, 'protocol': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'sourceIP': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Optional': 'true', 'Type': 'schema.TypeString'}, 'flowAnalysisId': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'flowAnalysisId path parameter. Flow analysis request id\n'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'detailedStatus': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclTraceCalculation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'aclTraceCalculationFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'lastUpdate': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'networkElements': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'accuracyList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'percent': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'reason': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'detailedStatus': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclTraceCalculation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'aclTraceCalculationFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'deviceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'cpuStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'fiveMinUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'fiveSecsUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'oneMinUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'memoryStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'memoryUsage': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'totalMemory': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'deviceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'deviceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'egressPhysicalInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'egressVirtualInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'flexConnect': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'authentication': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataSwitching': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'egressAclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'ingressAclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'wirelessLanControllerId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'wirelessLanControllerName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ingressPhysicalInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'ingressVirtualInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'ip': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'linkInformationSource': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'byteRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'destIpAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputInterface': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipv4DSCP': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipv4TTL': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputInterface': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'packetBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetLoss': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetLossPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMax': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMean': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMin': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'sourceIpAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'role': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ssid': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'tunnels': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'wlanId': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'networkElementsInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'accuracyList': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'percent': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'reason': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'detailedStatus': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclTraceCalculation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'aclTraceCalculationFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'deviceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'cpuStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'fiveMinUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'fiveSecsUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'oneMinUsageInPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'memoryStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'memoryUsage': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'totalMemory': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}}}}, 'deviceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'deviceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'egressInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'physicalInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'virtualInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'flexConnect': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'authentication': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataSwitching': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'egressAclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'ingressAclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'wirelessLanControllerId': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'wirelessLanControllerName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ingressInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'physicalInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'virtualInterface': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'aclAnalysis': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'aclName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingAces': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ace': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'matchingPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'ports': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'destPorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'sourcePorts': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}}}}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'result': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'adminStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueFlushes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputQueueMaxDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'inputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'operationalStatus': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'outputDrop': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputQueueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputRatebps': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'interfaceStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'interfaceStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'pathOverlayInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'controlPlane': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dataPacketEncapsulation': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vxlanInfo': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'dscp': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vnid': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'qosStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'classMapName': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'dropRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'numPackets': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'offeredRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueBandwidthbps': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'queueDepth': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueNoBufferDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'queueTotalDrops': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}}}}, 'qosStatsCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'qosStatsCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'usedVlan': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'vrfName': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}, 'ip': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'linkInformationSource': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonCollection': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonCollectionFailureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'perfMonitorStatistics': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'byteRate': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'destIpAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inputInterface': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipv4DSCP': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ipv4TTL': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'outputInterface': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'packetBytes': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetCount': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetLoss': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'packetLossPercentage': {'Computed': 'true', 'Type': 'schema.TypeFloat'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'refreshedAt': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMax': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMean': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'rtpJitterMin': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'sourceIpAddress': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'role': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'ssid': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'tunnels': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'type': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'wlanId': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}, 'properties': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'request': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'controlPath': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'createTime': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'destIP': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'destPort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'failureReason': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'inclusions': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}}, 'lastUpdateTime': {'Computed': 'true', 'Type': 'schema.TypeInt'}, 'periodicRefresh': {'Computed': 'true', 'Type': 'schema.TypeBool'}, 'protocol': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourceIP': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'sourcePort': {'Computed': 'true', 'Type': 'schema.TypeString'}, 'status': {'Computed': 'true', 'Type': 'schema.TypeString'}}}}}}}}, 'metadata': {'item': {'operation_id': ['InitiateANewPathtrace', 'RetrievesPreviousPathtrace'], 'new_flat_structure': [{'RequestPathTraceInitiateANewPathtrace': {'type': 'obj', 'data': [{'name': 'controlPath', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'destIP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inclusions', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'periodicRefresh', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'ResponsePathTraceRetrievesPreviousPathtraceResponse': {'type': 'obj', 'data': [{'name': 'detailedStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseDetailedStatus'}, {'name': 'lastUpdate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'networkElements', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElements'}, {'name': 'networkElementsInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfo'}, {'name': 'properties', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'request', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseRequest'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseDetailedStatus': {'type': 'obj', 'data': [{'name': 'aclTraceCalculation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'aclTraceCalculationFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElements': {'type': 'obj', 'data': [{'name': 'accuracyList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsAccuracyList'}, {'name': 'detailedStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDetailedStatus'}, {'name': 'deviceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatistics'}, {'name': 'deviceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'egressPhysicalInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterface'}, {'name': 'egressVirtualInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterface'}, {'name': 'flexConnect', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnect'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ingressPhysicalInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterface'}, {'name': 'ingressVirtualInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterface'}, {'name': 'ip', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'linkInformationSource', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsPerfMonStatistics'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ssid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tunnels', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'wlanId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsAccuracyList': {'type': 'obj', 'data': [{'name': 'percent', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'reason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDetailedStatus': {'type': 'obj', 'data': [{'name': 'aclTraceCalculation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'aclTraceCalculationFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatistics': {'type': 'obj', 'data': [{'name': 'cpuStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCpuStatistics'}, {'name': 'memoryStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsMemoryStatistics'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCpuStatistics': {'type': 'obj', 'data': [{'name': 'fiveMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'fiveSecsUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'oneMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsMemoryStatistics': {'type': 'obj', 'data': [{'name': 'memoryUsage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'totalMemory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnect': {'type': 'obj', 'data': [{'name': 'authentication', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataSwitching', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'egressAclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysis'}, {'name': 'ingressAclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysis'}, {'name': 'wirelessLanControllerId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'wirelessLanControllerName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsPerfMonStatistics': {'type': 'obj', 'data': [{'name': 'byteRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'destIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4DSCP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4TTL', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'packetBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetLoss', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetLossPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMax', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMean', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMin', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'sourceIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfo': {'type': 'obj', 'data': [{'name': 'accuracyList', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoAccuracyList'}, {'name': 'detailedStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDetailedStatus'}, {'name': 'deviceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatistics'}, {'name': 'deviceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'egressInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterface'}, {'name': 'flexConnect', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnect'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ingressInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterface'}, {'name': 'ip', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'linkInformationSource', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'perfMonitorStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoPerfMonitorStatistics'}, {'name': 'role', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ssid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'tunnels', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'type', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'wlanId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoAccuracyList': {'type': 'obj', 'data': [{'name': 'percent', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'reason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDetailedStatus': {'type': 'obj', 'data': [{'name': 'aclTraceCalculation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'aclTraceCalculationFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatistics': {'type': 'obj', 'data': [{'name': 'cpuStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCpuStatistics'}, {'name': 'memoryStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCpuStatistics': {'type': 'obj', 'data': [{'name': 'fiveMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'fiveSecsUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'oneMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics': {'type': 'obj', 'data': [{'name': 'memoryUsage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'totalMemory', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterface': {'type': 'obj', 'data': [{'name': 'physicalInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterface'}, {'name': 'virtualInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterface'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnect': {'type': 'obj', 'data': [{'name': 'authentication', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataSwitching', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'egressAclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysis'}, {'name': 'ingressAclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysis'}, {'name': 'wirelessLanControllerId', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'wirelessLanControllerName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterface': {'type': 'obj', 'data': [{'name': 'physicalInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterface'}, {'name': 'virtualInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterface'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterface': {'type': 'obj', 'data': [{'name': 'aclAnalysis', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysis'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics'}, {'name': 'interfaceStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'interfaceStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'pathOverlayInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo'}, {'name': 'qosStatistics', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics'}, {'name': 'qosStatsCollection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'qosStatsCollectionFailureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'usedVlan', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vrfName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics': {'type': 'obj', 'data': [{'name': 'adminStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueFlushes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputQueueMaxDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'inputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'operationalStatus', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'outputDrop', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputQueueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputRatebps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo': {'type': 'obj', 'data': [{'name': 'controlPlane', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dataPacketEncapsulation', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vxlanInfo', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo': {'type': 'obj', 'data': [{'name': 'dscp', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'vnid', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics': {'type': 'obj', 'data': [{'name': 'classMapName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'dropRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'numPackets', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'offeredRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueBandwidthbps', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'queueDepth', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueNoBufferDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'queueTotalDrops', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoPerfMonitorStatistics': {'type': 'obj', 'data': [{'name': 'byteRate', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'destIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inputInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4DSCP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'ipv4TTL', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'outputInterface', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'packetBytes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetCount', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetLoss', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'packetLossPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMax', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMean', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'rtpJitterMin', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'sourceIpAddress', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseRequest': {'type': 'obj', 'data': [{'name': 'controlPath', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'createTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'destIP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'destPort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'failureReason', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'inclusions', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'lastUpdateTime', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'periodicRefresh', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourceIP', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'sourcePort', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'status', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCPUStatistics': {'type': 'obj', 'data': [{'name': 'fiveMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'fiveSecsUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'oneMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCPUStatistics': {'type': 'obj', 'data': [{'name': 'fiveMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'fiveSecsUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'oneMinUsageInPercentage', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'float64'}, {'name': 'refreshedAt', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis': {'type': 'obj', 'data': [{'name': 'aclName', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingAces', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAces'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces': {'type': 'obj', 'data': [{'name': 'ace', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'matchingPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPorts'}, {'name': 'result', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts': {'type': 'obj', 'data': [{'name': 'ports', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceAclAnalysisMatchingAcesMatchingPortsPorts'}, {'name': 'protocol', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts': {'type': 'obj', 'data': [{'name': 'destPorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'sourcePorts', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': ['RequestPathTraceInitiateANewPathtrace', 'ResponsePathTraceRetrievesPreviousPathtraceResponse'], 'access_list': [[], ['response']]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detailed_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"acl_trace_calculation": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"acl_trace_calculation_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"last_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_elements": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"egress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
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
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"egress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
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
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"wireless_lan_controller_name": &schema.Schema{
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
									"ingress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
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
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ingress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
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
												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"network_elements_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"egress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
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
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
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
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"wireless_lan_controller_name": &schema.Schema{
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
									"ingress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
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
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},
																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"result": &schema.Schema{
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
															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"perf_monitor_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"request": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_path": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"create_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"dest_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dest_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"inclusions": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"last_update_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"periodic_refresh": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

						"control_path": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"dest_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dest_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"flow_analysis_id": &schema.Schema{
							Description: `flowAnalysisId path parameter. Flow analysis request id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"inclusions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"periodic_refresh": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePathTraceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestPathTraceInitiateANewPathtrace(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vFlowAnalysisID, okFlowAnalysisID := resourceItem["flow_analysis_id"]
	vvFlowAnalysisID := interfaceToString(vFlowAnalysisID)
	if okFlowAnalysisID && vvFlowAnalysisID != "" {
		getResponse2, _, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["flow_analysis_id"] = vvFlowAnalysisID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.PathTrace.RetrivesAllPreviousPathtracesSummary(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsPathTraceRetrivesAllPreviousPathtracesSummary(m, response2, nil)
			item2, err := searchPathTraceRetrivesAllPreviousPathtracesSummary(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["flow_analysis_id"] = vvFlowAnalysisID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.PathTrace.InitiateANewPathtrace(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing InitiateANewPathtrace", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing InitiateANewPathtrace", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["flow_analysis_id"] = vvFlowAnalysisID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourcePathTraceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPeriodicRefresh, okPeriodicRefresh := resourceMap["periodic_refresh"]
	vSourceIP, okSourceIP := resourceMap["source_ip"]
	vDestIP, okDestIP := resourceMap["dest_ip"]
	vSourcePort, okSourcePort := resourceMap["source_port"]
	vDestPort, okDestPort := resourceMap["dest_port"]
	vGtCreateTime, okGtCreateTime := resourceMap["gt_create_time"]
	vLtCreateTime, okLtCreateTime := resourceMap["lt_create_time"]
	vProtocol, okProtocol := resourceMap["protocol"]
	vStatus, okStatus := resourceMap["status"]
	vTaskID, okTaskID := resourceMap["task_id"]
	vLastUpdateTime, okLastUpdateTime := resourceMap["last_update_time"]
	vLimit, okLimit := resourceMap["limit"]
	vOffset, okOffset := resourceMap["offset"]
	vOrder, okOrder := resourceMap["order"]
	vSortBy, okSortBy := resourceMap["sort_by"]
	vFlowAnalysisID, okFlowAnalysisID := resourceMap["flow_analysis_id"]

	method1 := []bool{okPeriodicRefresh, okSourceIP, okDestIP, okSourcePort, okDestPort, okGtCreateTime, okLtCreateTime, okProtocol, okStatus, okTaskID, okLastUpdateTime, okLimit, okOffset, okOrder, okSortBy}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okFlowAnalysisID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RetrivesAllPreviousPathtracesSummary")
		queryParams1 := dnacentersdkgo.RetrivesAllPreviousPathtracesSummaryQueryParams{}

		if okPeriodicRefresh {
			queryParams1.PeriodicRefresh = vPeriodicRefresh.(bool)
		}
		if okSourceIP {
			queryParams1.SourceIP = vSourceIP.(string)
		}
		if okDestIP {
			queryParams1.DestIP = vDestIP.(string)
		}
		if okSourcePort {
			queryParams1.SourcePort = vSourcePort.(string)
		}
		if okDestPort {
			queryParams1.DestPort = vDestPort.(string)
		}
		if okGtCreateTime {
			queryParams1.GtCreateTime = vGtCreateTime.(string)
		}
		if okLtCreateTime {
			queryParams1.LtCreateTime = vLtCreateTime.(string)
		}
		if okProtocol {
			queryParams1.Protocol = vProtocol.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
		if okLastUpdateTime {
			queryParams1.LastUpdateTime = vLastUpdateTime.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}

		response1, restyResp1, err := client.PathTrace.RetrivesAllPreviousPathtracesSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetrivesAllPreviousPathtracesSummary", err,
				"Failure at RetrivesAllPreviousPathtracesSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: RetrievesPreviousPathtrace")
		vvFlowAnalysisID := vFlowAnalysisID.(string)

		response2, restyResp2, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RetrievesPreviousPathtrace", err,
				"Failure at RetrievesPreviousPathtrace, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourcePathTraceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePathTraceRead(ctx, d, m)
}

func resourcePathTraceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestPathTraceInitiateANewPathtrace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPathTraceInitiateANewPathtrace {
	request := dnacentersdkgo.RequestPathTraceInitiateANewPathtrace{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".control_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".control_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".control_path")))) {
		request.ControlPath = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_ip")))) {
		request.DestIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_port")))) {
		request.DestPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".inclusions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".inclusions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".inclusions")))) {
		request.Inclusions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".periodic_refresh")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".periodic_refresh")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".periodic_refresh")))) {
		request.PeriodicRefresh = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_ip")))) {
		request.SourceIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_port")))) {
		request.SourcePort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
