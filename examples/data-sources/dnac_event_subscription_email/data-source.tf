
data "dnac_event_subscription_email" "example" {
    provider = dnac
    event_ids = "string"
    limit = ------
    offset = ------
    order = "string"
    sort_by = "string"
}

output "dnac_event_subscription_email_example" {
    value = data.dnac_event_subscription_email.example.items
}
