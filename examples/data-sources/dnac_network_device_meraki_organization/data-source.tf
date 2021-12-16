
data "dnac_network_device_meraki_organization" "example" {
    provider = dnac
    id = "string"
}

output "dnac_network_device_meraki_organization_example" {
    value = data.dnac_network_device_meraki_organization.example.items
}
