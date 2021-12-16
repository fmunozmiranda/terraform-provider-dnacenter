
data "dnac_wireless_sensor_test_results" "example" {
    provider = dnac
    end_time = ------
    site_id = "string"
    start_time = ------
    test_failure_by = "string"
}

output "dnac_wireless_sensor_test_results_example" {
    value = data.dnac_wireless_sensor_test_results.example.item
}
