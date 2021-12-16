
data "dnacenter_discovery" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_discovery_example" {
    value = data.dnac_discovery.example.item
}
