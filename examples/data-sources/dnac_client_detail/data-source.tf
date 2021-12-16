
data "dnac_client_detail" "example" {
    provider = dnac
    mac_address = "string"
    timestamp = "string"
}

output "dnac_client_detail_example" {
    value = data.dnac_client_detail.example.item
}
