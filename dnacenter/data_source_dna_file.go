package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on File.

- Downloads a file specified by fileId
`,

		ReadContext: dataSourceFileRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_id": &schema.Schema{
				Description: `fileId path parameter. File Identification number
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFileID := d.Get("file_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DownloadAFileByFileID")
		vvFileID := vFileID.(string)

		response1, _, err := client.File.DownloadAFileByFileID(vvFileID)

		if err != nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DownloadAFileByFileID", err,
				"Failure at DownloadAFileByFileID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", response1)

		vItem1 := flattenFileDownloadAFileByFileIDItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DownloadAFileByFileID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())

	}
	return diags
}

func flattenFileDownloadAFileByFileIDItem(item dnacentersdkgo.FileDownload) []map[string]interface{} {
	respItem := make(map[string]interface{})
	return []map[string]interface{}{
		respItem,
	}
}
