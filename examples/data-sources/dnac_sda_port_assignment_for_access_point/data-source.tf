
data "dnac_sda_port_assignment_for_access_point" "example" {
    provider = dnac
    device_management_ip_address = "string"
    interface_name = "string"
}

output "dnac_sda_port_assignment_for_access_point_example" {
    value = data.dnac_sda_port_assignment_for_access_point.example.item
}
