
data "dnac_topology_site" "example" {
    provider = dnac
}

output "dnac_topology_site_example" {
    value = data.dnac_topology_site.example.item
}
