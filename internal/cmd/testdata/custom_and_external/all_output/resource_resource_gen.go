// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package generated

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/my_account/my_project/boolplanmodifier"
	"github.com/my_account/my_project/myboolvalidator"
	"github.com/my_account_my_project/bool"
)

var resourceResourceSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"bool_attribute": schema.BoolAttribute{
			CustomType: my_bool_type,
			Computed:   true,
			PlanModifiers: []planmodifier.Bool{
				myboolplanmodifier.Modify(),
			},
			Validators: []validator.Bool{
				myboolvalidator.Validate(),
			},
			Default: booldefault.StaticBool(true),
		},
	},
}

type ResourceModel struct {
	BoolAttribute my_bool_value `tfsdk:"bool_attribute"`
}
