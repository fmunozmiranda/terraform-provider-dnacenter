package dnacenter

import (
	"context"
	"fmt"
	"io"
	"os"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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
			"cert_file_path": &schema.Schema{
				Description: `Cert file absolute path.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_name": &schema.Schema{
				Description: `File name.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_name2": &schema.Schema{
				Description: `File name.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"list_of_users": &schema.Schema{
				Description: `listOfUsers query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pk_file_path": &schema.Schema{
				Description: `Pk file absolute path.`,
				Type:        schema.TypeString,
				Required:    true,
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
	vCertFilePath := d.Get("cert_file_path")
	vFileName := d.Get("file_name")
	vFileName2 := d.Get("file_name2")
	vPkFilePath := d.Get("pk_file_path")

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

		f, err := os.Open(vCertFilePath.(string))
		if err != nil {
			fmt.Println(err)
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportCertificateP12", err,
				"Failure at ImportCertificateP12, unexpected response", ""))
			return diags
		}
		second_file, err := os.Open(vPkFilePath.(string))
		if err != nil {
			fmt.Println(err)
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportCertificateP12", err,
				"Failure at ImportCertificateP12, unexpected response", ""))
			return diags
		}
		defer func() {
			if err = f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		var r io.Reader
		var r2 io.Reader
		r = f
		r2 = second_file
		response1, restyResp1, err := client.AuthenticationManagement.ImportCertificate(
			&queryParams1,
			&dnacentersdkgo.ImportCertificateMultipartFields{
				PkFileUploadName:   vFileName.(string),
				PkFileUpload:       r,
				CertFileUploadName: vFileName2.(string),
				CertFileUpload:     r2,
			},
		)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportCertificate", err,
				"Failure at ImportCertificate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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
