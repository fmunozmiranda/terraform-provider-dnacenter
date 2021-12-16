
data "dnac_pnp_virtual_account_sync_result" "example" {
    provider = dnac
    domain = "string"
    name = "string"
}

output "dnac_pnp_virtual_account_sync_result_example" {
    value = data.dnac_pnp_virtual_account_sync_result.example.item
}
