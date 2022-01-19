
resource "dnac_endpoint_analytics_profiling_rules" "example" {
    provider = dnac
    item {
      
      # cluster_id = ------
      condition_groups {
        
        condition {
          
          # attribute = ------
          # attribute_dictionary = ------
          # operator = ------
          # value = ------
        }
        # condition_group = [------]
        # operator = ------
        # type = ------
      }
      # is_deleted = ------
      # last_modified_by = ------
      # last_modified_on = ------
      # plugin_id = ------
      # rejected = ------
      result {
        
        # device_type = [------]
        # hardware_manufacturer = [------]
        # hardware_model = [------]
        # operating_system = [------]
      }
      # rule_id = ------
      # rule_name = ------
      # rule_priority = ------
      # rule_type = ------
      # rule_version = ------
      # source_priority = ------
      # used_attributes = [------]
    }
    parameters {
      
      cluster_id = "string"
      condition_groups {
        
        condition {
          
          attribute = "string"
          attribute_dictionary = "string"
          operator = "string"
          value = "string"
        }
        condition_group = ["string"]
        operator = "string"
        type = "string"
      }
      is_deleted = "false"
      last_modified_by = "string"
      last_modified_on = 1
      plugin_id = "string"
      rejected = "false"
      result {
        
        device_type = ["string"]
        hardware_manufacturer = ["string"]
        hardware_model = ["string"]
        operating_system = ["string"]
      }
      rule_id = "string"
      rule_name = "string"
      rule_priority = 1
      rule_type = "string"
      rule_version = 1
      source_priority = 1
      used_attributes = ["string"]
    }
}

output "dnac_endpoint_analytics_profiling_rules_example" {
    value = dnac_endpoint_analytics_profiling_rules.example
}