
data "dnacenter_applications" "example" {
    provider = dnac
    limit = ------
    name = "string"
    offset = ------
}

output "dnacenter_applications_example" {
    value = data.dnac_applications.example.items
}
