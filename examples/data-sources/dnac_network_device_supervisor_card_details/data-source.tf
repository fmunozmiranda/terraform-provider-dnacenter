
data "dnac_network_device_supervisor_card_details" "example" {
    provider = dnac
    device_uuid = "string"
}

output "dnac_network_device_supervisor_card_details_example" {
    value = data.dnac_network_device_supervisor_card_details.example.items
}
