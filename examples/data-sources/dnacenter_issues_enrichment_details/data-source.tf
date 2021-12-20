
data "dnacenter_issues_enrichment_details" "example" {
  provider = dnacenter
}

output "dnacenter_issues_enrichment_details_example" {
  value = data.dnacenter_issues_enrichment_details.example.item
}
