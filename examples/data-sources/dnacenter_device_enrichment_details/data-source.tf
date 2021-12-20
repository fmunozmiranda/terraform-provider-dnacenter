
data "dnacenter_device_enrichment_details" "example" {
  provider = dnacenter
}

output "dnacenter_device_enrichment_details_example" {
  value = data.dnacenter_device_enrichment_details.example.items
}
