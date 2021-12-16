
data "dnac_device_interface_ospf" "example" {
    provider = dnac
}

output "dnac_device_interface_ospf_example" {
    value = data.dnac_device_interface_ospf.example.items
}
