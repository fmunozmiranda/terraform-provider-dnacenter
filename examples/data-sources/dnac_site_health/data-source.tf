
data "dnac_site_health" "example" {
    provider = dnac
    limit = ------
    offset = ------
    site_type = "string"
    timestamp = "string"
}

output "dnac_site_health_example" {
    value = data.dnac_site_health.example.items
}