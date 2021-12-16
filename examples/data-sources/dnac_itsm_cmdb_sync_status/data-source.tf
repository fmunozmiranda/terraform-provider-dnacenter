
data "dnac_itsm_cmdb_sync_status" "example" {
    provider = dnac
    date = "string"
    status = "string"
}

output "dnac_itsm_cmdb_sync_status_example" {
    value = data.dnac_itsm_cmdb_sync_status.example.items
}
