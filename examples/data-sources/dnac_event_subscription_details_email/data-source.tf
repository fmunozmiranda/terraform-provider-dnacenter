
data "dnac_event_subscription_details_email" "example" {
    provider = dnac
    connector_type = "string"
    instance_id = "string"
    name = "string"
}

output "dnac_event_subscription_details_email_example" {
    value = data.dnac_event_subscription_details_email.example.items
}
