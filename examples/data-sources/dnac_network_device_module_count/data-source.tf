
data "dnac_network_device_module_count" "example" {
    provider = dnac
    device_id = "string"
    name_list = ["string"]
    operational_state_code_list = ["string"]
    part_number_list = ["string"]
    vendor_equipment_type_list = ["string"]
}

output "dnac_network_device_module_count_example" {
    value = data.dnac_network_device_module_count.example.item
}