
data "dnac_interface_network_device_detail" "example" {
    provider = dnac
    device_id = "string"
    name = "string"
}

output "dnac_interface_network_device_detail_example" {
    value = data.dnac_interface_network_device_detail.example.item
}