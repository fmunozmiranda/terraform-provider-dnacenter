
resource "dnac_sda_virtual_network" "example" {
    provider = dnac
    parameters {
      
      site_name_hierarchy = "string"
      virtual_network_name = "string"
    }
}

output "dnac_sda_virtual_network_example" {
    value = dnac_sda_virtual_network.example
}