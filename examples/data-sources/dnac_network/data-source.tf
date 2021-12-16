
data "dnacenter_network" "example" {
    provider = dnac
    site_id = "string"
}

output "dnacenter_network_example" {
    value = data.dnac_network.example.items
}
