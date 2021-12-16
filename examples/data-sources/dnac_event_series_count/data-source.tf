
data "dnac_event_series_count" "example" {
    provider = dnac
    category = "string"
    domain = "string"
    end_time = ------
    event_ids = "string"
    severity = "string"
    source = "string"
    start_time = ------
    sub_domain = "string"
    type = "string"
}

output "dnac_event_series_count_example" {
    value = data.dnac_event_series_count.example.item
}
