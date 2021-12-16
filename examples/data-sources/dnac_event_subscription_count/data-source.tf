
data "dnac_event_subscription_count" "example" {
    provider = dnac
    event_ids = "string"
}

output "dnac_event_subscription_count_example" {
    value = data.dnac_event_subscription_count.example.item
}
