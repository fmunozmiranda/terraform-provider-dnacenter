
data "dnacenter_sensor" "example" {
    provider = dnac
    site_id = "string"
}

output "dnacenter_sensor_example" {
    value = data.dnac_sensor.example.items
}
