
data "dnacenter_applications_health" "example" {
    provider = dnac
    application_health = "string"
    application_name = "string"
    device_id = "string"
    end_time = ------
    limit = ------
    mac_address = "string"
    offset = ------
    site_id = "string"
    start_time = ------
}

output "dnacenter_applications_health_example" {
    value = data.dnac_applications_health.example.items
}
