
data "dnacenter_issues" "example" {
    provider = dnac
    ai_driven = "string"
    device_id = "string"
    end_time = ------
    issue_status = "string"
    mac_address = "string"
    priority = "string"
    site_id = "string"
    start_time = ------
}

output "dnacenter_issues_example" {
    value = data.dnac_issues.example.items
}
