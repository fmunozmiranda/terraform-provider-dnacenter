
data "dnacenter_nfv_profile" "example" {
    provider = dnac
    id = "string"
    limit = "string"
    name = "string"
    offset = "string"
}

output "dnacenter_nfv_profile_example" {
    value = data.dnac_nfv_profile.example.items
}
