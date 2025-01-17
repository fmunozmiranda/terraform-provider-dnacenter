terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

resource "dna_sda_fabric" "response" {
  provider    = dnacenter
  fabric_name = "MyFabricName2"
}
output "dna_sda_fabric_response" {
  value = dna_sda_fabric.response
}
