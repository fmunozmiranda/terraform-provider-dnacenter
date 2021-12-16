
data "dnacenter_network_device_stack_details" "example" {
    provider = dnac
    device_id = "string"
}

output "dnacenter_network_device_stack_details_example" {
    value = data.dnac_network_device_stack_details.example.item
}
