
resource "dnac_business_sda_hostonboarding_ssid_ippool" "example" {
    provider = dnac
    parameters {
      
      scalable_group_name = "string"
      site_name_hierarchy = "string"
      ssid_names = ["string"]
      vlan_name = "string"
    }
}

output "dnac_business_sda_hostonboarding_ssid_ippool_example" {
    value = dnac_business_sda_hostonboarding_ssid_ippool.example
}