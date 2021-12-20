
data "dnacenter_client_enrichment_details" "example" {
  provider = dnacenter
}

output "dnacenter_client_enrichment_details_example" {
  value = data.dnacenter_client_enrichment_details.example.items
}
