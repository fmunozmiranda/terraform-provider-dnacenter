
data "dnac_reports_executions" "example" {
    provider = dnac
    report_id = "string"
}

output "dnac_reports_executions_example" {
    value = data.dnac_reports_executions.example.item
}