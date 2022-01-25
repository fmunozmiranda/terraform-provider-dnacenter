---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_pnp_device_reset Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Device Onboarding (PnP).
  Recovers a device from a Workflow Execution Error state
---

# dnacenter_pnp_device_reset (Data Source)

It performs create operation on Device Onboarding (PnP).

- Recovers a device from a Workflow Execution Error state

## Example Usage

```terraform
data "dnacdnacenter_pnp_device_reset" "example" {
  provider = dnacenter
  config_list {

    config_id = "string"
    config_parameters {

      key   = "string"
      value = "string"
    }
  }
  device_id                  = "string"
  license_level              = "string"
  license_type               = "string"
  top_of_stack_serial_number = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **device_reset_list** (Block List) (see [below for nested schema](#nestedblock--device_reset_list))
- **id** (String) The ID of this resource.
- **project_id** (String)
- **workflow_id** (String)

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedblock--device_reset_list"></a>
### Nested Schema for `device_reset_list`

Optional:

- **config_list** (Block List) (see [below for nested schema](#nestedblock--device_reset_list--config_list))
- **device_id** (String)
- **license_level** (String)
- **license_type** (String)
- **top_of_stack_serial_number** (String)

<a id="nestedblock--device_reset_list--config_list"></a>
### Nested Schema for `device_reset_list.config_list`

Optional:

- **config_id** (String)
- **config_parameters** (Block List) (see [below for nested schema](#nestedblock--device_reset_list--config_list--config_parameters))

<a id="nestedblock--device_reset_list--config_list--config_parameters"></a>
### Nested Schema for `device_reset_list.config_list.config_parameters`

Optional:

- **key** (String)
- **value** (String)




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **json_array_response** (List of String)
- **json_response** (String)
- **message** (String)
- **status_code** (Number)

