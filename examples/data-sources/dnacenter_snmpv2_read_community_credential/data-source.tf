
data "dnacdnacenter_snmpv2_read_community_credential" "example" {
  provider = dnacenter
  payload {

    comments        = "string"
    credential_type = "string"
    description     = "string"
    instance_uuid   = "string"
    read_community  = "string"
  }
}