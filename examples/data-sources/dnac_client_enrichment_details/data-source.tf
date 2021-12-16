
data "dnac_client_enrichment_details" "example" {
    provider = dnac
}

output "dnac_client_enrichment_details_example" {
    value = data.dnac_client_enrichment_details.example.items
}
