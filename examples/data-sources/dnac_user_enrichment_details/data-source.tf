
data "dnac_user_enrichment_details" "example" {
    provider = dnac
}

output "dnac_user_enrichment_details_example" {
    value = data.dnac_user_enrichment_details.example.items
}
