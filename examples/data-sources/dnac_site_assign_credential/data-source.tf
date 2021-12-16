
data "dnacenter_site_assign_credential" "example" {
    provider = dnac
    site_id = "string"
    cli_id = "string"
    http_read = "string"
    http_write = "string"
    item {
      
      # execution_id = ------
      # execution_status_url = ------
      # message = ------
    }
    snmp_v2_read_id = "string"
    snmp_v2_write_id = "string"
    snmp_v3_id = "string"
}