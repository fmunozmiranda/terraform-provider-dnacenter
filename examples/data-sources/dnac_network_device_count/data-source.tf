
data "dnacenter_network_device_count" "example" {
    provider = dnac
    device_id = "string"
}

output "dnacenter_network_device_count_example" {
    value = data.dnac_network_device_count.example.item_name
}

data "dnacenter_network_device_count" "example" {
    provider = dnac
    device_id = "string"
}

output "dnacenter_network_device_count_example" {
    value = data.dnac_network_device_count.example.item_id
}
