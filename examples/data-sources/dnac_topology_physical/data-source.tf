
data "dnacenter_topology_physical" "example" {
    provider = dnac
    node_type = "string"
}

output "dnacenter_topology_physical_example" {
    value = data.dnac_topology_physical.example.item
}
