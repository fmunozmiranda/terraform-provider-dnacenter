---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_configuration_template_clone Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Configuration Templates.
  API to clone template
---

# dnacenter_configuration_template_clone (Data Source)

It performs create operation on Configuration Templates.

- API to clone template

## Example Usage

```terraform
data "dnacdnacenter_configuration_template_clone" "example" {
  provider    = dnacenter
  name        = "string"
  project_id  = "string"
  template_id = "string"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) name path parameter. Template name to clone template(Name should be different than existing template name within same project)
- **project_id** (String) projectId path parameter. UUID of the project in which the template needs to be created
- **template_id** (String) templateId path parameter. UUID of the template to clone it

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **task_id** (String)
- **url** (String)

