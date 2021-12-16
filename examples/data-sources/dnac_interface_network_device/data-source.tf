
data "dnac_interface_network_device" "example" {
    provider = dnac
    device_id = "string"
}

output "dnac_interface_network_device_example" {
    value = data.dnac_interface_network_device.example.items
}