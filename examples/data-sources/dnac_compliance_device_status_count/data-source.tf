
data "dnacenter_compliance_device_status_count" "example" {
    provider = dnac
    compliance_status = "string"
}

output "dnacenter_compliance_device_status_count_example" {
    value = data.dnac_compliance_device_status_count.example.item
}
