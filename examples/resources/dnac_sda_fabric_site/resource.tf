
resource "dnac_sda_fabric_site" "example" {
    provider = dnac
    parameters {
      
      fabric_name = "string"
      site_name_hierarchy = "string"
    }
}

output "dnac_sda_fabric_site_example" {
    value = dnac_sda_fabric_site.example
}