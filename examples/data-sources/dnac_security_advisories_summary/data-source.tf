
data "dnacenter_security_advisories_summary" "example" {
    provider = dnac
}

output "dnacenter_security_advisories_summary_example" {
    value = data.dnac_security_advisories_summary.example.item
}
