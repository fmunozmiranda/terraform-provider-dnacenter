
data "dnac_pnp_global_settings" "example" {
    provider = dnac
}

output "dnac_pnp_global_settings_example" {
    value = data.dnac_pnp_global_settings.example.item
}
