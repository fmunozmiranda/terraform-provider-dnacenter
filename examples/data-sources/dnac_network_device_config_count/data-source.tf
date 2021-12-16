
data "dnacenter_network_device_config_count" "example" {
    provider = dnac
}

output "dnacenter_network_device_config_count_example" {
    value = data.dnac_network_device_config_count.example.item
}
