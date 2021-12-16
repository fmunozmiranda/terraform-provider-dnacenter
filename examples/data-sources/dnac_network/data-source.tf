
data "dnac_network" "example" {
    provider = dnac
    site_id = "string"
}

output "dnac_network_example" {
    value = data.dnac_network.example.items
}
