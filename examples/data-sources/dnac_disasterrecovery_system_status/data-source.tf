
data "dnac_disasterrecovery_system_status" "example" {
    provider = dnac
}

output "dnac_disasterrecovery_system_status_example" {
    value = data.dnac_disasterrecovery_system_status.example.item
}
