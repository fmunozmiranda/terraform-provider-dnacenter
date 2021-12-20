
data "dnacdnacenter_authentication_import_certificate" "example" {
  provider      = dnacenter
  list_of_users = ["string"]
  pk_password   = "******"
}