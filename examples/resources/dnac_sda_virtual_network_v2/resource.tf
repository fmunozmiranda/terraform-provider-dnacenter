
resource "dnac_sda_virtual_network_v2" "example" {
    provider = dnac
    parameters {
      
      is_guest_virtual_network = "false"
      scalable_group_names = ["string"]
      virtual_network_name = "string"
      virtual_network_type = "string"
    }
}

output "dnac_sda_virtual_network_v2_example" {
    value = dnac_sda_virtual_network_v2.example
}