
data "dnac_app_policy_default" "example" {
    provider = dnac
}

output "dnac_app_policy_default_example" {
    value = data.dnac_app_policy_default.example.items
}
