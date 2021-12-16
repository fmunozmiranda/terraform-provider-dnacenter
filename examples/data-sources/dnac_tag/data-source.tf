
data "dnacenter_tag" "example" {
    provider = dnac
    additional_info_attributes = "string"
    additional_info_name_space = "string"
    field = "string"
    level = "string"
    limit = "string"
    name = "string"
    offset = "string"
    order = "string"
    size = "string"
    sort_by = "string"
    system_tag = "string"
}

output "dnacenter_tag_example" {
    value = data.dnac_tag.example.items
}

data "dnacenter_tag" "example" {
    provider = dnac
    id = "string"
}

output "dnacenter_tag_example" {
    value = data.dnac_tag.example.item
}
