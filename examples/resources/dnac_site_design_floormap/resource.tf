
resource "dnac_site_design_floormap" "example" {
    provider = dnac
    # item = ------
    parameters {
      
      floor_id = "string"
    }
}

output "dnac_site_design_floormap_example" {
    value = dnac_site_design_floormap.example
}