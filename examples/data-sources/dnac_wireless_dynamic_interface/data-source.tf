
data "dnac_wireless_dynamic_interface" "example" {
    provider = dnac
    interface_name = "string"
}

output "dnac_wireless_dynamic_interface_example" {
    value = data.dnac_wireless_dynamic_interface.example.items
}
