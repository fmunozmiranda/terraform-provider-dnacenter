
data "dnac_file_namespace_files" "example" {
    provider = dnac
    name_space = "string"
}

output "dnac_file_namespace_files_example" {
    value = data.dnac_file_namespace_files.example.items
}
