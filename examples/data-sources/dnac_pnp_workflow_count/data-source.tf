
data "dnacenter_pnp_workflow_count" "example" {
    provider = dnac
    name = ["string"]
}

output "dnacenter_pnp_workflow_count_example" {
    value = data.dnac_pnp_workflow_count.example.item
}
