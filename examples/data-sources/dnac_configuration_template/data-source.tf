
data "dnac_configuration_template" "example" {
    provider = dnac
    filter_conflicting_templates = "false"
    product_family = "string"
    product_series = "string"
    product_type = "string"
    project_id = "string"
    project_names = ["string"]
    software_type = "string"
    software_version = "string"
    sort_order = "string"
    tags = ["string"]
    un_committed = "false"
}

output "dnac_configuration_template_example" {
    value = data.dnac_configuration_template.example.items
}

data "dnac_configuration_template" "example" {
    provider = dnac
    latest_version = "false"
    template_id = "string"
}

output "dnac_configuration_template_example" {
    value = data.dnac_configuration_template.example.item
}
