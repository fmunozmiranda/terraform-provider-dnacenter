
data "dnac_reports" "example" {
    provider = dnac
    view_group_id = "string"
    view_id = "string"
}

output "dnac_reports_example" {
    value = data.dnac_reports.example.items
}

data "dnac_reports" "example" {
    provider = dnac
    report_id = "string"
}

output "dnac_reports_example" {
    value = data.dnac_reports.example.item
}
