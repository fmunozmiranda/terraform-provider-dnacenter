
data "dnacdnacenter_swim_trigger_distribution" "example" {
  provider = dnacenter
  payload {

    device_uuid = "string"
    image_uuid  = "string"
  }
}