
data "dnac_sda_device_role" "example" {
    provider = dnac
    device_management_ip_address = "string"
}

output "dnac_sda_device_role_example" {
    value = data.dnac_sda_device_role.example.item
}
