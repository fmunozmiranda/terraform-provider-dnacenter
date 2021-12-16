
data "dnacenter_topology_site" "example" {
    provider = dnac
}

output "dnacenter_topology_site_example" {
    value = data.dnac_topology_site.example.item
}
