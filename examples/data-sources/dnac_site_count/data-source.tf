
data "dnacenter_site_count" "example" {
    provider = dnac
    site_id = "string"
}

output "dnacenter_site_count_example" {
    value = data.dnac_site_count.example.item
}
