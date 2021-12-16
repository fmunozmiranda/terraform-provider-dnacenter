
data "dnac_license_device_license_details" "example" {
    provider = dnac
    device_uuid = "string"
}

output "dnac_license_device_license_details_example" {
    value = data.dnac_license_device_license_details.example.items
}
