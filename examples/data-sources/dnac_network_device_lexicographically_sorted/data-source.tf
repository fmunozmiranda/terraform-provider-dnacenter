
data "dnac_network_device_lexicographically_sorted" "example" {
    provider = dnac
    associated_wlc_ip = "string"
    collection_interval = "string"
    collection_status = "string"
    error_code = "string"
    family = "string"
    hostname = "string"
    limit = ------
    mac_address = "string"
    management_ip_address = "string"
    offset = ------
    platform_id = "string"
    reachability_failure_reason = "string"
    reachability_status = "string"
    role = "string"
    role_source = "string"
    serial_number = "string"
    series = "string"
    software_type = "string"
    software_version = "string"
    type = "string"
    up_time = "string"
    vrf_name = "string"
}

output "dnac_network_device_lexicographically_sorted_example" {
    value = data.dnac_network_device_lexicographically_sorted.example.items
}