
data "dnac_site" "example" {
    provider = dnac
    limit = "string"
    name = "string"
    offset = "string"
    site_id = "string"
    type = "string"
}

output "dnac_site_example" {
    value = data.dnac_site.example.items
}
