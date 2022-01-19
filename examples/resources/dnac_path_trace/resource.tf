
resource "dnac_path_trace" "example" {
    provider = dnac
    item {
      
      detailed_status {
        
        # acl_trace_calculation = ------
        # acl_trace_calculation_failure_reason = ------
      }
      # last_update = ------
      network_elements {
        
        accuracy_list {
          
          # percent = ------
          # reason = ------
        }
        detailed_status {
          
          # acl_trace_calculation = ------
          # acl_trace_calculation_failure_reason = ------
        }
        device_statistics {
          
          cpu_statistics {
            
            # five_min_usage_in_percentage = ------
            # five_secs_usage_in_percentage = ------
            # one_min_usage_in_percentage = ------
            # refreshed_at = ------
          }
          memory_statistics {
            
            # memory_usage = ------
            # refreshed_at = ------
            # total_memory = ------
          }
        }
        # device_stats_collection = ------
        # device_stats_collection_failure_reason = ------
        egress_physical_interface {
          
          acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # id = ------
          interface_statistics {
            
            # admin_status = ------
            # input_packets = ------
            # input_queue_count = ------
            # input_queue_drops = ------
            # input_queue_flushes = ------
            # input_queue_max_depth = ------
            # input_ratebps = ------
            # operational_status = ------
            # output_drop = ------
            # output_packets = ------
            # output_queue_count = ------
            # output_queue_depth = ------
            # output_ratebps = ------
            # refreshed_at = ------
          }
          # interface_stats_collection = ------
          # interface_stats_collection_failure_reason = ------
          # name = ------
          path_overlay_info {
            
            # control_plane = ------
            # data_packet_encapsulation = ------
            # dest_ip = ------
            # dest_port = ------
            # protocol = ------
            # source_ip = ------
            # source_port = ------
            vxlan_info {
              
              # dscp = ------
              # vnid = ------
            }
          }
          qos_statistics {
            
            # class_map_name = ------
            # drop_rate = ------
            # num_bytes = ------
            # num_packets = ------
            # offered_rate = ------
            # queue_bandwidthbps = ------
            # queue_depth = ------
            # queue_no_buffer_drops = ------
            # queue_total_drops = ------
            # refreshed_at = ------
          }
          # qos_stats_collection = ------
          # qos_stats_collection_failure_reason = ------
          # used_vlan = ------
          # vrf_name = ------
        }
        egress_virtual_interface {
          
          acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # id = ------
          interface_statistics {
            
            # admin_status = ------
            # input_packets = ------
            # input_queue_count = ------
            # input_queue_drops = ------
            # input_queue_flushes = ------
            # input_queue_max_depth = ------
            # input_ratebps = ------
            # operational_status = ------
            # output_drop = ------
            # output_packets = ------
            # output_queue_count = ------
            # output_queue_depth = ------
            # output_ratebps = ------
            # refreshed_at = ------
          }
          # interface_stats_collection = ------
          # interface_stats_collection_failure_reason = ------
          # name = ------
          path_overlay_info {
            
            # control_plane = ------
            # data_packet_encapsulation = ------
            # dest_ip = ------
            # dest_port = ------
            # protocol = ------
            # source_ip = ------
            # source_port = ------
            vxlan_info {
              
              # dscp = ------
              # vnid = ------
            }
          }
          qos_statistics {
            
            # class_map_name = ------
            # drop_rate = ------
            # num_bytes = ------
            # num_packets = ------
            # offered_rate = ------
            # queue_bandwidthbps = ------
            # queue_depth = ------
            # queue_no_buffer_drops = ------
            # queue_total_drops = ------
            # refreshed_at = ------
          }
          # qos_stats_collection = ------
          # qos_stats_collection_failure_reason = ------
          # used_vlan = ------
          # vrf_name = ------
        }
        flex_connect {
          
          # authentication = ------
          # data_switching = ------
          egress_acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          ingress_acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # wireless_lan_controller_id = ------
          # wireless_lan_controller_name = ------
        }
        # id = ------
        ingress_physical_interface {
          
          acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # id = ------
          interface_statistics {
            
            # admin_status = ------
            # input_packets = ------
            # input_queue_count = ------
            # input_queue_drops = ------
            # input_queue_flushes = ------
            # input_queue_max_depth = ------
            # input_ratebps = ------
            # operational_status = ------
            # output_drop = ------
            # output_packets = ------
            # output_queue_count = ------
            # output_queue_depth = ------
            # output_ratebps = ------
            # refreshed_at = ------
          }
          # interface_stats_collection = ------
          # interface_stats_collection_failure_reason = ------
          # name = ------
          path_overlay_info {
            
            # control_plane = ------
            # data_packet_encapsulation = ------
            # dest_ip = ------
            # dest_port = ------
            # protocol = ------
            # source_ip = ------
            # source_port = ------
            vxlan_info {
              
              # dscp = ------
              # vnid = ------
            }
          }
          qos_statistics {
            
            # class_map_name = ------
            # drop_rate = ------
            # num_bytes = ------
            # num_packets = ------
            # offered_rate = ------
            # queue_bandwidthbps = ------
            # queue_depth = ------
            # queue_no_buffer_drops = ------
            # queue_total_drops = ------
            # refreshed_at = ------
          }
          # qos_stats_collection = ------
          # qos_stats_collection_failure_reason = ------
          # used_vlan = ------
          # vrf_name = ------
        }
        ingress_virtual_interface {
          
          acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # id = ------
          interface_statistics {
            
            # admin_status = ------
            # input_packets = ------
            # input_queue_count = ------
            # input_queue_drops = ------
            # input_queue_flushes = ------
            # input_queue_max_depth = ------
            # input_ratebps = ------
            # operational_status = ------
            # output_drop = ------
            # output_packets = ------
            # output_queue_count = ------
            # output_queue_depth = ------
            # output_ratebps = ------
            # refreshed_at = ------
          }
          # interface_stats_collection = ------
          # interface_stats_collection_failure_reason = ------
          # name = ------
          path_overlay_info {
            
            # control_plane = ------
            # data_packet_encapsulation = ------
            # dest_ip = ------
            # dest_port = ------
            # protocol = ------
            # source_ip = ------
            # source_port = ------
            vxlan_info {
              
              # dscp = ------
              # vnid = ------
            }
          }
          qos_statistics {
            
            # class_map_name = ------
            # drop_rate = ------
            # num_bytes = ------
            # num_packets = ------
            # offered_rate = ------
            # queue_bandwidthbps = ------
            # queue_depth = ------
            # queue_no_buffer_drops = ------
            # queue_total_drops = ------
            # refreshed_at = ------
          }
          # qos_stats_collection = ------
          # qos_stats_collection_failure_reason = ------
          # used_vlan = ------
          # vrf_name = ------
        }
        # ip = ------
        # link_information_source = ------
        # name = ------
        # perf_mon_collection = ------
        # perf_mon_collection_failure_reason = ------
        perf_mon_statistics {
          
          # byte_rate = ------
          # dest_ip_address = ------
          # dest_port = ------
          # input_interface = ------
          # ipv4_dsc_p = ------
          # ipv4_ttl = ------
          # output_interface = ------
          # packet_bytes = ------
          # packet_count = ------
          # packet_loss = ------
          # packet_loss_percentage = ------
          # protocol = ------
          # refreshed_at = ------
          # rtp_jitter_max = ------
          # rtp_jitter_mean = ------
          # rtp_jitter_min = ------
          # source_ip_address = ------
          # source_port = ------
        }
        # role = ------
        # ssid = ------
        # tunnels = [------]
        # type = ------
        # wlan_id = ------
      }
      network_elements_info {
        
        accuracy_list {
          
          # percent = ------
          # reason = ------
        }
        detailed_status {
          
          # acl_trace_calculation = ------
          # acl_trace_calculation_failure_reason = ------
        }
        device_statistics {
          
          cpu_statistics {
            
            # five_min_usage_in_percentage = ------
            # five_secs_usage_in_percentage = ------
            # one_min_usage_in_percentage = ------
            # refreshed_at = ------
          }
          memory_statistics {
            
            # memory_usage = ------
            # refreshed_at = ------
            # total_memory = ------
          }
        }
        # device_stats_collection = ------
        # device_stats_collection_failure_reason = ------
        egress_interface {
          
          physical_interface {
            
            acl_analysis {
              
              # acl_name = ------
              matching_aces {
                
                # ace = ------
                matching_ports {
                  
                  ports {
                    
                    # dest_ports = [------]
                    # source_ports = [------]
                  }
                  # protocol = ------
                }
                # result = ------
              }
              # result = ------
            }
            # id = ------
            interface_statistics {
              
              # admin_status = ------
              # input_packets = ------
              # input_queue_count = ------
              # input_queue_drops = ------
              # input_queue_flushes = ------
              # input_queue_max_depth = ------
              # input_ratebps = ------
              # operational_status = ------
              # output_drop = ------
              # output_packets = ------
              # output_queue_count = ------
              # output_queue_depth = ------
              # output_ratebps = ------
              # refreshed_at = ------
            }
            # interface_stats_collection = ------
            # interface_stats_collection_failure_reason = ------
            # name = ------
            path_overlay_info {
              
              # control_plane = ------
              # data_packet_encapsulation = ------
              # dest_ip = ------
              # dest_port = ------
              # protocol = ------
              # source_ip = ------
              # source_port = ------
              vxlan_info {
                
                # dscp = ------
                # vnid = ------
              }
            }
            qos_statistics {
              
              # class_map_name = ------
              # drop_rate = ------
              # num_bytes = ------
              # num_packets = ------
              # offered_rate = ------
              # queue_bandwidthbps = ------
              # queue_depth = ------
              # queue_no_buffer_drops = ------
              # queue_total_drops = ------
              # refreshed_at = ------
            }
            # qos_stats_collection = ------
            # qos_stats_collection_failure_reason = ------
            # used_vlan = ------
            # vrf_name = ------
          }
          virtual_interface {
            
            acl_analysis {
              
              # acl_name = ------
              matching_aces {
                
                # ace = ------
                matching_ports {
                  
                  ports {
                    
                    # dest_ports = [------]
                    # source_ports = [------]
                  }
                  # protocol = ------
                }
                # result = ------
              }
              # result = ------
            }
            # id = ------
            interface_statistics {
              
              # admin_status = ------
              # input_packets = ------
              # input_queue_count = ------
              # input_queue_drops = ------
              # input_queue_flushes = ------
              # input_queue_max_depth = ------
              # input_ratebps = ------
              # operational_status = ------
              # output_drop = ------
              # output_packets = ------
              # output_queue_count = ------
              # output_queue_depth = ------
              # output_ratebps = ------
              # refreshed_at = ------
            }
            # interface_stats_collection = ------
            # interface_stats_collection_failure_reason = ------
            # name = ------
            path_overlay_info {
              
              # control_plane = ------
              # data_packet_encapsulation = ------
              # dest_ip = ------
              # dest_port = ------
              # protocol = ------
              # source_ip = ------
              # source_port = ------
              vxlan_info {
                
                # dscp = ------
                # vnid = ------
              }
            }
            qos_statistics {
              
              # class_map_name = ------
              # drop_rate = ------
              # num_bytes = ------
              # num_packets = ------
              # offered_rate = ------
              # queue_bandwidthbps = ------
              # queue_depth = ------
              # queue_no_buffer_drops = ------
              # queue_total_drops = ------
              # refreshed_at = ------
            }
            # qos_stats_collection = ------
            # qos_stats_collection_failure_reason = ------
            # used_vlan = ------
            # vrf_name = ------
          }
        }
        flex_connect {
          
          # authentication = ------
          # data_switching = ------
          egress_acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          ingress_acl_analysis {
            
            # acl_name = ------
            matching_aces {
              
              # ace = ------
              matching_ports {
                
                ports {
                  
                  # dest_ports = [------]
                  # source_ports = [------]
                }
                # protocol = ------
              }
              # result = ------
            }
            # result = ------
          }
          # wireless_lan_controller_id = ------
          # wireless_lan_controller_name = ------
        }
        # id = ------
        ingress_interface {
          
          physical_interface {
            
            acl_analysis {
              
              # acl_name = ------
              matching_aces {
                
                # ace = ------
                matching_ports {
                  
                  ports {
                    
                    # dest_ports = [------]
                    # source_ports = [------]
                  }
                  # protocol = ------
                }
                # result = ------
              }
              # result = ------
            }
            # id = ------
            interface_statistics {
              
              # admin_status = ------
              # input_packets = ------
              # input_queue_count = ------
              # input_queue_drops = ------
              # input_queue_flushes = ------
              # input_queue_max_depth = ------
              # input_ratebps = ------
              # operational_status = ------
              # output_drop = ------
              # output_packets = ------
              # output_queue_count = ------
              # output_queue_depth = ------
              # output_ratebps = ------
              # refreshed_at = ------
            }
            # interface_stats_collection = ------
            # interface_stats_collection_failure_reason = ------
            # name = ------
            path_overlay_info {
              
              # control_plane = ------
              # data_packet_encapsulation = ------
              # dest_ip = ------
              # dest_port = ------
              # protocol = ------
              # source_ip = ------
              # source_port = ------
              vxlan_info {
                
                # dscp = ------
                # vnid = ------
              }
            }
            qos_statistics {
              
              # class_map_name = ------
              # drop_rate = ------
              # num_bytes = ------
              # num_packets = ------
              # offered_rate = ------
              # queue_bandwidthbps = ------
              # queue_depth = ------
              # queue_no_buffer_drops = ------
              # queue_total_drops = ------
              # refreshed_at = ------
            }
            # qos_stats_collection = ------
            # qos_stats_collection_failure_reason = ------
            # used_vlan = ------
            # vrf_name = ------
          }
          virtual_interface {
            
            acl_analysis {
              
              # acl_name = ------
              matching_aces {
                
                # ace = ------
                matching_ports {
                  
                  ports {
                    
                    # dest_ports = [------]
                    # source_ports = [------]
                  }
                  # protocol = ------
                }
                # result = ------
              }
              # result = ------
            }
            # id = ------
            interface_statistics {
              
              # admin_status = ------
              # input_packets = ------
              # input_queue_count = ------
              # input_queue_drops = ------
              # input_queue_flushes = ------
              # input_queue_max_depth = ------
              # input_ratebps = ------
              # operational_status = ------
              # output_drop = ------
              # output_packets = ------
              # output_queue_count = ------
              # output_queue_depth = ------
              # output_ratebps = ------
              # refreshed_at = ------
            }
            # interface_stats_collection = ------
            # interface_stats_collection_failure_reason = ------
            # name = ------
            path_overlay_info {
              
              # control_plane = ------
              # data_packet_encapsulation = ------
              # dest_ip = ------
              # dest_port = ------
              # protocol = ------
              # source_ip = ------
              # source_port = ------
              vxlan_info {
                
                # dscp = ------
                # vnid = ------
              }
            }
            qos_statistics {
              
              # class_map_name = ------
              # drop_rate = ------
              # num_bytes = ------
              # num_packets = ------
              # offered_rate = ------
              # queue_bandwidthbps = ------
              # queue_depth = ------
              # queue_no_buffer_drops = ------
              # queue_total_drops = ------
              # refreshed_at = ------
            }
            # qos_stats_collection = ------
            # qos_stats_collection_failure_reason = ------
            # used_vlan = ------
            # vrf_name = ------
          }
        }
        # ip = ------
        # link_information_source = ------
        # name = ------
        # perf_mon_collection = ------
        # perf_mon_collection_failure_reason = ------
        perf_monitor_statistics {
          
          # byte_rate = ------
          # dest_ip_address = ------
          # dest_port = ------
          # input_interface = ------
          # ipv4_dsc_p = ------
          # ipv4_ttl = ------
          # output_interface = ------
          # packet_bytes = ------
          # packet_count = ------
          # packet_loss = ------
          # packet_loss_percentage = ------
          # protocol = ------
          # refreshed_at = ------
          # rtp_jitter_max = ------
          # rtp_jitter_mean = ------
          # rtp_jitter_min = ------
          # source_ip_address = ------
          # source_port = ------
        }
        # role = ------
        # ssid = ------
        # tunnels = [------]
        # type = ------
        # wlan_id = ------
      }
      # properties = [------]
      request {
        
        # control_path = ------
        # create_time = ------
        # dest_ip = ------
        # dest_port = ------
        # failure_reason = ------
        # id = ------
        # inclusions = [------]
        # last_update_time = ------
        # periodic_refresh = ------
        # protocol = ------
        # source_ip = ------
        # source_port = ------
        # status = ------
      }
    }
    parameters {
      
      control_path = "false"
      dest_ip = "string"
      dest_port = "string"
      flow_analysis_id = "string"
      inclusions = ["string"]
      periodic_refresh = "false"
      protocol = "string"
      source_ip = "string"
      source_port = "string"
    }
}

output "dnac_path_trace_example" {
    value = dnac_path_trace.example
}