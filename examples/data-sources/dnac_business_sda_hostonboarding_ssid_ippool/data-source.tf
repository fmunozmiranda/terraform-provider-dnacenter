
data "dnac_business_sda_hostonboarding_ssid_ippool" "example" {
    provider = dnac
    site_name_hierarchy = "string"
    vlan_name = "string"
}

output "dnac_business_sda_hostonboarding_ssid_ippool_example" {
    value = data.dnac_business_sda_hostonboarding_ssid_ippool.example.item
}
