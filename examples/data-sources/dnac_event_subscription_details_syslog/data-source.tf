
data "dnac_event_subscription_details_syslog" "example" {
    provider = dnac
    connector_type = "string"
    instance_id = "string"
    name = "string"
}

output "dnac_event_subscription_details_syslog_example" {
    value = data.dnac_event_subscription_details_syslog.example.items
}
