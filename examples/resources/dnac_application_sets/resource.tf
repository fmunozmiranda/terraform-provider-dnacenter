
resource "dnac_application_sets" "example" {
    provider = dnac
    parameters {
      
      name = "string"
    }
}

output "dnac_application_sets_example" {
    value = dnac_application_sets.example
}