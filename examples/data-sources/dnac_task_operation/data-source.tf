
data "dnac_task_operation" "example" {
    provider = dnac
    limit = 1
    offset = 1
    operation_id = "string"
}

output "dnac_task_operation_example" {
    value = data.dnac_task_operation.example.items
}
