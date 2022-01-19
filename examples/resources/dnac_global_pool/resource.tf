
resource "dnac_global_pool" "example" {
    provider = dnac
    parameters {
      
      id = "string"
      settings {
        
        ippool {
          
          ip_address_space = "string"
          dhcp_server_ips = ["string"]
          dns_server_ips = ["string"]
          gateway = "string"
          id = "string"
          ip_pool_cidr = "string"
          ip_pool_name = "string"
          type = "string"
        }
      }
    }
}

output "dnac_global_pool_example" {
    value = dnac_global_pool.example
}