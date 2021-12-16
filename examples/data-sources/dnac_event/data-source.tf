
data "dnacenter_event" "example" {
    provider = dnac
    event_id = "string"
    limit = ------
    offset = ------
    order = "string"
    sort_by = "string"
    tags = "string"
}

output "dnacenter_event_example" {
    value = data.dnac_event.example.items
}
