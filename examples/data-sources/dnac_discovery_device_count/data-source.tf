
data "dnacenter_discovery_device_count" "example" {
    provider = dnac
    id = "string"
    task_id = "string"
}

output "dnacenter_discovery_device_count_example" {
    value = data.dnac_discovery_device_count.example.item
}
