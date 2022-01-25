
resource "dnacenter_sda_fabric" "example" {
  provider = dnacenter
  parameters {
    fabric_name = "b"
  }
}

output "dnacenter_sda_fabric_example" {
  value = dnacenter_sda_fabric.example
}