
data "dnacsensor_test_template_duplicate" "example" {
    provider = dnac
    item {
      
      # id = ------
      ap_coverage {
        
        # bands = ------
        # number_of_aps_to_test = ------
        # rssi_threshold = ------
      }
      # connection = ------
      # encryption_mode = ------
      # last_modified_time = ------
      # legacy_test_suite = ------
      location_info_list {
        
        # all_sensors = ------
        # location_id = ------
        # location_type = ------
        # site_hierarchy = ------
      }
      # model_version = ------
      # name = ------
      # num_associated_sensor = ------
      # num_neighbor_apthreshold = ------
      # radio_as_sensor_removed = ------
      # rssi_threshold = ------
      # run_now = ------
      schedule {
        
        frequency {
          
          # unit = ------
          # value = ------
        }
        schedule_range {
          
          # day = ------
          time_range {
            
            frequency {
              
              # unit = ------
              # value = ------
            }
            # from = ------
            # to = ------
          }
        }
        # start_time = ------
        # test_schedule_mode = ------
      }
      # schedule_in_days = ------
      # show_wlc_upgrade_banner = ------
      ssids {
        
        # auth_type = ------
        # certstatus = ------
        # certxferprotocol = ------
        # ext_web_auth = ------
        # id = ------
        # num_aps = ------
        # num_sensors = ------
        # profile_name = ------
        # psk = ------
        # qos_policy = ------
        # scep = ------
        # ssid = ------
        # status = ------
        tests {
          
          # name = ------
        }
        third_party {
          
          # selected = ------
        }
        # valid_from = ------
        # valid_to = ------
        # white_list = ------
        # wlan_id = ------
      }
      # start_time = ------
      # status = ------
      # test_duration_estimate = ------
      # test_schedule_mode = ------
      # test_template = ------
      # version = ------
    }
    new_template_name = "string"
    template_name = "string"
}