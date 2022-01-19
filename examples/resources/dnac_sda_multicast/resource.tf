
resource "dnac_sda_multicast" "example" {
    provider = dnac
    parameters {
      
      multicast_method = "string"
      multicast_vn_info {
        
        external_rp_ip_address = "string"
        ip_pool_name = "string"
        ssm_group_range = "string"
        ssm_info = ["string"]
        ssm_wildcard_mask = "string"
        virtual_network_name = "string"
      }
      muticast_type = "string"
      site_name_hierarchy = "string"
    }
}

output "dnac_sda_multicast_example" {
    value = dnac_sda_multicast.example
}