
data "dnac_network_device_interface_poe" "example" {
    provider = dnac
    device_uuid = "string"
    interface_name_list = "string"
}

output "dnac_network_device_interface_poe_example" {
    value = data.dnac_network_device_interface_poe.example.items
}
