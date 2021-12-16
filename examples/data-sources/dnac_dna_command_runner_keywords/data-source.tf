
data "dnac_dna_command_runner_keywords" "example" {
    provider = dnac
}

output "dnac_dna_command_runner_keywords_example" {
    value = data.dnac_dna_command_runner_keywords.example.items
}
