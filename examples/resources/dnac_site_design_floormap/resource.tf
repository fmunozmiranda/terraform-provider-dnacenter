
resource "dnacenter_site_design_floormap" "example" {
    provider = dnacenter
    # item = ------
    parameters {
      
      floor_id = "string"
    }
}

output "dnacenter_site_design_floormap_example" {
    value = dnacenter_site_design_floormap.example
}