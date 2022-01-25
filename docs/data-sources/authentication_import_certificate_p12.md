---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_authentication_import_certificate_p12 Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Authentication Management.
  This method is used to upload a PKCS#12 file.
  Upload the file to the p12FileUpload form data field
---

# dnacenter_authentication_import_certificate_p12 (Data Source)

It performs create operation on Authentication Management.

- This method is used to upload a PKCS#12 file.
Upload the file to the **p12FileUpload** form data field

## Example Usage

```terraform
data "dnacdnacenter_authentication_import_certificate_p12" "example" {
  provider      = dnacenter
  list_of_users = ["string"]
  p12_password  = "******"
  pk_password   = "******"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **file_name** (String) File name.
- **p12_file_path** (String) P12 file absolute path.

### Optional

- **id** (String) The ID of this resource.
- **list_of_users** (List of String) listOfUsers query parameter.
- **p12_password** (String, Sensitive) p12Password query parameter. P12 Passsword
- **pk_password** (String, Sensitive) pkPassword query parameter. Private Key Passsword

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **task_id** (String)
- **url** (String)

