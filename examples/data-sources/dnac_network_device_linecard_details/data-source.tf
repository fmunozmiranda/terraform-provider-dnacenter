
data "dnacenter_network_device_linecard_details" "example" {
    provider = dnac
    device_uuid = "string"
}

output "dnacenter_network_device_linecard_details_example" {
    value = data.dnac_network_device_linecard_details.example.items
}
