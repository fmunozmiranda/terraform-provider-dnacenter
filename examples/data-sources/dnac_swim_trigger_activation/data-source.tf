
data "dnacswim_trigger_activation" "example" {
    provider = dnac
    schedule_validate = "false"
    activate_lower_image_version = "false"
    device_upgrade_mode = "string"
    device_uuid = "string"
    distribute_if_needed = "false"
    image_uuid_list = ["string"]
    item {
      
      # task_id = ------
      # url = ------
    }
    smu_image_uuid_list = ["string"]
}