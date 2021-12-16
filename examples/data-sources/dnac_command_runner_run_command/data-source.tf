
data "dnacenter_command_runner_run_command" "example" {
    provider = dnac
    commands = ["string"]
    description = "string"
    device_uuids = ["string"]
    item {
      
      # task_id = ------
      # url = ------
    }
    name = "string"
    timeout = 1
}