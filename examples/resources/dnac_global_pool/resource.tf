
resource "dnacenter_global_pool" "example" {
  provider = dnacenter
  parameters {

    id = "f736f1ba-ea3e-4ac5-94db-e6577eebf69b"
    settings {

      ippool {

        ip_address_space = ""
        dhcp_server_ips  = []
        dns_server_ips   = ["101.101.101.101"]
        gateway          = ""
        id               = "f736f1ba-ea3e-4ac5-94db-e6577eebf69b"
        ip_pool_cidr     = "12.0.0.0/8"
        ip_pool_name     = "12Network"
        type             = "generic"
      }
    }
  }
}

output "dnacenter_global_pool_example" {
  value = dnacenter_global_pool.example
}