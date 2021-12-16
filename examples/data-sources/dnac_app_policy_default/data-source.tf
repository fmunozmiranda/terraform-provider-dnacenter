
data "dnacenter_app_policy_default" "example" {
    provider = dnac
}

output "dnacenter_app_policy_default_example" {
    value = data.dnac_app_policy_default.example.items
}
