
data "dnacenter_device_family_identifiers_details" "example" {
    provider = dnac
}

output "dnacenter_device_family_identifiers_details_example" {
    value = data.dnac_device_family_identifiers_details.example.items
}
