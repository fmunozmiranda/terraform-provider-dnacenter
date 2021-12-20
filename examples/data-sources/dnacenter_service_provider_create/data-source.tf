
data "dnacdnacenter_service_provider_create" "example" {
  provider = dnacenter
  qos {

    model        = "string"
    profile_name = "string"
    wan_provider = "string"
  }
}