
data "dnac_sda_virtual_network_ip_pool" "example" {
    provider = dnac
    ip_pool_name = "string"
    virtual_network_name = "string"
}

output "dnac_sda_virtual_network_ip_pool_example" {
    value = data.dnac_sda_virtual_network_ip_pool.example.item
}
