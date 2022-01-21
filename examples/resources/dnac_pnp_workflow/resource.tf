
resource "dnacenter_pnp_workflow" "example" {
    provider = dnacenter
    item {
      
      # id = ------
      # add_to_inventory = ------
      # added_on = ------
      # config_id = ------
      # curr_task_idx = ------
      # description = ------
      # end_time = ------
      # exec_time = ------
      # image_id = ------
      # instance_type = ------
      # lastupdate_on = ------
      # name = ------
      # start_time = ------
      # state = ------
      tasks {
        
        # curr_work_item_idx = ------
        # end_time = ------
        # name = ------
        # start_time = ------
        # state = ------
        # task_seq_no = ------
        # time_taken = ------
        # type = ------
        work_item_list {
          
          # command = ------
          # end_time = ------
          # output_str = ------
          # start_time = ------
          # state = ------
          # time_taken = ------
        }
      }
      # tenant_id = ------
      # type = ------
      # use_state = ------
      # version = ------
    }
    parameters {
      
      id = "string"
      add_to_inventory = "false"
      added_on = 1
      config_id = "string"
      curr_task_idx = 1
      description = "string"
      end_time = 1
      exec_time = 1
      image_id = "string"
      instance_type = "string"
      lastupdate_on = 1
      name = "string"
      start_time = 1
      state = "string"
      tasks {
        
        curr_work_item_idx = 1
        end_time = 1
        name = "string"
        start_time = 1
        state = "string"
        task_seq_no = 1
        time_taken = 1
        type = "string"
        work_item_list {
          
          command = "string"
          end_time = 1
          output_str = "string"
          start_time = 1
          state = "string"
          time_taken = 1
        }
      }
      tenant_id = "string"
      type = "string"
      use_state = "string"
      version = 1
    }
}

output "dnacenter_pnp_workflow_example" {
    value = dnacenter_pnp_workflow.example
}