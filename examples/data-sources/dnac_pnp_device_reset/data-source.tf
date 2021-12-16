
data "dnacpnp_device_reset" "example" {
    provider = dnac
    config_list {
      
      config_id = "string"
      config_parameters {
        
        key = "string"
        value = "string"
      }
    }
    device_id = "string"
    item {
      
      # message = ------
      # status_code = ------
    }
    license_level = "string"
    license_type = "string"
    top_of_stack_serial_number = "string"
}