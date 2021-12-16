
data "dnacenter_snmp_properties" "example" {
    provider = dnac
}

output "dnacenter_snmp_properties_example" {
    value = data.dnac_snmp_properties.example.items
}
