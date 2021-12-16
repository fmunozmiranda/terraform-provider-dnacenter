
data "dnacenter_event_count" "example" {
    provider = dnac
    event_id = "string"
    tags = "string"
}

output "dnacenter_event_count_example" {
    value = data.dnac_event_count.example.item
}
