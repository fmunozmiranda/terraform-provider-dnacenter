
data "dnac_wireless_rf_profile" "example" {
    provider = dnac
    rf_profile_name = "string"
}

output "dnac_wireless_rf_profile_example" {
    value = data.dnac_wireless_rf_profile.example.items
}