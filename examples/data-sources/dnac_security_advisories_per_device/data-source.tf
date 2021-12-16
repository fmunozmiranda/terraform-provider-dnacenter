
data "dnacenter_security_advisories_per_device" "example" {
    provider = dnac
    device_id = "string"
}

output "dnacenter_security_advisories_per_device_example" {
    value = data.dnac_security_advisories_per_device.example.items
}
