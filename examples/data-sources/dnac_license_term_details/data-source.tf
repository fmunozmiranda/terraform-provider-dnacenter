
data "dnacenter_license_term_details" "example" {
    provider = dnac
    device_type = "string"
    smart_account_id = "string"
    virtual_account_name = "string"
}

output "dnacenter_license_term_details_example" {
    value = data.dnac_license_term_details.example.items
}
