
data "dnacenter_authentication_import_certificate_p12" "example" {
    provider = dnac
    list_of_users = ["string"]
    p12_password = "******"
    pk_password = "******"
    item {
      
      # task_id = ------
      # url = ------
    }
}