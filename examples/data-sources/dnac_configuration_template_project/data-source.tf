
data "dnacenter_configuration_template_project" "example" {
    provider = dnac
    name = "string"
    sort_order = "string"
}

output "dnacenter_configuration_template_project_example" {
    value = data.dnac_configuration_template_project.example.items
}

data "dnacenter_configuration_template_project" "example" {
    provider = dnac
    project_id = "string"
}

output "dnacenter_configuration_template_project_example" {
    value = data.dnac_configuration_template_project.example.item
}
