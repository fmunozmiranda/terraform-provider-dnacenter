
data "dnac_discovery_device_range" "example" {
    provider = dnac
    id = "string"
    records_to_return = 1
    start_index = 1
    task_id = "string"
}

output "dnac_discovery_device_range_example" {
    value = data.dnac_discovery_device_range.example.items
}
