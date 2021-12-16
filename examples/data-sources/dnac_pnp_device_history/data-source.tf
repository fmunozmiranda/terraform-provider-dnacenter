
data "dnacenter_pnp_device_history" "example" {
    provider = dnac
    serial_number = "string"
    sort = ["string"]
    sort_order = "string"
}

output "dnacenter_pnp_device_history_example" {
    value = data.dnac_pnp_device_history.example.items
}
