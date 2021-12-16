
data "dnac_app_policy_queuing_profile" "example" {
    provider = dnac
    name = "string"
}

output "dnac_app_policy_queuing_profile_example" {
    value = data.dnac_app_policy_queuing_profile.example.items
}
