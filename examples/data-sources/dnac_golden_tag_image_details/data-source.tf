
data "dnac_golden_tag_image_details" "example" {
    provider = dnac
    device_family_identifier = "string"
    device_role = "string"
    image_id = "string"
    site_id = "string"
}

output "dnac_golden_tag_image_details_example" {
    value = data.dnac_golden_tag_image_details.example.item
}
