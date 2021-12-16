
data "dnacenter_license_virtual_account_details" "example" {
    provider = dnac
    smart_account_id = "string"
}

output "dnacenter_license_virtual_account_details_example" {
    value = data.dnac_license_virtual_account_details.example.item
}
