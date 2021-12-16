
data "dnac_global_credential" "example" {
    provider = dnac
    credential_sub_type = "string"
    order = "string"
    sort_by = "string"
}

output "dnac_global_credential_example" {
    value = data.dnac_global_credential.example.items
}

data "dnac_global_credential" "example" {
    provider = dnac
    id = "string"
}

output "dnac_global_credential_example" {
    value = data.dnac_global_credential.example.item
}
