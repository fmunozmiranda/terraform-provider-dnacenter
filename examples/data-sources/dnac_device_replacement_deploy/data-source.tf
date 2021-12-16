
data "dnacenter_device_replacement_deploy" "example" {
    provider = dnac
    faulty_device_serial_number = "string"
    item {
      
      # task_id = ------
      # url = ------
    }
    replacement_device_serial_number = "string"
}