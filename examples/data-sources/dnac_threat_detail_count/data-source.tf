
data "dnacthreat_detail_count" "example" {
    provider = dnac
    end_time = 1
    is_new_threat = "false"
    item {
      
      # response = ------
      # version = ------
    }
    limit = 1
    offset = 1
    site_id = ["string"]
    start_time = 1
    threat_level = ["string"]
    threat_type = ["string"]
}