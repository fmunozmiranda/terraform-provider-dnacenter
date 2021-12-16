
data "dnacenter_wireless_provision_ssid_create_provision" "example" {
    provider = dnac
    enable_broadcast_ssi_d = "false"
    enable_fast_lane = "false"
    enable_mac_filtering = "false"
    fast_transition = "string"
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    name = "string"
    passphrase = "string"
    radio_policy = "string"
    security_level = "string"
    traffic_type = "string"
    web_auth_url = "string"
}