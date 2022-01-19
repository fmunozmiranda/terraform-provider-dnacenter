
resource "dnac_qos_device_interface" "example" {
    provider = dnac
    parameters {
      
      excluded_interfaces = ["string"]
      id = "string"
      name = "string"
      network_device_id = "string"
      qos_device_interface_info {
        
        dmvpn_remote_sites_bw = [1]
        instance_id = 1
        interface_id = "string"
        interface_name = "string"
        label = "string"
        role = "string"
        upload_bw = 1
      }
    }
}

output "dnac_qos_device_interface_example" {
    value = dnac_qos_device_interface.example
}