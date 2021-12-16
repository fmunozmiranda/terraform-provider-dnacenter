
data "dnacenter_event_subscription" "example" {
    provider = dnac
    event_ids = "string"
    limit = ------
    offset = ------
    order = "string"
    sort_by = "string"
}

output "dnacenter_event_subscription_example" {
    value = data.dnac_event_subscription.example.items
}
