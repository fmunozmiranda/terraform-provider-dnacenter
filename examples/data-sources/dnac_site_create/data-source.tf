
data "dnacsite_create" "example" {
    provider = dnac
    area {
      
      name = "string"
      parent_name = "string"
    }
    building {
      
      address = "string"
      latitude = ------
      longitude = ------
      name = "string"
      parent_name = "string"
    }
    floor {
      
      height = ------
      length = ------
      name = "string"
      parent_name = "string"
      rf_model = "string"
      width = ------
    }
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
}