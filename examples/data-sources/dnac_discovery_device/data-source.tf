
data "dnacenter_discovery_device" "example" {
    provider = dnac
    id = "string"
    task_id = "string"
}

output "dnacenter_discovery_device_example" {
    value = data.dnac_discovery_device.example.items
}
