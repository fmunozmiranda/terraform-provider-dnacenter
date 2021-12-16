
data "dnacenter_device_interface" "example" {
    provider = dnac
    limit = ------
    offset = ------
}

output "dnacenter_device_interface_example" {
    value = data.dnac_device_interface.example.items
}

data "dnacenter_device_interface" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_device_interface_example" {
    value = data.dnac_device_interface.example.item
}
