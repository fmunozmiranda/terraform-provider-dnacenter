
data "dnacenter_security_advisories" "example" {
    provider = dnac
}

output "dnacenter_security_advisories_example" {
    value = data.dnac_security_advisories.example.items
}
