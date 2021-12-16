
data "dnac_client_proximity" "example" {
    provider = dnac
    number_days = ------
    time_resolution = ------
    username = "string"
}

output "dnac_client_proximity_example" {
    value = data.dnac_client_proximity.example.item
}
