
data "dnacenter_platform_release_summary" "example" {
    provider = dnac
}

output "dnacenter_platform_release_summary_example" {
    value = data.dnac_platform_release_summary.example.item
}
