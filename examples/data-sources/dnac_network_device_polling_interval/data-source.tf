
data "dnac_network_device_polling_interval" "example" {
    provider = dnac
    id = "string"
}

output "dnac_network_device_polling_interval_example" {
    value = data.dnac_network_device_polling_interval.example.item
}