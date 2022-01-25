---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_nfv_provision_details Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Site Design.
  Checks the provisioning detail of an ENCS device including log information.
---

# dnacenter_nfv_provision_details (Data Source)

It performs create operation on Site Design.

- Checks the provisioning detail of an ENCS device including log information.

## Example Usage

```terraform
data "dnacdnacenter_nfv_provision_details" "example" {
  provider  = dnacenter
  device_ip = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **device_ip** (String) Device Ip
- **id** (String) The ID of this resource.
- **persistbapioutput** (String) __persistbapioutput header parameter. Persist bapi sync response
- **runsync** (String) __runsync header parameter. Enable this parameter to execute the API and return a response synchronously
- **runsynctimeout** (String) __runsynctimeout header parameter. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **execution_id** (String)
- **execution_status_url** (String)
- **message** (String)

