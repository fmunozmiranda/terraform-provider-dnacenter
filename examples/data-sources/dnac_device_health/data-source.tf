
data "dnac_device_health" "example" {
    provider = dnac
    device_role = "string"
    end_time = ------
    health = "string"
    limit = ------
    offset = ------
    site_id = "string"
    start_time = ------
}

output "dnac_device_health_example" {
    value = data.dnac_device_health.example.items
}
