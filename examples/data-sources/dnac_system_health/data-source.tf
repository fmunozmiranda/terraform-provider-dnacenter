
data "dnac_system_health" "example" {
    provider = dnac
    domain = "string"
    limit = ------
    offset = ------
    subdomain = "string"
    summary = "false"
}

output "dnac_system_health_example" {
    value = data.dnac_system_health.example.item
}
