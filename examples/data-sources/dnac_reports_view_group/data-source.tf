
data "dnac_reports_view_group" "example" {
    provider = dnac
}

output "dnac_reports_view_group_example" {
    value = data.dnac_reports_view_group.example.items
}

data "dnac_reports_view_group" "example" {
    provider = dnac
    view_group_id = "string"
}

output "dnac_reports_view_group_example" {
    value = data.dnac_reports_view_group.example.item
}
