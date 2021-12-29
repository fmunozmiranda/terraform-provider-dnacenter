package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTagMember() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Tag.

- Adds members to the tag specified by id

- Removes Tag member from the tag specified by id
`,

		CreateContext: resourceTagMemberCreate,
		ReadContext:   resourceTagMemberRead,
		UpdateContext: resourceTagMemberUpdate,
		DeleteContext: resourceTagMemberDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. Tag ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"member_id": &schema.Schema{
							Description: `memberId path parameter. TagMember id to be removed from tag
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceTagMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTagMemberAddMembersToTheTag(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vMemberID, okMemberID := resourceItem["member_id"]
	vvMemberID := interfaceToString(vMemberID)
	if okID && vvID != "" {
		getResponse1, _, err := client.Tag.GetTagMembersByID(vvID, nil)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["member_id"] = vvMemberID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Tag.AddMembersToTheTag(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddMembersToTheTag", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddMembersToTheTag", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["member_id"] = vvMemberID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceTagMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vMemberType, okMemberType := resourceMap["member_type"]
	vOffset, okOffset := resourceMap["offset"]
	vLimit, okLimit := resourceMap["limit"]
	vMemberAssociationType, okMemberAssociationType := resourceMap["member_association_type"]
	vLevel, okLevel := resourceMap["level"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTagMembersByID")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetTagMembersByIDQueryParams{}

		queryParams1.MemberType = vMemberType.(string)

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okMemberAssociationType {
			queryParams1.MemberAssociationType = vMemberAssociationType.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}

		response1, restyResp1, err := client.Tag.GetTagMembersByID(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTagMembersByID", err,
				"Failure at GetTagMembersByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	return diags
}

func resourceTagMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceTagMemberRead(ctx, d, m)
}

func resourceTagMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestTagMemberAddMembersToTheTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestTagAddMembersToTheTag {
	var request dnacentersdkgo.RequestTagAddMembersToTheTag
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

//TODO
