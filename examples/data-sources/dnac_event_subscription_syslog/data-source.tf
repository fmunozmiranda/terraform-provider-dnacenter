
data "dnac_event_subscription_syslog" "example" {
    provider = dnac
    event_ids = "string"
    limit = ------
    offset = ------
    order = "string"
    sort_by = "string"
}

output "dnac_event_subscription_syslog_example" {
    value = data.dnac_event_subscription_syslog.example.items
}
