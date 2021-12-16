
data "dnac_network_device_vlan" "example" {
    provider = dnac
    id = "string"
    interface_type = "string"
}

output "dnac_network_device_vlan_example" {
    value = data.dnac_network_device_vlan.example.items
}
