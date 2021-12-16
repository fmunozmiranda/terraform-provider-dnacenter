
data "dnacenter_sda_fabric_site" "example" {
    provider = dnac
    site_name_hierarchy = "string"
}

output "dnacenter_sda_fabric_site_example" {
    value = data.dnac_sda_fabric_site.example.item
}
