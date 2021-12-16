
data "dnac_security_advisories_devices" "example" {
    provider = dnac
    advisory_id = "string"
}

output "dnac_security_advisories_devices_example" {
    value = data.dnac_security_advisories_devices.example.items
}
