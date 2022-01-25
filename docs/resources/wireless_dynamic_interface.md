---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_dynamic_interface Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and delete operations on Wireless.
  API to create or update an dynamic interfaceDelete a dynamic interface
---

# dnacenter_wireless_dynamic_interface (Resource)

It manages create, read and delete operations on Wireless.

- API to create or update an dynamic interface

- Delete a dynamic interface



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **parameters** (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **interface_name** (String) dynamic-interface name
- **vlan_id** (Number) Vlan Id


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **interface_name** (String)
- **vlan_id** (Number)

