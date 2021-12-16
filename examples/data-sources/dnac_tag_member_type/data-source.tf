
data "dnac_tag_member_type" "example" {
    provider = dnac
}

output "dnac_tag_member_type_example" {
    value = data.dnac_tag_member_type.example.items
}
