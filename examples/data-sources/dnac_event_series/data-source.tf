
data "dnac_event_series" "example" {
    provider = dnac
    category = "string"
    domain = "string"
    end_time = ------
    event_ids = "string"
    limit = ------
    offset = ------
    order = "string"
    severity = "string"
    sort_by = "string"
    source = "string"
    start_time = ------
    sub_domain = "string"
    type = "string"
}

output "dnac_event_series_example" {
    value = data.dnac_event_series.example.items
}
