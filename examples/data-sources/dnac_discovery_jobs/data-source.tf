
data "dnac_discovery_jobs" "example" {
    provider = dnac
    ip_address = "string"
    limit = 1
    name = "string"
    offset = 1
}

output "dnac_discovery_jobs_example" {
    value = data.dnac_discovery_jobs.example.items
}
