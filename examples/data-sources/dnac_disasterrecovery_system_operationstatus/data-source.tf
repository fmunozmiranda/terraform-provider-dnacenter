
data "dnac_disasterrecovery_system_operationstatus" "example" {
    provider = dnac
}

output "dnac_disasterrecovery_system_operationstatus_example" {
    value = data.dnac_disasterrecovery_system_operationstatus.example.item
}
