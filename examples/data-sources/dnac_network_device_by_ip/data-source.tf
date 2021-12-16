
data "dnacenter_network_device_by_ip" "example" {
    provider = dnac
    ip_address = "string"
}

output "dnacenter_network_device_by_ip_example" {
    value = data.dnac_network_device_by_ip.example.item
}
