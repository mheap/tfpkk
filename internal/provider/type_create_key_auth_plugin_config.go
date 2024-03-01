// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateKeyAuthPluginConfig struct {
	Anonymous       types.String   `tfsdk:"anonymous"`
	HideCredentials types.Bool     `tfsdk:"hide_credentials"`
	KeyInBody       types.Bool     `tfsdk:"key_in_body"`
	KeyInHeader     types.Bool     `tfsdk:"key_in_header"`
	KeyInQuery      types.Bool     `tfsdk:"key_in_query"`
	KeyNames        []types.String `tfsdk:"key_names"`
	RunOnPreflight  types.Bool     `tfsdk:"run_on_preflight"`
}
