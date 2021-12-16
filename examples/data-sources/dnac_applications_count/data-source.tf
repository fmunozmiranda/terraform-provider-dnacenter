
data "dnac_applications_count" "example" {
    provider = dnac
}

output "dnac_applications_count_example" {
    value = data.dnac_applications_count.example.item
}
