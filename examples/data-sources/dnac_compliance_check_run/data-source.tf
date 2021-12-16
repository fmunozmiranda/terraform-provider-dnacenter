
data "dnaccompliance_check_run" "example" {
    provider = dnac
    categories = ["string"]
    device_uuids = ["string"]
    item {
      
      # task_id = ------
      # url = ------
    }
    trigger_full = "false"
}