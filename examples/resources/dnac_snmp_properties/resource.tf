
resource "dnac_snmp_properties" "example" {
    provider = dnac
    parameters {
      
      id = "string"
      instance_tenant_id = "string"
      instance_uuid = "string"
      int_value = 1
      system_property_name = "string"
    }
}

output "dnac_snmp_properties_example" {
    value = dnac_snmp_properties.example
}