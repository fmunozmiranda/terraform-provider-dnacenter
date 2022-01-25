---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_configuration_template_create Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Configuration Templates.
  API to create a template by project id.
---

# dnacenter_configuration_template_create (Data Source)

It performs create operation on Configuration Templates.

- API to create a template by project id.

## Example Usage

```terraform
data "dnacdnacenter_configuration_template_create" "example" {
  provider   = dnacenter
  project_id = "string"
  id         = "string"
  name       = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **project_id** (String) projectId path parameter. UUID of the project in which the template needs to be created

### Optional

- **author** (String) Author of template
- **composite** (String) Is it composite template
- **containing_templates** (Block List) (see [below for nested schema](#nestedblock--containing_templates))
- **create_time** (Number) Create time of template
- **custom_params_order** (String) Custom Params Order
- **description** (String) Description of template
- **device_types** (Block List) (see [below for nested schema](#nestedblock--device_types))
- **failure_policy** (String) Define failure policy if template provisioning fails
- **id** (String) UUID of template
- **language** (String) Template language (JINJA or VELOCITY)
- **last_update_time** (Number) Update time of template
- **latest_version_time** (Number) Latest versioned template time
- **name** (String) Name of template
- **parent_template_id** (String) Parent templateID
- **project_name** (String) Project name
- **rollback_template_content** (String) Rollback template content
- **rollback_template_params** (Block List) (see [below for nested schema](#nestedblock--rollback_template_params))
- **software_type** (String) Applicable device software type
- **software_variant** (String) Applicable device software variant
- **software_version** (String) Applicable device software version
- **tags** (Block List) (see [below for nested schema](#nestedblock--tags))
- **template_content** (String) Template content
- **template_params** (Block List) (see [below for nested schema](#nestedblock--template_params))
- **validation_errors** (Block List) (see [below for nested schema](#nestedblock--validation_errors))
- **version** (String) Current version of template

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedblock--containing_templates"></a>
### Nested Schema for `containing_templates`

Optional:

- **composite** (String) Is it composite template
- **description** (String) Description of template
- **device_types** (Block List) (see [below for nested schema](#nestedblock--containing_templates--device_types))
- **id** (String) UUID of template
- **language** (String) Template language (JINJA or VELOCITY)
- **name** (String) Name of template
- **project_name** (String) Project name
- **rollback_template_params** (Block List) (see [below for nested schema](#nestedblock--containing_templates--rollback_template_params))
- **tags** (Block List) (see [below for nested schema](#nestedblock--containing_templates--tags))
- **template_content** (String) Template content
- **template_params** (Block List) (see [below for nested schema](#nestedblock--containing_templates--template_params))
- **version** (String) Current version of template

<a id="nestedblock--containing_templates--device_types"></a>
### Nested Schema for `containing_templates.device_types`

Optional:

- **product_family** (String) Device family
- **product_series** (String) Device series
- **product_type** (String) Device type


<a id="nestedblock--containing_templates--rollback_template_params"></a>
### Nested Schema for `containing_templates.rollback_template_params`

Optional:

- **binding** (String) Bind to source
- **custom_order** (Number) CustomOrder of template param
- **data_type** (String) Datatype of template param
- **default_value** (String) Default value of template param
- **description** (String) Description of template param
- **display_name** (String) Display name of param
- **group** (String) group
- **id** (String) UUID of template param
- **instruction_text** (String) Instruction text for param
- **key** (String) key
- **not_param** (String) Is it not a variable
- **order** (Number) Order of template param
- **param_array** (String) Is it an array
- **parameter_name** (String) Name of template param
- **provider** (String) provider
- **range** (Block List) (see [below for nested schema](#nestedblock--containing_templates--rollback_template_params--range))
- **required** (String) Is param required
- **selection** (Block List) (see [below for nested schema](#nestedblock--containing_templates--rollback_template_params--selection))

<a id="nestedblock--containing_templates--rollback_template_params--range"></a>
### Nested Schema for `containing_templates.rollback_template_params.range`

Optional:

- **id** (String) UUID of range
- **max_value** (Number) Max value of range
- **min_value** (Number) Min value of range


<a id="nestedblock--containing_templates--rollback_template_params--selection"></a>
### Nested Schema for `containing_templates.rollback_template_params.selection`

Optional:

- **default_selected_values** (List of String) Default selection values
- **id** (String) UUID of selection
- **selection_type** (String) Type of selection(SINGLE_SELECT or MULTI_SELECT)
- **selection_values** (List of String) Selection values



<a id="nestedblock--containing_templates--tags"></a>
### Nested Schema for `containing_templates.tags`

Optional:

- **id** (String) UUID of tag
- **name** (String) Name of tag


<a id="nestedblock--containing_templates--template_params"></a>
### Nested Schema for `containing_templates.template_params`

Optional:

- **binding** (String) Bind to source
- **custom_order** (Number) CustomOrder of template param
- **data_type** (String) Datatype of template param
- **default_value** (String) Default value of template param
- **description** (String) Description of template param
- **display_name** (String) Display name of param
- **group** (String) group
- **id** (String) UUID of template param
- **instruction_text** (String) Instruction text for param
- **key** (String) key
- **not_param** (String) Is it not a variable
- **order** (Number) Order of template param
- **param_array** (String) Is it an array
- **parameter_name** (String) Name of template param
- **provider** (String) provider
- **range** (Block List) (see [below for nested schema](#nestedblock--containing_templates--template_params--range))
- **required** (String) Is param required
- **selection** (Block List) (see [below for nested schema](#nestedblock--containing_templates--template_params--selection))

<a id="nestedblock--containing_templates--template_params--range"></a>
### Nested Schema for `containing_templates.template_params.range`

Optional:

- **id** (String) UUID of range
- **max_value** (Number) Max value of range
- **min_value** (Number) Min value of range


<a id="nestedblock--containing_templates--template_params--selection"></a>
### Nested Schema for `containing_templates.template_params.selection`

Optional:

- **default_selected_values** (List of String) Default selection values
- **id** (String) UUID of selection
- **selection_type** (String) Type of selection(SINGLE_SELECT or MULTI_SELECT)
- **selection_values** (List of String) Selection values




<a id="nestedblock--device_types"></a>
### Nested Schema for `device_types`

Optional:

- **product_family** (String) Device family
- **product_series** (String) Device series
- **product_type** (String) Device type


<a id="nestedblock--rollback_template_params"></a>
### Nested Schema for `rollback_template_params`

Optional:

- **binding** (String) Bind to source
- **custom_order** (Number) CustomOrder of template param
- **data_type** (String) Datatype of template param
- **default_value** (String) Default value of template param
- **description** (String) Description of template param
- **display_name** (String) Display name of param
- **group** (String) group
- **id** (String) UUID of template param
- **instruction_text** (String) Instruction text for param
- **key** (String) key
- **not_param** (String) Is it not a variable
- **order** (Number) Order of template param
- **param_array** (String) Is it an array
- **parameter_name** (String) Name of template param
- **provider** (String) provider
- **range** (Block List) (see [below for nested schema](#nestedblock--rollback_template_params--range))
- **required** (String) Is param required
- **selection** (Block List) (see [below for nested schema](#nestedblock--rollback_template_params--selection))

<a id="nestedblock--rollback_template_params--range"></a>
### Nested Schema for `rollback_template_params.range`

Optional:

- **id** (String) UUID of range
- **max_value** (Number) Max value of range
- **min_value** (Number) Min value of range


<a id="nestedblock--rollback_template_params--selection"></a>
### Nested Schema for `rollback_template_params.selection`

Optional:

- **default_selected_values** (List of String) Default selection values
- **id** (String) UUID of selection
- **selection_type** (String) Type of selection(SINGLE_SELECT or MULTI_SELECT)
- **selection_values** (List of String) Selection values



<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Optional:

- **id** (String) UUID of tag
- **name** (String) Name of tag


<a id="nestedblock--template_params"></a>
### Nested Schema for `template_params`

Optional:

- **binding** (String) Bind to source
- **custom_order** (Number) CustomOrder of template param
- **data_type** (String) Datatype of template param
- **default_value** (String) Default value of template param
- **description** (String) Description of template param
- **display_name** (String) Display name of param
- **group** (String) group
- **id** (String) UUID of template param
- **instruction_text** (String) Instruction text for param
- **key** (String) key
- **not_param** (String) Is it not a variable
- **order** (Number) Order of template param
- **param_array** (String) Is it an array
- **parameter_name** (String) Name of template param
- **provider** (String) provider
- **range** (Block List) (see [below for nested schema](#nestedblock--template_params--range))
- **required** (String) Is param required
- **selection** (Block List) (see [below for nested schema](#nestedblock--template_params--selection))

<a id="nestedblock--template_params--range"></a>
### Nested Schema for `template_params.range`

Optional:

- **id** (String) UUID of range
- **max_value** (Number) Max value of range
- **min_value** (Number) Min value of range


<a id="nestedblock--template_params--selection"></a>
### Nested Schema for `template_params.selection`

Optional:

- **default_selected_values** (List of String) Default selection values
- **id** (String) UUID of selection
- **selection_type** (String) Type of selection(SINGLE_SELECT or MULTI_SELECT)
- **selection_values** (List of String) Selection values



<a id="nestedblock--validation_errors"></a>
### Nested Schema for `validation_errors`

Optional:

- **rollback_template_errors** (List of String) Validation or design conflicts errors of rollback template
- **template_errors** (List of String) Validation or design conflicts errors
- **template_id** (String) UUID of template
- **template_version** (String) Current version of template


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **task_id** (String)
- **url** (String)

