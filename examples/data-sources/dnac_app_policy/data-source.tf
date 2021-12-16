
data "dnacenter_app_policy" "example" {
    provider = dnac
    policy_scope = "string"
}

output "dnacenter_app_policy_example" {
    value = data.dnac_app_policy.example.items
}
