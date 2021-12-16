
data "dnacservice_provider_create" "example" {
    provider = dnac
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    qos {
      
      model = "string"
      profile_name = "string"
      wan_provider = "string"
    }
}