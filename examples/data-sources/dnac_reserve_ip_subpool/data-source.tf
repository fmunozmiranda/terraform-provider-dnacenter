
data "dnac_reserve_ip_subpool" "example" {
    provider = dnac
    limit = "string"
    offset = "string"
    site_id = "string"
}

output "dnac_reserve_ip_subpool_example" {
    value = data.dnac_reserve_ip_subpool.example.items
}
