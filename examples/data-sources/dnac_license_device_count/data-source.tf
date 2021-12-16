
data "dnac_license_device_count" "example" {
    provider = dnac
    device_type = "string"
    dna_level = "string"
    registration_status = "string"
    smart_account_id = "string"
    virtual_account_name = "string"
}

output "dnac_license_device_count_example" {
    value = data.dnac_license_device_count.example.item
}
