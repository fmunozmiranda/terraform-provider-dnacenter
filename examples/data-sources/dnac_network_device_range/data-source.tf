
data "dnacenter_network_device_range" "example" {
    provider = dnac
    records_to_return = 1
    start_index = 1
}

output "dnacenter_network_device_range_example" {
    value = data.dnac_network_device_range.example.items
}
