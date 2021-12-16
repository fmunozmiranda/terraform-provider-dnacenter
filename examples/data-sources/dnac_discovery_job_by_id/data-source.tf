
data "dnac_discovery_job_by_id" "example" {
    provider = dnac
    id = "string"
    ip_address = "string"
    limit = 1
    offset = 1
}

output "dnac_discovery_job_by_id_example" {
    value = data.dnac_discovery_job_by_id.example.items
}
