
data "dnacenter_device_credential" "example" {
    provider = dnac
    site_id = "string"
}

output "dnacenter_device_credential_example" {
    value = data.dnac_device_credential.example.item
}
