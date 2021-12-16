
data "dnacenter_device_enrichment_details" "example" {
    provider = dnac
}

output "dnacenter_device_enrichment_details_example" {
    value = data.dnac_device_enrichment_details.example.items
}
