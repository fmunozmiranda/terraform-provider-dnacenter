
data "dnac_pnp_smart_account_domains" "example" {
    provider = dnac
}

output "dnac_pnp_smart_account_domains_example" {
    value = data.dnac_pnp_smart_account_domains.example.items
}
