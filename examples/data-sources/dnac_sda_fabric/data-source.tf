
data "dnacenter_sda_fabric" "example" {
    provider = dnac
    fabric_name = "string"
}

output "dnacenter_sda_fabric_example" {
    value = data.dnac_sda_fabric.example.item
}
