
data "dnacenter_tag_count" "example" {
    provider = dnac
    attribute_name = "string"
    level = "string"
    name = "string"
    name_space = "string"
    size = "string"
    system_tag = "string"
}

output "dnacenter_tag_count_example" {
    value = data.dnac_tag_count.example.item
}
