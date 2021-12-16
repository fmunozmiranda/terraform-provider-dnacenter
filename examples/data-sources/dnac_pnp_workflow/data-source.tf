
data "dnacenter_pnp_workflow" "example" {
    provider = dnac
    limit = 1
    name = ["string"]
    offset = 1
    sort = ["string"]
    sort_order = "string"
    type = ["string"]
}

output "dnacenter_pnp_workflow_example" {
    value = data.dnac_pnp_workflow.example.items
}

data "dnacenter_pnp_workflow" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_pnp_workflow_example" {
    value = data.dnac_pnp_workflow.example.item
}
