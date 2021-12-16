
data "dnacenter_network_device_wireless_lan" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_network_device_wireless_lan_example" {
    value = data.dnac_network_device_wireless_lan.example.item
}
