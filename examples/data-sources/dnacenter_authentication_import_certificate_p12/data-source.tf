
data "dnacdnacenter_authentication_import_certificate_p12" "example" {
  provider      = dnacenter
  list_of_users = ["string"]
  p12_password  = "******"
  pk_password   = "******"
}