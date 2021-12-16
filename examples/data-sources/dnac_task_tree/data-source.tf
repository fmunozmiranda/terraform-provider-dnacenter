
data "dnacenter_task_tree" "example" {
    provider = dnac
    task_id = "string"
}

output "dnacenter_task_tree_example" {
    value = data.dnac_task_tree.example.items
}
