
data "dnacenter_network_device_by_serial_number" "example" {
    provider = dnac
    serial_number = "string"
}

output "dnacenter_network_device_by_serial_number_example" {
    value = data.dnac_network_device_by_serial_number.example.item
}
