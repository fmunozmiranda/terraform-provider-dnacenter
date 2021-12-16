
data "dnac_network_device_count" "example" {
    provider = dnac
    device_id = "string"
}

output "dnac_network_device_count_example" {
    value = data.dnac_network_device_count.example.item_name
}

data "dnac_network_device_count" "example" {
    provider = dnac
    device_id = "string"
}

output "dnac_network_device_count_example" {
    value = data.dnac_network_device_count.example.item_id
}
