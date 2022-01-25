---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_application_sets Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and delete operations on Application Policy.
  Delete existing application-set by it's idCreate new custom application-set/s
---

# dnacenter_application_sets (Resource)

It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **parameters** (Block List) Array of RequestApplicationPolicyCreateApplicationSet (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- **name** (String) Name

Optional:

- **id** (String) Name


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **application_set** (List of Object) (see [below for nested schema](#nestedobjatt--item--application_set))
- **id** (String)
- **indicative_network_identity** (List of Object) (see [below for nested schema](#nestedobjatt--item--indicative_network_identity))
- **name** (String)
- **network_applications** (List of Object) (see [below for nested schema](#nestedobjatt--item--network_applications))
- **network_identity** (List of Object) (see [below for nested schema](#nestedobjatt--item--network_identity))

<a id="nestedobjatt--item--application_set"></a>
### Nested Schema for `item.application_set`

Read-Only:

- **id_ref** (String)


<a id="nestedobjatt--item--indicative_network_identity"></a>
### Nested Schema for `item.indicative_network_identity`

Read-Only:

- **display_name** (String)
- **id** (String)
- **lower_port** (Number)
- **ports** (String)
- **protocol** (String)
- **upper_port** (Number)


<a id="nestedobjatt--item--network_applications"></a>
### Nested Schema for `item.network_applications`

Read-Only:

- **app_protocol** (String)
- **application_sub_type** (String)
- **application_type** (String)
- **category_id** (String)
- **display_name** (String)
- **dscp** (String)
- **engine_id** (String)
- **help_string** (String)
- **id** (String)
- **ignore_conflict** (String)
- **long_description** (String)
- **name** (String)
- **popularity** (Number)
- **rank** (Number)
- **server_name** (String)
- **traffic_class** (String)
- **url** (String)


<a id="nestedobjatt--item--network_identity"></a>
### Nested Schema for `item.network_identity`

Read-Only:

- **display_name** (String)
- **id** (String)
- **lower_port** (Number)
- **ports** (String)
- **protocol** (String)
- **upper_port** (Number)

