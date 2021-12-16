
data "dnacenter_platform_nodes_configuration_summary" "example" {
    provider = dnac
}

output "dnacenter_platform_nodes_configuration_summary_example" {
    value = data.dnac_platform_nodes_configuration_summary.example.item
}
