
data "dnacenter_event_api_status" "example" {
    provider = dnac
    execution_id = "string"
}

output "dnacenter_event_api_status_example" {
    value = data.dnac_event_api_status.example.item
}
