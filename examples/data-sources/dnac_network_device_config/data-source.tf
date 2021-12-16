
data "dnac_network_device_config" "example" {
    provider = dnac
}

output "dnac_network_device_config_example" {
    value = data.dnac_network_device_config.example.items
}

data "dnac_network_device_config" "example" {
    provider = dnac
    network_device_id = "string"
}

output "dnac_network_device_config_example" {
    value = data.dnac_network_device_config.example.item
}
