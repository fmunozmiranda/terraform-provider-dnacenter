
data "dnacenter_sda_virtual_network_v2" "example" {
    provider = dnac
    virtual_network_name = "string"
}

output "dnacenter_sda_virtual_network_v2_example" {
    value = data.dnac_sda_virtual_network_v2.example.item
}
