
data "dnacenter_endpoint_analytics_profiling_rules" "example" {
    provider = dnac
    include_deleted = "false"
    limit = ------
    offset = ------
    order = "string"
    rule_type = "string"
    sort_by = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
    value = data.dnac_endpoint_analytics_profiling_rules.example.items
}

data "dnacenter_endpoint_analytics_profiling_rules" "example" {
    provider = dnac
    rule_id = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
    value = data.dnac_endpoint_analytics_profiling_rules.example.item
}
