
data "dnac_discovery_device" "example" {
    provider = dnac
    id = "string"
    task_id = "string"
}

output "dnac_discovery_device_example" {
    value = data.dnac_discovery_device.example.items
}
