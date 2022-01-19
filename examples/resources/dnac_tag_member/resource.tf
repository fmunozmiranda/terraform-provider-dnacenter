
resource "dnac_tag_member" "example" {
    provider = dnac
    parameters {
      
      id = "string"
      member_id = "string"
    }
}

output "dnac_tag_member_example" {
    value = dnac_tag_member.example
}