package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterfaceNetworkDeviceRange() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the list of interfaces for the device for the specified range
`,

		ReadContext: dataSourceInterfaceNetworkDeviceRangeRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"records_to_return": &schema.Schema{
				Description: `recordsToReturn path parameter. Number of records to return
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_index": &schema.Schema{
				Description: `startIndex path parameter. Start index
`,
				Type:     schema.TypeInt,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"class_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"duplex": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"if_index": &schema.Schema{
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

						"interface_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_mask": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"isis_support": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"media_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"native_vlan_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ospf_support": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"pid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_no": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"speed": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"voice_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInterfaceNetworkDeviceRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")
	vStartIndex := d.Get("start_index")
	vRecordsToReturn := d.Get("records_to_return")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceInterfacesBySpecifiedRange")
		vvDeviceID := vDeviceID.(string)
		vvStartIndex := vStartIndex.(int)
		vvRecordsToReturn := vRecordsToReturn.(int)

		response1, _, err := client.Devices.GetDeviceInterfacesBySpecifiedRange(vvDeviceID, vvStartIndex, vvRecordsToReturn)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceInterfacesBySpecifiedRange", err,
				"Failure at GetDeviceInterfacesBySpecifiedRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDevicesGetDeviceInterfacesBySpecifiedRangeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceInterfacesBySpecifiedRange response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceInterfacesBySpecifiedRangeItems(items *[]dnacentersdkgo.ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["admin_status"] = item.AdminStatus
		respItem["class_name"] = item.ClassName
		respItem["description"] = item.Description
		respItem["device_id"] = item.DeviceID
		respItem["duplex"] = item.Duplex
		respItem["id"] = item.ID
		respItem["if_index"] = item.IfIndex
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["interface_type"] = item.InterfaceType
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv4_mask"] = item.IPv4Mask
		respItem["isis_support"] = item.IsisSupport
		respItem["last_updated"] = item.LastUpdated
		respItem["mac_address"] = item.MacAddress
		respItem["mapped_physical_interface_id"] = item.MappedPhysicalInterfaceID
		respItem["mapped_physical_interface_name"] = item.MappedPhysicalInterfaceName
		respItem["media_type"] = item.MediaType
		respItem["native_vlan_id"] = item.NativeVLANID
		respItem["ospf_support"] = item.OspfSupport
		respItem["pid"] = item.Pid
		respItem["port_mode"] = item.PortMode
		respItem["port_name"] = item.PortName
		respItem["port_type"] = item.PortType
		respItem["serial_no"] = item.SerialNo
		respItem["series"] = item.Series
		respItem["speed"] = item.Speed
		respItem["status"] = item.Status
		respItem["vlan_id"] = item.VLANID
		respItem["voice_vlan"] = item.VoiceVLAN
		respItems = append(respItems, respItem)
	}
	return respItems
}
