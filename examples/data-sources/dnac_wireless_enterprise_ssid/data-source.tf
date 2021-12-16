
data "dnac_wireless_enterprise_ssid" "example" {
    provider = dnac
    ssid_name = "string"
}

output "dnac_wireless_enterprise_ssid_example" {
    value = data.dnac_wireless_enterprise_ssid.example.items
}
