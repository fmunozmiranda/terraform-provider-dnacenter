
data "dnac_site_design_floormap" "example" {
    provider = dnac
    floor_id = "string"
}

output "dnac_site_design_floormap_example" {
    value = data.dnac_site_design_floormap.example.item
}