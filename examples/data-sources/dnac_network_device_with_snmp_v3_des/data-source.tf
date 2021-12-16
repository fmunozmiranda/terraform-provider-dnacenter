
data "dnac_network_device_with_snmp_v3_des" "example" {
    provider = dnac
    limit = "string"
    offset = "string"
    order = "string"
    site_id = "string"
    sort_by = "string"
}

output "dnac_network_device_with_snmp_v3_des_example" {
    value = data.dnac_network_device_with_snmp_v3_des.example.items
}