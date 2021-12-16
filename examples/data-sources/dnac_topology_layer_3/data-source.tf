
data "dnacenter_topology_layer_3" "example" {
    provider = dnac
    topology_type = "string"
}

output "dnacenter_topology_layer_3_example" {
    value = data.dnac_topology_layer_3.example.item
}
