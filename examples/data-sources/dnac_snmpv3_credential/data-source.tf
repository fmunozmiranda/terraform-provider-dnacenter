
data "dnacenter_snmpv3_credential" "example" {
    provider = dnac
    auth_password = "string"
    auth_type = "string"
    comments = "string"
    credential_type = "string"
    description = "string"
    id = "string"
    instance_tenant_id = "string"
    instance_uuid = "string"
    item_id {
      
      # task_id = ------
      # url = ------
    }
    item_name {
      
      # task_id = ------
      # url = ------
    }
    privacy_password = "string"
    privacy_type = "string"
    snmp_mode = "string"
    username = "string"
}