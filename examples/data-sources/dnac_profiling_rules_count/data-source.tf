
data "dnac_profiling_rules_count" "example" {
    provider = dnac
    include_deleted = "false"
    rule_type = "string"
}

output "dnac_profiling_rules_count_example" {
    value = data.dnac_profiling_rules_count.example.item
}
