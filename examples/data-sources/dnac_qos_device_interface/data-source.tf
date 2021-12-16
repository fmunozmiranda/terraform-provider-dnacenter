
data "dnac_qos_device_interface" "example" {
    provider = dnac
    network_device_id = "string"
}

output "dnac_qos_device_interface_example" {
    value = data.dnac_qos_device_interface.example.items
}
