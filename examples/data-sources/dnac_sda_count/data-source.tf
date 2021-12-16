
data "dnacenter_sda_count" "example" {
    provider = dnac
}

output "dnacenter_sda_count_example" {
    value = data.dnac_sda_count.example.item
}
