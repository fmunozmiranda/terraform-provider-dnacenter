
data "dnac_task" "example" {
    provider = dnac
    data = "string"
    end_time = "string"
    error_code = "string"
    failure_reason = "string"
    is_error = "string"
    limit = "string"
    offset = "string"
    order = "string"
    parent_id = "string"
    progress = "string"
    service_type = "string"
    sort_by = "string"
    start_time = "string"
    username = "string"
}

output "dnac_task_example" {
    value = data.dnac_task.example.items
}

data "dnac_task" "example" {
    provider = dnac
    task_id = "string"
}

output "dnac_task_example" {
    value = data.dnac_task.example.item
}
