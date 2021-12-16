
data "dnacenter_topology_network_health" "example" {
    provider = dnac
    timestamp = "string"
}

output "dnacenter_topology_network_health_example" {
    value = data.dnac_topology_network_health.example.items
}
