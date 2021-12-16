
data "dnac_topology_layer_2" "example" {
    provider = dnac
    vlan_id = "string"
}

output "dnac_topology_layer_2_example" {
    value = data.dnac_topology_layer_2.example.item
}
