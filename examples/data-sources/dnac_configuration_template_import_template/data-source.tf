
data "dnacconfiguration_template_import_template" "example" {
    provider = dnac
    do_version = "false"
    project_name = "string"
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
    item {
      
      # task_id = ------
      # url = ------
    }
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
      }
    }
    validation_errors {
      
      template_id = "string"
      template_version = "string"
    }
    version = "string"
}