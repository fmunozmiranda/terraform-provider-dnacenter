
resource "dnacenter_network_device" "example" {
    provider = dnacenter
    item {
      
      # ap_manager_interface_ip = ------
      # associated_wlc_ip = ------
      # boot_date_time = ------
      # collection_interval = ------
      # collection_status = ------
      # error_code = ------
      # error_description = ------
      # family = ------
      # hostname = ------
      # id = ------
      # instance_tenant_id = ------
      # instance_uuid = ------
      # interface_count = ------
      # inventory_status_detail = ------
      # last_update_time = ------
      # last_updated = ------
      # line_card_count = ------
      # line_card_id = ------
      # location = ------
      # location_name = ------
      # mac_address = ------
      # management_ip_address = ------
      # memory_size = ------
      # platform_id = ------
      # reachability_failure_reason = ------
      # reachability_status = ------
      # role = ------
      # role_source = ------
      # serial_number = ------
      # series = ------
      # snmp_contact = ------
      # snmp_location = ------
      # software_type = ------
      # software_version = ------
      # tag_count = ------
      # tunnel_udp_port = ------
      # type = ------
      # up_time = ------
      # waas_device_mode = ------
    }
    parameters {
      
      id = "string"
    }
}

output "dnacenter_network_device_example" {
    value = dnacenter_network_device.example
}