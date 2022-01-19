
resource "dnac_wireless_dynamic_interface" "example" {
    provider = dnac
    parameters {
      
      interface_name = "string"
      vlan_id = ------
    }
}

output "dnac_wireless_dynamic_interface_example" {
    value = dnac_wireless_dynamic_interface.example
}