
data "dnac_application_sets" "example" {
    provider = dnac
    limit = ------
    name = "string"
    offset = ------
}

output "dnac_application_sets_example" {
    value = data.dnac_application_sets.example.items
}