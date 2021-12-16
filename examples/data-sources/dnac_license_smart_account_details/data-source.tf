
data "dnacenter_license_smart_account_details" "example" {
    provider = dnac
}

output "dnacenter_license_smart_account_details_example" {
    value = data.dnac_license_smart_account_details.example.items
}
