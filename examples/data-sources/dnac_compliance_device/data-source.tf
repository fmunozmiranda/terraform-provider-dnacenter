
data "dnac_compliance_device" "example" {
    provider = dnac
    compliance_status = "string"
    device_uuid = "string"
    limit = ------
    offset = ------
}

output "dnac_compliance_device_example" {
    value = data.dnac_compliance_device.example.items
}
