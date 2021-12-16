
data "dnacenter_device_interface_count" "example" {
    provider = dnac
}

output "dnacenter_device_interface_count_example" {
    value = data.dnac_device_interface_count.example.item
}
