
data "dnacwireless_provision_device_create" "example" {
    provider = dnac
    device_name = "string"
    dynamic_interfaces {
      
      interface_gateway = "string"
      interface_ipaddress = "string"
      interface_name = "string"
      interface_netmask_in_cid_r = 1
      lag_or_port_number = 1
      vlan_id = 1
    }
    item {
      
      # execution_id = ------
      # execution_url = ------
      provisioning_tasks {
        
        # failed = [------]
        # success = [------]
      }
    }
    managed_aplocations = ["string"]
    site = "string"
}