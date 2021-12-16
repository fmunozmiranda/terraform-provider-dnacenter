
data "dnacenter_system_performance_historical" "example" {
    provider = dnac
    end_time = ------
    kpi = "string"
    start_time = ------
}

output "dnacenter_system_performance_historical_example" {
    value = data.dnac_system_performance_historical.example.item
}
