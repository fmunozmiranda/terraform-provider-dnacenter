
data "dnacdevice_credential_create" "example" {
    provider = dnac
    cli_credential {
      
      description = "string"
      enable_password = "string"
      password = "******"
      username = "string"
    }
    https_read {
      
      name = "string"
      password = "******"
      port = ------
      username = "string"
    }
    https_write {
      
      name = "string"
      password = "******"
      port = ------
      username = "string"
    }
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    snmp_v2c_read {
      
      description = "string"
      read_community = "string"
    }
    snmp_v2c_write {
      
      description = "string"
      write_community = "string"
    }
    snmp_v3 {
      
      auth_password = "string"
      auth_type = "string"
      description = "string"
      privacy_password = "string"
      privacy_type = "string"
      snmp_mode = "string"
      username = "string"
    }
}