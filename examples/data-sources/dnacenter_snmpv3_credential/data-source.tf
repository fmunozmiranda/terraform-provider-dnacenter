
data "dnacdnacenter_snmpv3_credential" "example" {
  provider = dnacenter
  payload {

    auth_password      = "string"
    auth_type          = "string"
    comments           = "string"
    credential_type    = "string"
    description        = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    privacy_password   = "string"
    privacy_type       = "string"
    snmp_mode          = "string"
    username           = "string"
  }
}