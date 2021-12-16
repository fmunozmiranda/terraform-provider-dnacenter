
data "dnac_event_series_audit_logs_summary" "example" {
    provider = dnac
    category = "string"
    context = "string"
    description = "string"
    device_id = "string"
    domain = "string"
    end_time = ------
    event_hierarchy = "string"
    event_id = "string"
    instance_id = "string"
    is_parent_only = "false"
    is_system_events = "false"
    name = "string"
    parent_instance_id = "string"
    severity = "string"
    site_id = "string"
    source = "string"
    start_time = ------
    sub_domain = "string"
    user_id = "string"
}

output "dnac_event_series_audit_logs_summary_example" {
    value = data.dnac_event_series_audit_logs_summary.example.items
}