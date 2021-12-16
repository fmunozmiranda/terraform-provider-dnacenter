
data "dnac_system_health_count" "example" {
    provider = dnac
    domain = "string"
    subdomain = "string"
}

output "dnac_system_health_count_example" {
    value = data.dnac_system_health_count.example.item
}
