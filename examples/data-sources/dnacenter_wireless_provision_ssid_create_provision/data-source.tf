
data "dnacdnacenter_wireless_provision_ssid_create_provision" "example" {
  provider               = dnacenter
  enable_broadcast_ssi_d = "false"
  enable_fast_lane       = "false"
  enable_mac_filtering   = "false"
  fast_transition        = "string"
  name                   = "string"
  passphrase             = "string"
  radio_policy           = "string"
  security_level         = "string"
  traffic_type           = "string"
  web_auth_url           = "string"
}