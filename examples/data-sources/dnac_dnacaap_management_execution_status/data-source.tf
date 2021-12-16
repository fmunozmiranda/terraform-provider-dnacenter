
data "dnac_dnacaap_management_execution_status" "example" {
    provider = dnac
    execution_id = "string"
}

output "dnac_dnacaap_management_execution_status_example" {
    value = data.dnac_dnacaap_management_execution_status.example.item
}
