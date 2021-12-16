
data "dnacenter_network_device_global_polling_interval" "example" {
    provider = dnac
}

output "dnacenter_network_device_global_polling_interval_example" {
    value = data.dnac_network_device_global_polling_interval.example.item
}
