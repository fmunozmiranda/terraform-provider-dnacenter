package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceAuthenticationImportCertificateP12() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Authentication Management.

- This method is used to upload a PKCS#12 file.
Upload the file to the **p12FileUpload** form data field
`,

		ReadContext: dataSourceAuthenticationImportCertificateP12Read,
		Schema: map[string]*schema.Schema{
			"list_of_users": &schema.Schema{
				Description: `listOfUsers query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"p12_password": &schema.Schema{
				Description: `p12Password query parameter. P12 Passsword
`,
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
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

func dataSourceAuthenticationImportCertificateP12Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vP12Password, okP12Password := d.GetOk("p12_password")
	vPkPassword, okPkPassword := d.GetOk("pk_password")
	vListOfUsers, okListOfUsers := d.GetOk("list_of_users")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportCertificateP12")
		queryParams1 := dnacentersdkgo.ImportCertificateP12QueryParams{}

		if okP12Password {
			queryParams1.P12Password = vP12Password.(string)
		}
		if okPkPassword {
			queryParams1.PkPassword = vPkPassword.(string)
		}
		if okListOfUsers {
			queryParams1.ListOfUsers = interfaceToSliceString(vListOfUsers)
		}

		response1, _, err := client.AuthenticationManagement.ImportCertificateP12(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportCertificateP12", err,
				"Failure at ImportCertificateP12, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenAuthenticationManagementImportCertificateP12Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportCertificateP12 response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAuthenticationManagementImportCertificateP12Item(item *dnacentersdkgo.ResponseAuthenticationManagementImportCertificateP12Response) []map[string]interface{} {
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
