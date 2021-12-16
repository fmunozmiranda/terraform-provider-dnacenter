
data "dnacenter_event_subscription_details_rest" "example" {
    provider = dnac
    connector_type = "string"
    instance_id = "string"
    name = "string"
}

output "dnacenter_event_subscription_details_rest_example" {
    value = data.dnac_event_subscription_details_rest.example.items
}
