
data "dnacenter_site" "example" {
    provider = dnac
    limit = "string"
    name = "string"
    offset = "string"
    site_id = "string"
    type = "string"
}

output "dnacenter_site_example" {
    value = data.dnac_site.example.items
}
