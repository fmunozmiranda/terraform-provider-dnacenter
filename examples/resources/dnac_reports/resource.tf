
resource "dnac_reports" "example" {
    provider = dnac
    item {
      
      # data_category = ------
      # deliveries = [------]
      # execution_count = ------
      executions {
        
        # end_time = ------
        # errors = [------]
        # execution_id = ------
        # process_status = ------
        # request_status = ------
        # start_time = ------
        # warnings = [------]
      }
      # name = ------
      # report_id = ------
      # report_was_executed = ------
      # schedule = [------]
      # tags = [------]
      view {
        
        # description = ------
        field_groups {
          
          # field_group_display_name = ------
          # field_group_name = ------
          fields {
            
            # display_name = ------
            # name = ------
          }
        }
        filters {
          
          # display_name = ------
          # name = ------
          # type = ------
          # value = [------]
        }
        format {
          
          # default = ------
          # format_type = ------
          # name = ------
        }
        # name = ------
        # view_id = ------
        # view_info = ------
      }
      # view_group_id = ------
      # view_group_version = ------
    }
    parameters {
      
      deliveries = ["string"]
      name = "string"
      report_id = "string"
      schedule = ["string"]
      tags = ["string"]
      view {
        
        field_groups {
          
          field_group_display_name = "string"
          field_group_name = "string"
          fields {
            
            display_name = "string"
            name = "string"
          }
        }
        filters {
          
          display_name = "string"
          name = "string"
          type = "string"
          value = ["string"]
        }
        format {
          
          format_type = "string"
          name = "string"
        }
        name = "string"
        view_id = "string"
      }
      view_group_id = "string"
      view_group_version = "string"
    }
}

output "dnac_reports_example" {
    value = dnac_reports.example
}