
data "dnac_applications" "example" {
    provider = dnac
    limit = ------
    name = "string"
    offset = ------
}

output "dnac_applications_example" {
    value = data.dnac_applications.example.items
}
