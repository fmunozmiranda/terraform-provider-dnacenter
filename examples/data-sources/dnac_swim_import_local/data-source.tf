
data "dnacenter_swim_import_local" "example" {
    provider = dnac
    is_third_party = "false"
    third_party_application_type = "string"
    third_party_image_family = "string"
    third_party_vendor = "string"
    item {
      
      # task_id = ------
      # url = ------
    }
}