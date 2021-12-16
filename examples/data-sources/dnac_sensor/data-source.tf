
data "dnac_sensor" "example" {
    provider = dnac
    site_id = "string"
}

output "dnac_sensor_example" {
    value = data.dnac_sensor.example.items
}
