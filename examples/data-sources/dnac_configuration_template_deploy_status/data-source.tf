
data "dnacenter_configuration_template_deploy_status" "example" {
    provider = dnac
    deployment_id = "string"
}

output "dnacenter_configuration_template_deploy_status_example" {
    value = data.dnac_configuration_template_deploy_status.example.item
}
