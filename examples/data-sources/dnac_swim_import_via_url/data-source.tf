
data "dnacswim_import_via_url" "example" {
    provider = dnac
    schedule_at = "string"
    schedule_desc = "string"
    schedule_origin = "string"
    application_type = "string"
    image_family = "string"
    item {
      
      # task_id = ------
      # url = ------
    }
    source_url = "string"
    third_party = "false"
    vendor = "string"
}