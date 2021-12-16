
data "dnac_service_provider" "example" {
    provider = dnac
}

output "dnac_service_provider_example" {
    value = data.dnac_service_provider.example.items
}
