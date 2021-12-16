
data "dnac_event_artifact" "example" {
    provider = dnac
    event_ids = "string"
    limit = ------
    offset = ------
    order = "string"
    search = "string"
    sort_by = "string"
    tags = "string"
}

output "dnac_event_artifact_example" {
    value = data.dnac_event_artifact.example.items
}
