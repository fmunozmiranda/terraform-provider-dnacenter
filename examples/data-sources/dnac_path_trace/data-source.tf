
data "dnac_path_trace" "example" {
    provider = dnac
    dest_ip = "string"
    dest_port = "string"
    gt_create_time = "string"
    last_update_time = "string"
    limit = "string"
    lt_create_time = "string"
    offset = "string"
    order = "string"
    periodic_refresh = "false"
    protocol = "string"
    sort_by = "string"
    source_ip = "string"
    source_port = "string"
    status = "string"
    task_id = "string"
}

output "dnac_path_trace_example" {
    value = data.dnac_path_trace.example.items
}

data "dnac_path_trace" "example" {
    provider = dnac
    flow_analysis_id = "string"
}

output "dnac_path_trace_example" {
    value = data.dnac_path_trace.example.item
}
