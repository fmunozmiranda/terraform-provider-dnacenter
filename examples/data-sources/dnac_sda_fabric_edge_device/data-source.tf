
data "dnac_sda_fabric_edge_device" "example" {
    provider = dnac
    device_management_ip_address = "string"
}

output "dnac_sda_fabric_edge_device_example" {
    value = data.dnac_sda_fabric_edge_device.example.item
}
