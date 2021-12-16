
data "dnacenter_reports_view_group_view" "example" {
    provider = dnac
    view_group_id = "string"
    view_id = "string"
}

output "dnacenter_reports_view_group_view_example" {
    value = data.dnac_reports_view_group_view.example.item
}
