
data "dnacenter_device_interface_isis" "example" {
    provider = dnac
}

output "dnacenter_device_interface_isis_example" {
    value = data.dnac_device_interface_isis.example.items
}
