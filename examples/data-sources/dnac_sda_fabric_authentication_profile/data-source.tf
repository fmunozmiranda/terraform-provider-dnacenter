
data "dnacenter_sda_fabric_authentication_profile" "example" {
    provider = dnac
    authenticate_template_name = "string"
    site_name_hierarchy = "string"
}

output "dnacenter_sda_fabric_authentication_profile_example" {
    value = data.dnac_sda_fabric_authentication_profile.example.item
}
