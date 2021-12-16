
data "dnac_topology_vlan_details" "example" {
    provider = dnac
}

output "dnac_topology_vlan_details_example" {
    value = data.dnac_topology_vlan_details.example.items
}
