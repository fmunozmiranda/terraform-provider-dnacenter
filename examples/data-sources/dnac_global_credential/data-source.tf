
data "dnacenter_global_credential" "example" {
    provider = dnac
    credential_sub_type = "string"
    order = "string"
    sort_by = "string"
}

output "dnacenter_global_credential_example" {
    value = data.dnac_global_credential.example.items
}

data "dnacenter_global_credential" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_global_credential_example" {
    value = data.dnac_global_credential.example.item
}
