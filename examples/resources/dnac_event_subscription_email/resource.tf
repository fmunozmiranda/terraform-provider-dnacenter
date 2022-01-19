
resource "dnac_event_subscription_email" "example" {
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
          from_email_address = "string"
          subject = "string"
          to_email_addresses = ["string"]
        }
      }
      subscription_id = "string"
      version = "string"
    }
}

output "dnac_event_subscription_email_example" {
    value = dnac_event_subscription_email.example
}