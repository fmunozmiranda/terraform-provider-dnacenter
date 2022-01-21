
resource "dnacenter_configuration_template" "example" {
    provider = dnacenter
    item {
      
      # author = ------
      # composite = ------
      containing_templates {
        
        # composite = ------
        # description = ------
        device_types {
          
          # product_family = ------
          # product_series = ------
          # product_type = ------
        }
        # id = ------
        # language = ------
        # name = ------
        # project_name = ------
        rollback_template_params {
          
          # binding = ------
          # custom_order = ------
          # data_type = ------
          # default_value = ------
          # description = ------
          # display_name = ------
          # group = ------
          # id = ------
          # instruction_text = ------
          # key = ------
          # not_param = ------
          # order = ------
          # param_array = ------
          # parameter_name = ------
          # provider = ------
          range {
            
            # id = ------
            # max_value = ------
            # min_value = ------
          }
          # required = ------
          selection {
            
            # default_selected_values = [------]
            # id = ------
            # selection_type = ------
            # selection_values = [------]
          }
        }
        tags {
          
          # id = ------
          # name = ------
        }
        # template_content = ------
        template_params {
          
          # binding = ------
          # custom_order = ------
          # data_type = ------
          # default_value = ------
          # description = ------
          # display_name = ------
          # group = ------
          # id = ------
          # instruction_text = ------
          # key = ------
          # not_param = ------
          # order = ------
          # param_array = ------
          # parameter_name = ------
          # provider = ------
          range {
            
            # id = ------
            # max_value = ------
            # min_value = ------
          }
          # required = ------
          selection {
            
            # default_selected_values = [------]
            # id = ------
            # selection_type = ------
            # selection_values = [------]
          }
        }
        # version = ------
      }
      # create_time = ------
      # custom_params_order = ------
      # description = ------
      device_types {
        
        # product_family = ------
        # product_series = ------
        # product_type = ------
      }
      # document_database = ------
      # failure_policy = ------
      # id = ------
      # language = ------
      # last_update_time = ------
      # latest_version_time = ------
      # name = ------
      # parent_template_id = ------
      # project_associated = ------
      # project_id = ------
      # project_name = ------
      # rollback_template_content = ------
      rollback_template_params {
        
        # binding = ------
        # custom_order = ------
        # data_type = ------
        # default_value = ------
        # description = ------
        # display_name = ------
        # group = ------
        # id = ------
        # instruction_text = ------
        # key = ------
        # not_param = ------
        # order = ------
        # param_array = ------
        # parameter_name = ------
        # provider = ------
        range {
          
          # id = ------
          # max_value = ------
          # min_value = ------
        }
        # required = ------
        selection {
          
          # default_selected_values = [------]
          # id = ------
          # selection_type = ------
          # selection_values = [------]
        }
      }
      # software_type = ------
      # software_variant = ------
      # software_version = ------
      tags {
        
        # id = ------
        # name = ------
      }
      # template_content = ------
      template_params {
        
        # binding = ------
        # custom_order = ------
        # data_type = ------
        # default_value = ------
        # description = ------
        # display_name = ------
        # group = ------
        # id = ------
        # instruction_text = ------
        # key = ------
        # not_param = ------
        # order = ------
        # param_array = ------
        # parameter_name = ------
        # provider = ------
        range {
          
          # id = ------
          # max_value = ------
          # min_value = ------
        }
        # required = ------
        selection {
          
          # default_selected_values = [------]
          # id = ------
          # selection_type = ------
          # selection_values = [------]
        }
      }
      validation_errors {
        
        # rollback_template_errors = [------]
        # template_errors = [------]
        # template_id = ------
        # template_version = ------
      }
      # version = ------
    }
    parameters {
      
      author = "string"
      composite = "false"
      containing_templates {
        
        composite = "false"
        description = "string"
        device_types {
          
          product_family = "string"
          product_series = "string"
          product_type = "string"
        }
        id = "string"
        language = "string"
        name = "string"
        project_name = "string"
        rollback_template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        tags {
          
          id = "string"
          name = "string"
        }
        template_content = "string"
        template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        version = "string"
      }
      create_time = 1
      custom_params_order = "false"
      description = "string"
      device_types {
        
        product_family = "string"
        product_series = "string"
        product_type = "string"
      }
      failure_policy = "string"
      id = "string"
      language = "string"
      last_update_time = 1
      latest_version_time = 1
      name = "string"
      parent_template_id = "string"
      project_id = "string"
      project_name = "string"
      rollback_template_content = "string"
      rollback_template_params {
        
        binding = "string"
        custom_order = 1
        data_type = "string"
        default_value = "string"
        description = "string"
        display_name = "string"
        group = "string"
        id = "string"
        instruction_text = "string"
        key = "string"
        not_param = "false"
        order = 1
        param_array = "false"
        parameter_name = "string"
        provider = "string"
        range {
          
          id = "string"
          max_value = 1
          min_value = 1
        }
        required = "false"
        selection {
          
          default_selected_values = ["string"]
          id = "string"
          selection_type = "string"
          selection_values = ["string"]
        }
      }
      software_type = "string"
      software_variant = "string"
      software_version = "string"
      tags {
        
        id = "string"
        name = "string"
      }
      template_content = "string"
      template_id = "string"
      template_params {
        
        binding = "string"
        custom_order = 1
        data_type = "string"
        default_value = "string"
        description = "string"
        display_name = "string"
        group = "string"
        id = "string"
        instruction_text = "string"
        key = "string"
        not_param = "false"
        order = 1
        param_array = "false"
        parameter_name = "string"
        provider = "string"
        range {
          
          id = "string"
          max_value = 1
          min_value = 1
        }
        required = "false"
        selection {
          
          default_selected_values = ["string"]
          id = "string"
          selection_type = "string"
          selection_values = ["string"]
        }
      }
      validation_errors {
        
        rollback_template_errors = ["string"]
        template_errors = ["string"]
        template_id = "string"
        template_version = "string"
      }
      version = "string"
    }
}

output "dnacenter_configuration_template_example" {
    value = dnacenter_configuration_template.example
}