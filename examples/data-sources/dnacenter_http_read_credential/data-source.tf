
data "dnacdnacenter_http_read_credential" "example" {
  provider = dnacenter
  payload {

    comments           = "string"
    credential_type    = "string"
    description        = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "******"
    port               = 1
    secure             = "false"
    username           = "string"
  }
}