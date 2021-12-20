
data "dnacdnacenter_wireless_psk_override" "example" {
  provider = dnacenter
  payload {

    pass_phrase = "string"
    site        = "string"
    ssid        = "string"
  }
}