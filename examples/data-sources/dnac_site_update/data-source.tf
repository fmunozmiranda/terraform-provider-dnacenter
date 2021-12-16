
data "dnacenter_site_update" "example" {
    provider = dnac
    site_id = "string"
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
      rf_model = "string"
      width = ------
    }
    item {
      
      # data = ------
      # end_time = ------
      # id = ------
      # instance_tenant_id = ------
      # is_error = ------
      # operation_id_list = [------]
      # progress = ------
      # root_id = ------
      # service_type = ------
      # start_time = ------
      # version = ------
    }
}