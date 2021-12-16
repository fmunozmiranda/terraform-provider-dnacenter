
data "dnac_topology_physical" "example" {
    provider = dnac
    node_type = "string"
}

output "dnac_topology_physical_example" {
    value = data.dnac_topology_physical.example.item
}
