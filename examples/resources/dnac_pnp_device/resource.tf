
resource "dnacenter_pnp_device" "example" {
    provider = dnacenter
    item {
      
      # id = ------
      day_zero_config {
        
        # config = ------
      }
      # day_zero_config_preview = [------]
      device_info {
        
        aaa_credentials {
          
          # password = ------
          # username = ------
        }
        # added_on = ------
        # addn_mac_addrs = [------]
        # agent_type = ------
        # auth_status = ------
        # authenticated_mic_number = ------
        # authenticated_sudi_serial_no = ------
        # capabilities_supported = [------]
        # cm_state = ------
        # description = ------
        # device_sudi_serial_nos = [------]
        # device_type = ------
        # features_supported = [------]
        file_system_list {
          
          # freespace = ------
          # name = ------
          # readable = ------
          # size = ------
          # type = ------
          # writeable = ------
        }
        # first_contact = ------
        # hostname = ------
        http_headers {
          
          # key = ------
          # value = ------
        }
        # image_file = ------
        # image_version = ------
        ip_interfaces {
          
          # ipv4_address = [------]
          # ipv6_address_list = [------]
          # mac_address = ------
          # name = ------
          # status = ------
        }
        # last_contact = ------
        # last_sync_time = ------
        # last_update_on = ------
        location {
          
          # address = ------
          # altitude = ------
          # latitude = ------
          # longitude = ------
          # site_id = ------
        }
        # mac_address = ------
        # mode = ------
        # name = ------
        neighbor_links {
          
          # local_interface_name = ------
          # local_mac_address = ------
          # local_short_interface_name = ------
          # remote_device_name = ------
          # remote_interface_name = ------
          # remote_mac_address = ------
          # remote_platform = ------
          # remote_short_interface_name = ------
          # remote_version = ------
        }
        # onb_state = ------
        # pid = ------
        pnp_profile_list {
          
          # created_by = ------
          # discovery_created = ------
          primary_endpoint {
            
            # certificate = ------
            # fqdn = ------
            # ipv4_address = [------]
            # ipv6_address = [------]
            # port = ------
            # protocol = ------
          }
          # profile_name = ------
          secondary_endpoint {
            
            # certificate = ------
            # fqdn = ------
            # ipv4_address = [------]
            # ipv6_address = [------]
            # port = ------
            # protocol = ------
          }
        }
        # populate_inventory = ------
        pre_workflow_cli_ouputs {
          
          # cli = ------
          # cli_output = ------
        }
        # project_id = ------
        # project_name = ------
        # reload_requested = ------
        # serial_number = ------
        # site_id = ------
        # site_name = ------
        # smart_account_id = ------
        # source = ------
        # stack = ------
        stack_info {
          
          # is_full_ring = ------
          stack_member_list {
            
            # hardware_version = ------
            # license_level = ------
            # license_type = ------
            # mac_address = ------
            # pid = ------
            # priority = ------
            # role = ------
            # serial_number = ------
            # software_version = ------
            # stack_number = ------
            # state = ------
            # sudi_serial_number = ------
          }
          # stack_ring_protocol = ------
          # supports_stack_workflows = ------
          # total_member_count = ------
          # valid_license_levels = [------]
        }
        # state = ------
        # sudi_required = ------
        # tags = [------]
        # user_mic_numbers = [------]
        # user_sudi_serial_nos = [------]
        # virtual_account_id = ------
        # workflow_id = ------
        # workflow_name = ------
      }
      run_summary_list {
        
        # details = ------
        # error_flag = ------
        history_task_info {
          
          addn_details {
            
            # key = ------
            # value = ------
          }
          # name = ------
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
        # timestamp = ------
      }
      system_reset_workflow {
        
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
      system_workflow {
        
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
      # tenant_id = ------
      # version = ------
      workflow {
        
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
      workflow_parameters {
        
        config_list {
          
          # config_id = ------
          config_parameters {
            
            # key = ------
            # value = ------
          }
        }
        # license_level = ------
        # license_type = ------
        # top_of_stack_serial_number = ------
      }
    }
    parameters {
      
      id = "string"
      device_info {
        
        aaa_credentials {
          
          password = "******"
          username = "string"
        }
        added_on = 1
        addn_mac_addrs = ["string"]
        agent_type = "string"
        auth_status = "string"
        authenticated_sudi_serial_no = "string"
        capabilities_supported = ["string"]
        cm_state = "string"
        description = "string"
        device_sudi_serial_nos = ["string"]
        device_type = "string"
        features_supported = ["string"]
        file_system_list {
          
          freespace = 1
          name = "string"
          readable = "false"
          size = 1
          type = "string"
          writeable = "false"
        }
        first_contact = 1
        hostname = "string"
        http_headers {
          
          key = "string"
          value = "string"
        }
        image_file = "string"
        image_version = "string"
        ip_interfaces {
          
          ipv4_address = ["string"]
          ipv6_address_list = ["string"]
          mac_address = "string"
          name = "string"
          status = "string"
        }
        last_contact = 1
        last_sync_time = 1
        last_update_on = 1
        location {
          
          address = "string"
          altitude = "string"
          latitude = "string"
          longitude = "string"
          site_id = "string"
        }
        mac_address = "string"
        mode = "string"
        name = "string"
        neighbor_links {
          
          local_interface_name = "string"
          local_mac_address = "string"
          local_short_interface_name = "string"
          remote_device_name = "string"
          remote_interface_name = "string"
          remote_mac_address = "string"
          remote_platform = "string"
          remote_short_interface_name = "string"
          remote_version = "string"
        }
        onb_state = "string"
        pid = "string"
        pnp_profile_list {
          
          created_by = "string"
          discovery_created = "false"
          primary_endpoint {
            
            certificate = "string"
            fqdn = "string"
            ipv4_address = ["string"]
            ipv6_address = ["string"]
            port = 1
            protocol = "string"
          }
          profile_name = "string"
          secondary_endpoint {
            
            certificate = "string"
            fqdn = "string"
            ipv4_address = ["string"]
            ipv6_address = ["string"]
            port = 1
            protocol = "string"
          }
        }
        populate_inventory = "false"
        pre_workflow_cli_ouputs {
          
          cli = "string"
          cli_output = "string"
        }
        project_id = "string"
        project_name = "string"
        reload_requested = "false"
        serial_number = "string"
        smart_account_id = "string"
        source = "string"
        stack = "false"
        stack_info {
          
          is_full_ring = "false"
          stack_member_list {
            
            hardware_version = "string"
            license_level = "string"
            license_type = "string"
            mac_address = "string"
            pid = "string"
            priority = 1
            role = "string"
            serial_number = "string"
            software_version = "string"
            stack_number = 1
            state = "string"
            sudi_serial_number = "string"
          }
          stack_ring_protocol = "string"
          supports_stack_workflows = "false"
          total_member_count = 1
          valid_license_levels = ["string"]
        }
        state = "string"
        sudi_required = "false"
        tags = ["string"]
        user_sudi_serial_nos = ["string"]
        virtual_account_id = "string"
        workflow_id = "string"
        workflow_name = "string"
      }
      run_summary_list {
        
        details = "string"
        error_flag = "false"
        history_task_info {
          
          addn_details {
            
            key = "string"
            value = "string"
          }
          name = "string"
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
        timestamp = 1
      }
      system_reset_workflow {
        
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
      system_workflow {
        
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
      tenant_id = "string"
      version = 1
      workflow {
        
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
      workflow_parameters {
        
        config_list {
          
          config_id = "string"
          config_parameters {
            
            key = "string"
            value = "string"
          }
        }
        license_level = "string"
        license_type = "string"
        top_of_stack_serial_number = "string"
      }
    }
}

output "dnacenter_pnp_device_example" {
    value = dnacenter_pnp_device.example
}