
data "dnac_itsm_integration_events_failed" "example" {
    provider = dnac
    instance_id = "string"
}

output "dnac_itsm_integration_events_failed_example" {
    value = data.dnac_itsm_integration_events_failed.example.items
}
