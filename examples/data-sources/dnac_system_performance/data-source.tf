
data "dnac_system_performance" "example" {
    provider = dnac
    end_time = ------
    function = "string"
    kpi = "string"
    start_time = ------
}

output "dnac_system_performance_example" {
    value = data.dnac_system_performance.example.item
}
