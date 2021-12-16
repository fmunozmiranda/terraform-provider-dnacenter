
data "dnac_network_device_equipment" "example" {
    provider = dnac
    device_uuid = "string"
    type = "string"
}

output "dnac_network_device_equipment_example" {
    value = data.dnac_network_device_equipment.example.items
}