
data "dnac_configuration_template_version" "example" {
    provider = dnac
    template_id = "string"
}

output "dnac_configuration_template_version_example" {
    value = data.dnac_configuration_template_version.example.items
}
