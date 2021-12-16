
data "dnac_network_device_functional_capability" "example" {
    provider = dnac
    device_id = "string"
    function_name = ["string"]
}

output "dnac_network_device_functional_capability_example" {
    value = data.dnac_network_device_functional_capability.example.items
}

data "dnac_network_device_functional_capability" "example" {
    provider = dnac
    id = "string"
}

output "dnac_network_device_functional_capability_example" {
    value = data.dnac_network_device_functional_capability.example.item
}
