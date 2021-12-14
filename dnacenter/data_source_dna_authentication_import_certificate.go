package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceAuthenticationImportCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Authentication Management.

- This method is used to upload a certificate.
Upload the files to the **certFileUpload** and **pkFileUpload** form data fields
`,

		ReadContext: dataSourceAuthenticationImportCertificateRead,
		Schema: map[string]*schema.Schema{
			"list_of_users": &schema.Schema{
				Description: `listOfUsers query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pk_password": &schema.Schema{
				Description: `pkPassword query parameter. Private Key Passsword
`,
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAuthenticationImportCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPkPassword, okPkPassword := d.GetOk("pk_password")
	vListOfUsers, okListOfUsers := d.GetOk("list_of_users")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportCertificate")
		queryParams1 := dnacentersdkgo.ImportCertificateQueryParams{}

		if okPkPassword {
			queryParams1.PkPassword = vPkPassword.(string)
		}
		if okListOfUsers {
			queryParams1.ListOfUsers = interfaceToSliceString(vListOfUsers)
		}

		response1, _, err := client.AuthenticationManagement.ImportCertificate(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportCertificate", err,
				"Failure at ImportCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenAuthenticationManagementImportCertificateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportCertificate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAuthenticationManagementImportCertificateItem(item *dnacentersdkgo.ResponseAuthenticationManagementImportCertificateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
