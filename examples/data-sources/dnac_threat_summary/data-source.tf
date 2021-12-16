
data "dnacenter_threat_summary" "example" {
    provider = dnac
    end_time = 1
    items {
      
      threat_data {
        
        # threat_count = ------
        # threat_level = ------
        # threat_type = ------
      }
      # timestamp = ------
    }
    site_id = ["string"]
    start_time = 1
    threat_level = ["string"]
    threat_type = ["string"]
}