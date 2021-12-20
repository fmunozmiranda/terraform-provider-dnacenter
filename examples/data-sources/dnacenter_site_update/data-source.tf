
data "dnacdnacenter_site_update" "example" {
  provider = dnacenter
  site_id  = "string"
  area {

    name        = "string"
    parent_name = "string"
  }
  building {

    address     = "string"
    latitude    = 1.0
    longitude   = 1.0
    name        = "string"
    parent_name = "string"
  }
  floor {

    height   = 1.0
    length   = 1.0
    name     = "string"
    rf_model = "string"
    width    = 1.0
  }
}