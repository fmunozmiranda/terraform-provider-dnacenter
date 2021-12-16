
data "dnac_discovery_count" "example" {
    provider = dnac
}

output "dnac_discovery_count_example" {
    value = data.dnac_discovery_count.example.item
}
