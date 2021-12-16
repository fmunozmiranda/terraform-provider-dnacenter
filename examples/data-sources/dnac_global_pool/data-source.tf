
data "dnac_global_pool" "example" {
    provider = dnac
    limit = "string"
    offset = "string"
}

output "dnac_global_pool_example" {
    value = data.dnac_global_pool.example.items
}