
data "dnacenter_license_device_license_summary" "example" {
    provider = dnac
    device_type = "string"
    device_uuid = "string"
    dna_level = "string"
    limit = ------
    order = "string"
    page_number = ------
    registration_status = "string"
    smart_account_id = ------
    sort_by = "string"
    virtual_account_name = "string"
}

output "dnacenter_license_device_license_summary_example" {
    value = data.dnac_license_device_license_summary.example.items
}
