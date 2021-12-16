
data "dnacenter_tag_member_count" "example" {
    provider = dnac
    id = "string"
    level = "string"
    member_association_type = "string"
    member_type = "string"
}

output "dnacenter_tag_member_count_example" {
    value = data.dnac_tag_member_count.example.item
}
