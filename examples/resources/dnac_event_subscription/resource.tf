
resource "dnac_event_subscription" "example" {
    provider = dnac
    parameters {
      
      description = "string"
      filter {
        
        event_ids = ["string"]
      }
      name = "string"
      subscription_endpoints {
        
        instance_id = "string"
        subscription_details {
          
          connector_type = "string"
          method = "string"
          name = "string"
          url = "string"
        }
      }
      subscription_id = "string"
      version = "string"
    }
}

output "dnac_event_subscription_example" {
    value = dnac_event_subscription.example
}