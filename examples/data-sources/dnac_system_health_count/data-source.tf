
data "dnacenter_system_health_count" "example" {
    provider = dnac
    domain = "string"
    subdomain = "string"
}

output "dnacenter_system_health_count_example" {
    value = data.dnac_system_health_count.example.item
}
