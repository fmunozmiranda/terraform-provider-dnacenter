
data "dnac_event_artifact_count" "example" {
    provider = dnac
}

output "dnac_event_artifact_count_example" {
    value = data.dnac_event_artifact_count.example.item
}
