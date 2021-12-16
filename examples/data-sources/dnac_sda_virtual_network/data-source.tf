
data "dnac_sda_virtual_network" "example" {
    provider = dnac
    site_name_hierarchy = "string"
    virtual_network_name = "string"
}

output "dnac_sda_virtual_network_example" {
    value = data.dnac_sda_virtual_network.example.item
}
