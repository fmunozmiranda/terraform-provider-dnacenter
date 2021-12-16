
data "dnac_device_interface_by_ip" "example" {
    provider = dnac
    ip_address = "string"
}

output "dnac_device_interface_by_ip_example" {
    value = data.dnac_device_interface_by_ip.example.items
}
