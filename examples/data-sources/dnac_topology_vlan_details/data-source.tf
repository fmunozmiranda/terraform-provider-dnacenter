
data "dnacenter_topology_vlan_details" "example" {
    provider = dnac
}

output "dnacenter_topology_vlan_details_example" {
    value = data.dnac_topology_vlan_details.example.items
}
