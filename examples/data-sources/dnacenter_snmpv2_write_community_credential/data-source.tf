
data "dnacdnacenter_snmpv2_write_community_credential" "example" {
  provider = dnacenter
  payload {

    comments        = "string"
    credential_type = "string"
    description     = "string"
    instance_uuid   = "string"
    write_community = "string"
  }
}