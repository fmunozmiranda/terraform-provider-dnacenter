
data "dnacenter_wireless_profile" "example" {
    provider = dnac
    profile_name = "string"
}

output "dnacenter_wireless_profile_example" {
    value = data.dnac_wireless_profile.example.items
}
