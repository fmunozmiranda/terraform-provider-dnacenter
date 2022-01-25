---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_swim_import_via_url Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Software Image Management (SWIM).
  Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image
  files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
---

# dnacenter_swim_import_via_url (Data Source)

It performs create operation on Software Image Management (SWIM).

- Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image
files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2

## Example Usage

```terraform
data "dnacdnacenter_swim_import_via_url" "example" {
  provider        = dnacenter
  schedule_at     = "string"
  schedule_desc   = "string"
  schedule_origin = "string"
  payload {

    application_type = "string"
    image_family     = "string"
    source_url       = "string"
    third_party      = "false"
    vendor           = "string"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **payload** (Block List) Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURL (see [below for nested schema](#nestedblock--payload))
- **schedule_at** (String) scheduleAt query parameter. Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
- **schedule_desc** (String) scheduleDesc query parameter. Custom Description (Optional)
- **schedule_origin** (String) scheduleOrigin query parameter. Originator of this call (Optional)

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedblock--payload"></a>
### Nested Schema for `payload`

Optional:

- **application_type** (String)
- **image_family** (String)
- **source_url** (String)
- **third_party** (String)
- **vendor** (String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **task_id** (String)
- **url** (String)

