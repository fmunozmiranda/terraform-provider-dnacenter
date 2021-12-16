
data "dnac_application_sets_count" "example" {
    provider = dnac
}

output "dnac_application_sets_count_example" {
    value = data.dnac_application_sets_count.example.item
}
