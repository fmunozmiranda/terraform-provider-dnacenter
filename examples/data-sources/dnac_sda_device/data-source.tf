
data "dnac_sda_device" "example" {
    provider = dnac
    device_management_ip_address = "string"
}

output "dnac_sda_device_example" {
    value = data.dnac_sda_device.example.item
}