
data "dnacenter_network_device_register_for_wsa" "example" {
    provider = dnac
    macaddress = "string"
    serial_number = "string"
}

output "dnacenter_network_device_register_for_wsa_example" {
    value = data.dnac_network_device_register_for_wsa.example.item
}
