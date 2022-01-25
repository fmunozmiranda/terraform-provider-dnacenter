---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_rf_profile Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and delete operations on Wireless.
  Create or Update RF profileDelete RF profile(s)
---

# dnacenter_wireless_rf_profile (Resource)

It manages create, read and delete operations on Wireless.

- Create or Update RF profile

- Delete RF profile(s)



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

Required:

- **rf_profile_name** (String) rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted

Optional:

- **channel_width** (String) rf-profile channel width
- **default_rf_profile** (String) isDefault rf-profile
- **enable_brown_field** (String) true if enable brown field for rf-profile else false
- **enable_custom** (String) true if enable custom rf-profile else false
- **enable_radio_type_a** (String) tru if Enable Radio Type A else false
- **enable_radio_type_b** (String) true if Enable Radio Type B else false
- **name** (String) custom RF profile name
- **radio_type_a_properties** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--radio_type_a_properties))
- **radio_type_b_properties** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--radio_type_b_properties))

<a id="nestedblock--parameters--radio_type_a_properties"></a>
### Nested Schema for `parameters.radio_type_a_properties`

Optional:

- **data_rates** (String) Data Rates
- **mandatory_data_rates** (String) Mandatory Data Rates
- **max_power_level** (Number) Max Power Level
- **min_power_level** (Number) Min Power Level
- **parent_profile** (String) Parent rf-profile name
- **power_threshold_v1** (Number) Power Threshold V1
- **radio_channels** (String) Radio Channels
- **rx_sop_threshold** (String) Rx Sop Threshold


<a id="nestedblock--parameters--radio_type_b_properties"></a>
### Nested Schema for `parameters.radio_type_b_properties`

Optional:

- **data_rates** (String) Data Rates
- **mandatory_data_rates** (String) Mandatory Data Rates
- **max_power_level** (Number) Max Power Level
- **min_power_level** (Number) Min Power Level
- **parent_profile** (String) Parent rf-profile name
- **power_threshold_v1** (Number) Power Threshold V1
- **radio_channels** (String) Radio Channels
- **rx_sop_threshold** (String) Rx Sop Threshold



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **channel_width** (String)
- **default_rf_profile** (String)
- **enable_brown_field** (String)
- **enable_custom** (String)
- **enable_radio_type_a** (String)
- **enable_radio_type_b** (String)
- **name** (String)
- **radio_type_a_properties** (List of Object) (see [below for nested schema](#nestedobjatt--item--radio_type_a_properties))
- **radio_type_b_properties** (List of Object) (see [below for nested schema](#nestedobjatt--item--radio_type_b_properties))

<a id="nestedobjatt--item--radio_type_a_properties"></a>
### Nested Schema for `item.radio_type_a_properties`

Read-Only:

- **data_rates** (String)
- **mandatory_data_rates** (String)
- **max_power_level** (Number)
- **min_power_level** (Number)
- **parent_profile** (String)
- **power_threshold_v1** (Number)
- **radio_channels** (String)
- **rx_sop_threshold** (String)


<a id="nestedobjatt--item--radio_type_b_properties"></a>
### Nested Schema for `item.radio_type_b_properties`

Read-Only:

- **data_rates** (String)
- **mandatory_data_rates** (String)
- **max_power_level** (Number)
- **min_power_level** (Number)
- **parent_profile** (String)
- **power_threshold_v1** (Number)
- **radio_channels** (String)
- **rx_sop_threshold** (String)

