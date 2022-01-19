
resource "dnac_sda_fabric_edge_device" "example" {
    provider = dnac
    parameters {
      
      device_management_ip_address = "string"
      site_name_hierarchy = "string"
    }
}

output "dnac_sda_fabric_edge_device_example" {
    value = dnac_sda_fabric_edge_device.example
}