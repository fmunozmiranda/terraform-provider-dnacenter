
data "dnac_network_device_inventory_insight_link_mismatch" "example" {
    provider = dnac
    category = "string"
    limit = "string"
    offset = "string"
    order = "string"
    site_id = "string"
    sort_by = "string"
}

output "dnac_network_device_inventory_insight_link_mismatch_example" {
    value = data.dnac_network_device_inventory_insight_link_mismatch.example.items
}
