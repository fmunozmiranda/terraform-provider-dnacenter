
data "dnac_device_replacement_count" "example" {
    provider = dnac
    replacement_status = ["string"]
}

output "dnac_device_replacement_count_example" {
    value = data.dnac_device_replacement_count.example.item
}
