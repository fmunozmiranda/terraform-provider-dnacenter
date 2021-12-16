
data "dnacenter_discovery_range_delete" "example" {
    provider = dnac
    records_to_delete = 1
    start_index = 1
    item {
      
      # task_id = ------
      # url = ------
    }
}