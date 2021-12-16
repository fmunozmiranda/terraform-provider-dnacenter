
data "dnac_issues_enrichment_details" "example" {
    provider = dnac
}

output "dnac_issues_enrichment_details_example" {
    value = data.dnac_issues_enrichment_details.example.item
}
