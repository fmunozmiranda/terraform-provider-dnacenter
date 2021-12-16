
data "dnac_discovery_summary" "example" {
    provider = dnac
    cli_status = ["string"]
    http_status = ["string"]
    id = "string"
    ip_address = ["string"]
    netconf_status = ["string"]
    ping_status = ["string"]
    snmp_status = ["string"]
    sort_by = "string"
    sort_order = "string"
    task_id = "string"
}

output "dnac_discovery_summary_example" {
    value = data.dnac_discovery_summary.example.item
}
