
data "dnac_device_interface" "example" {
    provider = dnac
    limit = ------
    offset = ------
}

output "dnac_device_interface_example" {
    value = data.dnac_device_interface.example.items
}

data "dnac_device_interface" "example" {
    provider = dnac
    id = "string"
}

output "dnac_device_interface_example" {
    value = data.dnac_device_interface.example.item
}
