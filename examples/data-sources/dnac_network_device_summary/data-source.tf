
data "dnac_network_device_summary" "example" {
    provider = dnac
    id = "string"
}

output "dnac_network_device_summary_example" {
    value = data.dnac_network_device_summary.example.item
}
