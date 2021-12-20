
data "dnacdnacenter_netconf_credential" "example" {
  provider = dnacenter
  payload {

    comments           = "string"
    credential_type    = "string"
    description        = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    netconf_port       = "string"
  }
}