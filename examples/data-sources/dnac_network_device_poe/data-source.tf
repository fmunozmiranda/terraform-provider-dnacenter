
data "dnacenter_network_device_poe" "example" {
    provider = dnac
    device_uuid = "string"
}

output "dnacenter_network_device_poe_example" {
    value = data.dnac_network_device_poe.example.item
}
