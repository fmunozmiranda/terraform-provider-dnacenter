
data "dnac_discovery" "example" {
    provider = dnac
    id = "string"
}

output "dnac_discovery_example" {
    value = data.dnac_discovery.example.item
}
