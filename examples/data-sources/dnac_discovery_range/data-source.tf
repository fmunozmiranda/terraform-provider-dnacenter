
data "dnac_discovery_range" "example" {
    provider = dnac
    records_to_return = 1
    start_index = 1
}

output "dnac_discovery_range_example" {
    value = data.dnac_discovery_range.example.items
}