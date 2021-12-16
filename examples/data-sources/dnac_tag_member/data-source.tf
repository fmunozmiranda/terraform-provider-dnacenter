
data "dnacenter_tag_member" "example" {
    provider = dnac
    id = "string"
    level = "string"
    limit = "string"
    member_association_type = "string"
    member_type = "string"
    offset = "string"
}

output "dnacenter_tag_member_example" {
    value = data.dnac_tag_member.example.items
}
