
data "dnac_qos_device_interface_info_count" "example" {
    provider = dnac
}

output "dnac_qos_device_interface_info_count_example" {
    value = data.dnac_qos_device_interface_info_count.example.item
}
