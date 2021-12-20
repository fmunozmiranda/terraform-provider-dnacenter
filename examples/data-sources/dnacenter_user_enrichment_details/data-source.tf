
data "dnacenter_user_enrichment_details" "example" {
  provider = dnacenter
}

output "dnacenter_user_enrichment_details_example" {
  value = data.dnacenter_user_enrichment_details.example.items
}
