
data "dnacenter_network_device_config" "example" {
    provider = dnac
}

output "dnacenter_network_device_config_example" {
    value = data.dnac_network_device_config.example.items
}

data "dnacenter_network_device_config" "example" {
    provider = dnac
    network_device_id = "string"
}

output "dnacenter_network_device_config_example" {
    value = data.dnac_network_device_config.example.item
}
