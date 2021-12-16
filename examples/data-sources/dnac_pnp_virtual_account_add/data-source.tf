
data "dnacpnp_virtual_account_add" "example" {
    provider = dnac
    address_fqdn = "string"
    address_ip_v4 = "string"
    cert = "string"
    item {
      
      # auto_sync_period = ------
      # cco_user = ------
      # expiry = ------
      # last_sync = ------
      profile {
        
        # address_fqdn = ------
        # address_ip_v4 = ------
        # cert = ------
        # make_default = ------
        # name = ------
        # port = ------
        # profile_id = ------
        # proxy = ------
      }
      # smart_account_id = ------
      sync_result {
        
        sync_list {
          
          # device_sn_list = [------]
          # sync_type = ------
        }
        # sync_msg = ------
      }
      # sync_result_str = ------
      # sync_start_time = ------
      # sync_status = ------
      # tenant_id = ------
      # token = ------
      # virtual_account_id = ------
    }
    make_default = "false"
    name = "string"
    port = 1
    profile_id = "string"
    proxy = "false"
}