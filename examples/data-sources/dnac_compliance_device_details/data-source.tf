
data "dnac_compliance_device_details" "example" {
    provider = dnac
    compliance_status = "string"
    compliance_type = "string"
    device_uuid = "string"
    limit = "string"
    offset = "string"
}

output "dnac_compliance_device_details_example" {
    value = data.dnac_compliance_device_details.example.items
}
