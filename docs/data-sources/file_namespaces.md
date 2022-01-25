---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_file_namespaces Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on File.
  Returns list of available namespaces
---

# dnacenter_file_namespaces (Data Source)

It performs read operation on File.

- Returns list of available namespaces

## Example Usage

```terraform
data "dnacenter_file_namespaces" "example" {
  provider = dnacenter
}

output "dnacenter_file_namespaces_example" {
  value = data.dnacenter_file_namespaces.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **items** (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **response** (List of String)
- **version** (String)

