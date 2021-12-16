
data "dnacenter_nfv_provision_detail" "example" {
    provider = dnac
    device_ip = "string"
}

output "dnacenter_nfv_provision_detail_example" {
    value = data.dnac_nfv_provision_detail.example.item
}
