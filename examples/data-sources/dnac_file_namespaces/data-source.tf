
data "dnac_file_namespaces" "example" {
    provider = dnac
}

output "dnac_file_namespaces_example" {
    value = data.dnac_file_namespaces.example.items
}
