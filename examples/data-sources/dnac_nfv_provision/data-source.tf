
data "dnacnfv_provision" "example" {
    provider = dnac
    device {
      
      custom_networks {
        
        connection_type = "string"
        name = "string"
        network_mode = "string"
        services_to_connect {
          
          service = "string"
        }
        vlan = "string"
      }
      custom_services {
        
        application_type = "string"
        image_name = "string"
        name = "string"
        profile = "string"
        topology {
          
          assign_ip = "string"
          name = "string"
          type = "string"
        }
      }
      custom_template {
        
        device_type = "string"
        template = "string"
      }
      device_type = "string"
      dia = "false"
      service_providers {
        
        connect = "false"
        default_gateway = "false"
        link_type = "string"
        service_provider = "string"
      }
      services {
        
        image_name = "string"
        mode = "string"
        name = "string"
        profile = "string"
        topology {
          
          assign_ip = "string"
          name = "string"
          type = "string"
        }
        type = "string"
      }
      tag_name = "string"
      vlan {
        
        id = "string"
        type = "string"
      }
    }
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    site_profile_name = "string"
}