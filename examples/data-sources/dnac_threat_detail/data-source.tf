
data "dnacenter_threat_detail" "example" {
    provider = dnac
    end_time = 1
    is_new_threat = "false"
    items {
      
      # ap_name = ------
      # mac_address = ------
      # site_name_hierarchy = ------
      # ssid = ------
      # threat_level = ------
      # threat_type = ------
      # updated_time = ------
      # vendor = ------
    }
    limit = 1
    offset = 1
    site_id = ["string"]
    start_time = 1
    threat_level = ["string"]
    threat_type = ["string"]
}