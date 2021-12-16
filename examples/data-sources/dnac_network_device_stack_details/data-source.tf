
data "dnac_network_device_stack_details" "example" {
    provider = dnac
    device_id = "string"
}

output "dnac_network_device_stack_details_example" {
    value = data.dnac_network_device_stack_details.example.item
}
