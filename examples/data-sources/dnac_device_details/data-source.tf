
data "dnac_device_details" "example" {
    provider = dnac
    identifier = "string"
    search_by = "string"
    timestamp = "string"
}

output "dnac_device_details_example" {
    value = data.dnac_device_details.example.item
}
