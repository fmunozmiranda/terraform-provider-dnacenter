
data "dnac_compliance_device_by_id" "example" {
    provider = dnac
    device_uuid = "string"
}

output "dnac_compliance_device_by_id_example" {
    value = data.dnac_compliance_device_by_id.example.item
}
