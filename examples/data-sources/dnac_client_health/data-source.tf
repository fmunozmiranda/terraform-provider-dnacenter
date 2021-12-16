
data "dnacenter_client_health" "example" {
    provider = dnac
    timestamp = "string"
}

output "dnacenter_client_health_example" {
    value = data.dnac_client_health.example.items
}
