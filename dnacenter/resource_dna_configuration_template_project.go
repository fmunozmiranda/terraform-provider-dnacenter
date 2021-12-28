package dnacenter

import (
	"context"
	"reflect"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationTemplateProject() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Configuration Templates.

- This resource is used to create a new project.

- This resource is used to update an existing project.

- Deletes the project by its id
`,

		CreateContext: resourceConfigurationTemplateProjectCreate,
		ReadContext:   resourceConfigurationTemplateProjectRead,
		UpdateContext: resourceConfigurationTemplateProjectUpdate,
		DeleteContext: resourceConfigurationTemplateProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//{'data': {'parameters': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'createTime': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Create time of project\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of project\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of project\n'}, 'lastUpdateTime': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Update time of project\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of project\n'}, 'templates': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'List of templates within the project\n', 'Elem': {'Schema': {'tags': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'author': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Author of template\n'}, 'composite': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it composite template\n'}, 'containingTemplates': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'composite': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it composite template\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template\n'}, 'deviceTypes': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'productFamily': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device family\n'}, 'productSeries': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device series\n'}, 'productType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device type\n'}}}}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'language': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Template language (JINJA or VELOCITY)\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template\n'}, 'projectName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Project name\n'}, 'rollbackTemplateParams': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'templateContent': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Template content\n'}, 'templateParams': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'version': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}, 'createTime': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Create time of template\n'}, 'customParamsOrder': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Custom Params Order\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template\n'}, 'deviceTypes': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'productFamily': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device family\n'}, 'productSeries': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device series\n'}, 'productType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Device type\n'}}}}, 'documentDatabase': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Document Database\n'}, 'failurePolicy': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Define failure policy if template provisioning fails\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'language': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Template language (JINJA or VELOCITY)\n'}, 'lastUpdateTime': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Update time of template\n'}, 'latestVersionTime': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Latest versioned template time\n'}, 'name': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template\n'}, 'parentTemplateId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Parent templateID\n'}, 'projectAssociated': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Project Associated\n'}, 'projectId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Project UUID\n'}, 'projectName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Project name\n'}, 'rollbackTemplateContent': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Rollback template content\n'}, 'rollbackTemplateParams': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'softwareType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software type\n'}, 'softwareVariant': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software variant\n'}, 'softwareVersion': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software version\n'}, 'templateContent': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Template content\n'}, 'templateParams': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Optional': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Optional': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'validationErrors': {'Optional': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'rollbackTemplateErrors': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Validation or design conflicts errors of rollback template\n', 'Elem': {'Schema': {}}}, 'templateErrors': {'Optional': 'true', 'Type': 'schema.TypeList', 'Description': 'Validation or design conflicts errors\n', 'Elem': {'Schema': {}}}, 'templateId': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'templateVersion': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}, 'version': {'Optional': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}, 'projectId': {'Required': 'true', 'Type': 'schema.TypeString', 'Description': 'projectId path parameter. projectId(UUID) of project to be deleted\n'}}}}, 'item': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'createTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Create time of project\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of project\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of project\n'}, 'lastUpdateTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Update time of project\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of project\n'}, 'templates': {'Computed': 'true', 'Type': 'schema.TypeList', 'Description': 'List of templates within the project\n', 'Elem': {'Schema': {'tags': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'author': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Author of template\n'}, 'composite': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it composite template\n'}, 'containingTemplates': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'tags': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of tag\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of tag\n'}}}}, 'composite': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it composite template\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template\n'}, 'deviceTypes': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'productFamily': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device family\n'}, 'productSeries': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device series\n'}, 'productType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device type\n'}}}}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'language': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Template language (JINJA or VELOCITY)\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template\n'}, 'projectName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Project name\n'}, 'rollbackTemplateParams': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'templateContent': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Template content\n'}, 'templateParams': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'version': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}, 'createTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Create time of template\n'}, 'customParamsOrder': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Custom Params Order\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template\n'}, 'deviceTypes': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'productFamily': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device family\n'}, 'productSeries': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device series\n'}, 'productType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Device type\n'}}}}, 'documentDatabase': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Document Database\n'}, 'failurePolicy': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Define failure policy if template provisioning fails\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'language': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Template language (JINJA or VELOCITY)\n'}, 'lastUpdateTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Update time of template\n'}, 'latestVersionTime': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Latest versioned template time\n'}, 'name': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template\n'}, 'parentTemplateId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Parent templateID\n'}, 'projectAssociated': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Project Associated\n'}, 'projectId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Project UUID\n'}, 'projectName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Project name\n'}, 'rollbackTemplateContent': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Rollback template content\n'}, 'rollbackTemplateParams': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'softwareType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software type\n'}, 'softwareVariant': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software variant\n'}, 'softwareVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Applicable device software version\n'}, 'templateContent': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Template content\n'}, 'templateParams': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'binding': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Bind to source\n'}, 'customOrder': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'CustomOrder of template param\n'}, 'dataType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Datatype of template param\n'}, 'defaultValue': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Default value of template param\n'}, 'description': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Description of template param\n'}, 'displayName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Display name of param\n'}, 'group': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'group\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template param\n'}, 'instructionText': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Instruction text for param\n'}, 'key': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'key\n'}, 'notParam': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it not a variable\n'}, 'order': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Order of template param\n'}, 'paramArray': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is it an array\n'}, 'parameterName': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Name of template param\n'}, 'provider': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'provider\n'}, 'range': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Schema': {'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of range\n'}, 'maxValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Max value of range\n'}, 'minValue': {'Computed': 'true', 'Type': 'schema.TypeInt', 'Description': 'Min value of range\n'}}}}, 'required': {'Computed': 'true', 'Type': 'schema.TypeBool', 'Description': 'Is param required\n'}, 'selection': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'defaultSelectedValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'Elem': {'Type': 'schema.TypeString'}, 'Description': 'Default selection values\n'}, 'id': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of selection\n'}, 'selectionType': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n'}, 'selectionValues': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Description': 'Selection values\n', 'Elem': {'Schema': {}}}}}}}}}, 'validationErrors': {'Computed': 'true', 'Type': 'schema.TypeList', 'MaxItems': 1, 'Elem': {'Schema': {'rollbackTemplateErrors': {'Computed': 'true', 'Type': 'schema.TypeList', 'Description': 'Validation or design conflicts errors of rollback template\n', 'Elem': {'Schema': {}}}, 'templateErrors': {'Computed': 'true', 'Type': 'schema.TypeList', 'Description': 'Validation or design conflicts errors\n', 'Elem': {'Schema': {}}}, 'templateId': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'UUID of template\n'}, 'templateVersion': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}, 'version': {'Computed': 'true', 'Type': 'schema.TypeString', 'Description': 'Current version of template\n'}}}}}}}}, 'metadata': {'item': {'operation_id': [['CreateProject', 'UpdateProject'], 'GetsTheDetailsOfAGivenProject'], 'new_flat_structure': [[{'RequestConfigurationTemplatesCreateProject': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTags'}, {'name': 'createTime', 'description': 'Create time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': 'Description of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templates', 'description': 'List of templates within the project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplates'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesTags'}, {'name': 'author', 'description': 'Author of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'containingTemplates', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates'}, {'name': 'createTime', 'description': 'Create time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'customParamsOrder', 'description': 'Custom Params Order\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes'}, {'name': 'documentDatabase', 'description': 'Document Database\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'failurePolicy', 'description': 'Define failure policy if template provisioning fails\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'latestVersionTime', 'description': 'Latest versioned template time\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'parentTemplateId', 'description': 'Parent templateID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectAssociated', 'description': 'Project Associated\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'projectId', 'description': 'Project UUID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateContent', 'description': 'Rollback template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams'}, {'name': 'softwareType', 'description': 'Applicable device software type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVariant', 'description': 'Applicable device software variant\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': 'Applicable device software version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams'}, {'name': 'validationErrors', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors': {'type': 'obj', 'data': [{'name': 'rollbackTemplateErrors', 'description': 'Validation or design conflicts errors of rollback template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors'}, {'name': 'templateErrors', 'description': 'Validation or design conflicts errors\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors'}, {'name': 'templateId', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateVersion', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}, {'RequestConfigurationTemplatesUpdateProject': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTags'}, {'name': 'createTime', 'description': 'Create time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': 'Description of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templates', 'description': 'List of templates within the project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplates'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesTags'}, {'name': 'author', 'description': 'Author of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'containingTemplates', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates'}, {'name': 'createTime', 'description': 'Create time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'customParamsOrder', 'description': 'Custom Params Order\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes'}, {'name': 'documentDatabase', 'description': 'Document Database\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'failurePolicy', 'description': 'Define failure policy if template provisioning fails\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'latestVersionTime', 'description': 'Latest versioned template time\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'parentTemplateId', 'description': 'Parent templateID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectAssociated', 'description': 'Project Associated\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'projectId', 'description': 'Project UUID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateContent', 'description': 'Rollback template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams'}, {'name': 'softwareType', 'description': 'Applicable device software type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVariant', 'description': 'Applicable device software variant\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': 'Applicable device software version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams'}, {'name': 'validationErrors', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors': {'type': 'obj', 'data': [{'name': 'rollbackTemplateErrors', 'description': 'Validation or design conflicts errors of rollback template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors'}, {'name': 'templateErrors', 'description': 'Validation or design conflicts errors\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors'}, {'name': 'templateId', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateVersion', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], {'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags'}, {'name': 'createTime', 'description': 'Create time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'description', 'description': 'Description of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templates', 'description': 'List of templates within the project\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags'}, {'name': 'author', 'description': 'Author of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'containingTemplates', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates'}, {'name': 'createTime', 'description': 'Create time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'customParamsOrder', 'description': 'Custom Params Order\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes'}, {'name': 'documentDatabase', 'description': 'Document Database\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'failurePolicy', 'description': 'Define failure policy if template provisioning fails\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'lastUpdateTime', 'description': 'Update time of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'latestVersionTime', 'description': 'Latest versioned template time\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'parentTemplateId', 'description': 'Parent templateID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectAssociated', 'description': 'Project Associated\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'projectId', 'description': 'Project UUID\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateContent', 'description': 'Rollback template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams'}, {'name': 'softwareType', 'description': 'Applicable device software type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVariant', 'description': 'Applicable device software variant\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'softwareVersion', 'description': 'Applicable device software version\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams'}, {'name': 'validationErrors', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates': {'type': 'obj', 'data': [{'name': 'tags', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags'}, {'name': 'composite', 'description': 'Is it composite template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'description', 'description': 'Description of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'deviceTypes', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes'}, {'name': 'id', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'language', 'description': 'Template language (JINJA or VELOCITY)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'projectName', 'description': 'Project name\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'rollbackTemplateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams'}, {'name': 'templateContent', 'description': 'Template content\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateParams', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams'}, {'name': 'version', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'name', 'description': 'Name of tag\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes': {'type': 'obj', 'data': [{'name': 'productFamily', 'description': 'Device family\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productSeries', 'description': 'Device series\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'productType', 'description': 'Device type\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams': {'type': 'obj', 'data': [{'name': 'binding', 'description': 'Bind to source\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'customOrder', 'description': 'CustomOrder of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'dataType', 'description': 'Datatype of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'defaultValue', 'description': 'Default value of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'description', 'description': 'Description of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'displayName', 'description': 'Display name of param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'group', 'description': 'group\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'id', 'description': 'UUID of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'instructionText', 'description': 'Instruction text for param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'key', 'description': 'key\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'notParam', 'description': 'Is it not a variable\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'order', 'description': 'Order of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'paramArray', 'description': 'Is it an array\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'parameterName', 'description': 'Name of template param\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'provider', 'description': 'provider\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'range', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange'}, {'name': 'required', 'description': 'Is param required\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'bool'}, {'name': 'selection', 'description': None, 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange': {'type': 'obj', 'data': [{'name': 'id', 'description': 'UUID of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'maxValue', 'description': 'Max value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}, {'name': 'minValue', 'description': 'Min value of range\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'int'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection': {'type': 'obj', 'data': [{'name': 'defaultSelectedValues', 'description': 'Default selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]string'}, {'name': 'id', 'description': 'UUID of selection\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionType', 'description': 'Type of selection(SINGLE_SELECT or MULTI_SELECT)\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'selectionValues', 'description': 'Selection values\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors': {'type': 'obj', 'data': [{'name': 'rollbackTemplateErrors', 'description': 'Validation or design conflicts errors of rollback template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors'}, {'name': 'templateErrors', 'description': 'Validation or design conflicts errors\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': '[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors'}, {'name': 'templateId', 'description': 'UUID of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}, {'name': 'templateVersion', 'description': 'Current version of template\n', 'has_rename': None, 'alt_name': None, 'endpoint_name': None, 'type': 'string'}], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}, 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors': {'type': 'obj', 'data': [], 'epType': 'json', 'has_rename': None, 'alt_name': None, 'endpoint_name': None}}], 'flatten_structure_key': [['RequestConfigurationTemplatesCreateProject', 'RequestConfigurationTemplatesUpdateProject'], 'ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject'], 'access_list': [[['tags', 'createTime', 'description', 'id', 'lastUpdateTime', 'name', 'templates'], ['tags', 'createTime', 'description', 'id', 'lastUpdateTime', 'name', 'templates']], []]}}}
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"containing_templates": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"composite": &schema.Schema{
													Description: `Is it composite template
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"device_types": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"product_family": &schema.Schema{
																Description: `Device family
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"product_series": &schema.Schema{
																Description: `Device series
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"product_type": &schema.Schema{
																Description: `Device type
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `UUID of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"language": &schema.Schema{
													Description: `Template language (JINJA or VELOCITY)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"project_name": &schema.Schema{
													Description: `Project name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"rollback_template_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"binding": &schema.Schema{
																Description: `Bind to source
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"custom_order": &schema.Schema{
																Description: `CustomOrder of template param
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"data_type": &schema.Schema{
																Description: `Datatype of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"default_value": &schema.Schema{
																Description: `Default value of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": &schema.Schema{
																Description: `Description of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"display_name": &schema.Schema{
																Description: `Display name of param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"group": &schema.Schema{
																Description: `group
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"id": &schema.Schema{
																Description: `UUID of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"instruction_text": &schema.Schema{
																Description: `Instruction text for param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"key": &schema.Schema{
																Description: `key
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"not_param": &schema.Schema{
																Description: `Is it not a variable
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"order": &schema.Schema{
																Description: `Order of template param
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"param_array": &schema.Schema{
																Description: `Is it an array
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"parameter_name": &schema.Schema{
																Description: `Name of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"provider": &schema.Schema{
																Description: `provider
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"range": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `UUID of range
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"max_value": &schema.Schema{
																			Description: `Max value of range
`,
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"min_value": &schema.Schema{
																			Description: `Min value of range
`,
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"required": &schema.Schema{
																Description: `Is param required
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_selected_values": &schema.Schema{
																			Description: `Default selection values
`,
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"id": &schema.Schema{
																			Description: `UUID of selection
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"selection_type": &schema.Schema{
																			Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"selection_values": &schema.Schema{
																			Description: `Selection values
`,
																			Type:     schema.TypeList,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"tags": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of tag
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"name": &schema.Schema{
																Description: `Name of tag
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"template_content": &schema.Schema{
													Description: `Template content
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"template_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"binding": &schema.Schema{
																Description: `Bind to source
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"custom_order": &schema.Schema{
																Description: `CustomOrder of template param
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"data_type": &schema.Schema{
																Description: `Datatype of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"default_value": &schema.Schema{
																Description: `Default value of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"description": &schema.Schema{
																Description: `Description of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"display_name": &schema.Schema{
																Description: `Display name of param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"group": &schema.Schema{
																Description: `group
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"id": &schema.Schema{
																Description: `UUID of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"instruction_text": &schema.Schema{
																Description: `Instruction text for param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"key": &schema.Schema{
																Description: `key
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"not_param": &schema.Schema{
																Description: `Is it not a variable
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"order": &schema.Schema{
																Description: `Order of template param
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"param_array": &schema.Schema{
																Description: `Is it an array
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"parameter_name": &schema.Schema{
																Description: `Name of template param
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"provider": &schema.Schema{
																Description: `provider
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"range": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `UUID of range
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"max_value": &schema.Schema{
																			Description: `Max value of range
`,
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"min_value": &schema.Schema{
																			Description: `Min value of range
`,
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},
															"required": &schema.Schema{
																Description: `Is param required
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_selected_values": &schema.Schema{
																			Description: `Default selection values
`,
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"id": &schema.Schema{
																			Description: `UUID of selection
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"selection_type": &schema.Schema{
																			Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"selection_values": &schema.Schema{
																			Description: `Selection values
`,
																			Type:     schema.TypeList,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"version": &schema.Schema{
													Description: `Current version of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"create_time": &schema.Schema{
										Description: `Create time of template
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"custom_params_order": &schema.Schema{
										Description: `Custom Params Order
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"document_database": &schema.Schema{
										Description: `Document Database
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_policy": &schema.Schema{
										Description: `Define failure policy if template provisioning fails
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_update_time": &schema.Schema{
										Description: `Update time of template
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"latest_version_time": &schema.Schema{
										Description: `Latest versioned template time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"parent_template_id": &schema.Schema{
										Description: `Parent templateID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_associated": &schema.Schema{
										Description: `Project Associated
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_id": &schema.Schema{
										Description: `Project UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rollback_template_content": &schema.Schema{
										Description: `Rollback template content
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeList,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"software_type": &schema.Schema{
										Description: `Applicable device software type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_variant": &schema.Schema{
										Description: `Applicable device software variant
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_version": &schema.Schema{
										Description: `Applicable device software version
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeList,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"validation_errors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rollback_template_errors": &schema.Schema{
													Description: `Validation or design conflicts errors of rollback template
`,
													Type:     schema.TypeList,
													Computed: true,
												},
												"template_errors": &schema.Schema{
													Description: `Validation or design conflicts errors
`,
													Type:     schema.TypeList,
													Computed: true,
												},
												"template_id": &schema.Schema{
													Description: `UUID of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"template_version": &schema.Schema{
													Description: `Current version of template
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"project_id": &schema.Schema{
							Description: `projectId path parameter. projectId(UUID) of project to be deleted
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"containing_templates": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"composite": &schema.Schema{
													Description: `Is it composite template
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"description": &schema.Schema{
													Description: `Description of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"device_types": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"product_family": &schema.Schema{
																Description: `Device family
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"product_series": &schema.Schema{
																Description: `Device series
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"product_type": &schema.Schema{
																Description: `Device type
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `UUID of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"language": &schema.Schema{
													Description: `Template language (JINJA or VELOCITY)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"project_name": &schema.Schema{
													Description: `Project name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"rollback_template_params": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"binding": &schema.Schema{
																Description: `Bind to source
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"custom_order": &schema.Schema{
																Description: `CustomOrder of template param
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"data_type": &schema.Schema{
																Description: `Datatype of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"default_value": &schema.Schema{
																Description: `Default value of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"description": &schema.Schema{
																Description: `Description of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"display_name": &schema.Schema{
																Description: `Display name of param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"group": &schema.Schema{
																Description: `group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"id": &schema.Schema{
																Description: `UUID of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"instruction_text": &schema.Schema{
																Description: `Instruction text for param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"key": &schema.Schema{
																Description: `key
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"not_param": &schema.Schema{
																Description: `Is it not a variable
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"order": &schema.Schema{
																Description: `Order of template param
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"param_array": &schema.Schema{
																Description: `Is it an array
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"parameter_name": &schema.Schema{
																Description: `Name of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"provider": &schema.Schema{
																Description: `provider
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"range": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `UUID of range
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"max_value": &schema.Schema{
																			Description: `Max value of range
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"min_value": &schema.Schema{
																			Description: `Min value of range
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"required": &schema.Schema{
																Description: `Is param required
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"selection": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_selected_values": &schema.Schema{
																			Description: `Default selection values
`,
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"id": &schema.Schema{
																			Description: `UUID of selection
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selection_type": &schema.Schema{
																			Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selection_values": &schema.Schema{
																			Description: `Selection values
`,
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																		},
																	},
																},
															},
														},
													},
												},
												"tags": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of tag
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"name": &schema.Schema{
																Description: `Name of tag
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"template_content": &schema.Schema{
													Description: `Template content
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"template_params": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"binding": &schema.Schema{
																Description: `Bind to source
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"custom_order": &schema.Schema{
																Description: `CustomOrder of template param
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"data_type": &schema.Schema{
																Description: `Datatype of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"default_value": &schema.Schema{
																Description: `Default value of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"description": &schema.Schema{
																Description: `Description of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"display_name": &schema.Schema{
																Description: `Display name of param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"group": &schema.Schema{
																Description: `group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"id": &schema.Schema{
																Description: `UUID of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"instruction_text": &schema.Schema{
																Description: `Instruction text for param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"key": &schema.Schema{
																Description: `key
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"not_param": &schema.Schema{
																Description: `Is it not a variable
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"order": &schema.Schema{
																Description: `Order of template param
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"param_array": &schema.Schema{
																Description: `Is it an array
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"parameter_name": &schema.Schema{
																Description: `Name of template param
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"provider": &schema.Schema{
																Description: `provider
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"range": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `UUID of range
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"max_value": &schema.Schema{
																			Description: `Max value of range
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"min_value": &schema.Schema{
																			Description: `Min value of range
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"required": &schema.Schema{
																Description: `Is param required
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"selection": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_selected_values": &schema.Schema{
																			Description: `Default selection values
`,
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"id": &schema.Schema{
																			Description: `UUID of selection
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selection_type": &schema.Schema{
																			Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selection_values": &schema.Schema{
																			Description: `Selection values
`,
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																		},
																	},
																},
															},
														},
													},
												},
												"version": &schema.Schema{
													Description: `Current version of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"create_time": &schema.Schema{
										Description: `Create time of template
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"custom_params_order": &schema.Schema{
										Description: `Custom Params Order
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"document_database": &schema.Schema{
										Description: `Document Database
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"failure_policy": &schema.Schema{
										Description: `Define failure policy if template provisioning fails
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"last_update_time": &schema.Schema{
										Description: `Update time of template
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"latest_version_time": &schema.Schema{
										Description: `Latest versioned template time
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"parent_template_id": &schema.Schema{
										Description: `Parent templateID
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"project_associated": &schema.Schema{
										Description: `Project Associated
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"project_id": &schema.Schema{
										Description: `Project UUID
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"rollback_template_content": &schema.Schema{
										Description: `Rollback template content
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
														},
													},
												},
											},
										},
									},
									"software_type": &schema.Schema{
										Description: `Applicable device software type
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"software_variant": &schema.Schema{
										Description: `Applicable device software variant
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"software_version": &schema.Schema{
										Description: `Applicable device software version
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
															},
														},
													},
												},
											},
										},
									},
									"validation_errors": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rollback_template_errors": &schema.Schema{
													Description: `Validation or design conflicts errors of rollback template
`,
													Type:     schema.TypeList,
													Optional: true,
												},
												"template_errors": &schema.Schema{
													Description: `Validation or design conflicts errors
`,
													Type:     schema.TypeList,
													Optional: true,
												},
												"template_id": &schema.Schema{
													Description: `UUID of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"template_version": &schema.Schema{
													Description: `Current version of template
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceConfigurationTemplateProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestConfigurationTemplateProjectCreateProject(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vProjectID, okProjectID := resourceItem["project_id"]
	vvProjectID := interfaceToString(vProjectID)
	if okProjectID && vvProjectID != "" {
		getResponse2, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["project_id"] = vvProjectID
			d.SetId(joinResourceID(resourceMap))
			return resourceRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.ConfigurationTemplates.GetsAListOfProjects(nil)
		if response2 != nil && err == nil {
			items2 := getAllItemsConfigurationTemplatesGetsAListOfProjects(m, response2, nil)
			item2, err := searchConfigurationTemplatesGetsAListOfProjects(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["project_id"] = vvProjectID
				d.SetId(joinResourceID(resourceMap))
				return resourceRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.CreateProject(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateProject", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateProject", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["project_id"] = vvProjectID
	d.SetId(joinResourceID(resourceMap))
	return resourceRead(ctx, d, m)
}

func resourceConfigurationTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vProjectID, okProjectID := resourceMap["project_id"]

	method1 := []bool{okName, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsAListOfProjects")
		queryParams1 := dnacentersdkgo.GetsAListOfProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsAListOfProjects", err,
				"Failure at GetsAListOfProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO Code Items for DNAC

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetsTheDetailsOfAGivenProject")
		vvProjectID := vProjectID.(string)

		response2, restyResp2, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

	}
	return diags
}

func resourceConfigurationTemplateProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vSortOrder, okSortOrder := resourceMap["sort_order"]
	vProjectID, okProjectID := resourceMap["project_id"]

	method1 := []bool{okName, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}
		//Set value vvName = getResp.
		if getResp.tags != nil {
			vvName = getResp.tags.Name
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestConfigurationTemplateProjectUpdateProject(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ConfigurationTemplates.UpdateProject(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateProject", err, restyResp1.String(),
					"Failure at UpdateProject, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateProject", err,
				"Failure at UpdateProject, unexpected response", ""))
			return diags
		}
	}

	return resourceConfigurationTemplateProjectRead(ctx, d, m)
}

func resourceConfigurationTemplateProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	//TODO

	return diags
}
func expandRequestConfigurationTemplateProjectCreateProject(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProject {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProject{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectCreateProjectTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".templates")))) {
		request.Templates = expandRequestConfigurationTemplateProjectCreateProjectTemplatesArray(ctx, key+".templates", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectCreateProjectTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesArray(ctx, key+".containing_templates", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_params_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_params_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_params_order")))) {
		request.CustomParamsOrder = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateProjectCreateProjectTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".document_database")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".document_database")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".document_database")))) {
		request.DocumentDatabase = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_policy")))) {
		request.FailurePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latest_version_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latest_version_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latest_version_time")))) {
		request.LatestVersionTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_template_id")))) {
		request.ParentTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_associated")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_associated")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_associated")))) {
		request.ProjectAssociated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_content")))) {
		request.RollbackTemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_type")))) {
		request.SoftwareType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_variant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_variant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_variant")))) {
		request.SoftwareVariant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsRollbackTemplateErrorsArray(ctx, key+".rollback_template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsTemplateErrorsArray(ctx, key+".template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_version")))) {
		request.TemplateVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsRollbackTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsRollbackTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectCreateProjectTemplatesValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProject(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProject {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProject{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectUpdateProjectTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".templates")))) {
		request.Templates = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesArray(ctx, key+".templates", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesArray(ctx, key+".containing_templates", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_params_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_params_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_params_order")))) {
		request.CustomParamsOrder = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".document_database")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".document_database")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".document_database")))) {
		request.DocumentDatabase = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_policy")))) {
		request.FailurePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latest_version_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latest_version_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latest_version_time")))) {
		request.LatestVersionTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_template_id")))) {
		request.ParentTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_associated")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_associated")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_associated")))) {
		request.ProjectAssociated = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_content")))) {
		request.RollbackTemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_type")))) {
		request.SoftwareType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_variant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_variant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_variant")))) {
		request.SoftwareVariant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors {
	request := dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsRollbackTemplateErrorsArray(ctx, key+".rollback_template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsTemplateErrorsArray(ctx, key+".template_errors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_version")))) {
		request.TemplateVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsRollbackTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsTemplateErrorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors {
	request := []dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsTemplateErrors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateProjectUpdateProjectTemplatesValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors {
	var request dnacentersdkgo.RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
