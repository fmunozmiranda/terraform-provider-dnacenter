
data "dnacenter_app_policy_queuing_profile_count" "example" {
    provider = dnac
}

output "dnacenter_app_policy_queuing_profile_count_example" {
    value = data.dnac_app_policy_queuing_profile_count.example.item
}
