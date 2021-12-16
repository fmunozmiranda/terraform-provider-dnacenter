
data "dnacenter_event_series_audit_logs_parent_records" "example" {
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
    is_system_events = "false"
    limit = ------
    name = "string"
    offset = ------
    order = "string"
    severity = "string"
    site_id = "string"
    sort_by = "string"
    source = "string"
    start_time = ------
    sub_domain = "string"
    user_id = "string"
}

output "dnacenter_event_series_audit_logs_parent_records_example" {
    value = data.dnac_event_series_audit_logs_parent_records.example.items
}
