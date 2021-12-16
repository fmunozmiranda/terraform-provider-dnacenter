
data "dnacenter_service_provider_update" "example" {
    provider = dnac
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    qos {
      
      model = "string"
      old_profile_name = "string"
      profile_name = "string"
      wan_provider = "string"
    }
}