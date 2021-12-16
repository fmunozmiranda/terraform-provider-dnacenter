
data "dnac_sda_count" "example" {
    provider = dnac
}

output "dnac_sda_count_example" {
    value = data.dnac_sda_count.example.item
}
