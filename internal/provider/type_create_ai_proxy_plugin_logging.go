// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateAIProxyPluginLogging struct {
	LogPayloads   types.Bool `tfsdk:"log_payloads"`
	LogStatistics types.Bool `tfsdk:"log_statistics"`
}
