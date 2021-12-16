
data "dnac_sda_multicast" "example" {
    provider = dnac
    site_name_hierarchy = "string"
}

output "dnac_sda_multicast_example" {
    value = data.dnac_sda_multicast.example.item
}
