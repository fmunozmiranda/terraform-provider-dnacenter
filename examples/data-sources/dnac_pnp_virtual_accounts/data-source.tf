
data "dnacenter_pnp_virtual_accounts" "example" {
    provider = dnac
    domain = "string"
}

output "dnacenter_pnp_virtual_accounts_example" {
    value = data.dnac_pnp_virtual_accounts.example.items
}
