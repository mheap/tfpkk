// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Unhealthy struct {
	HTTPFailures types.Int64   `tfsdk:"http_failures"`
	HTTPStatuses []types.Int64 `tfsdk:"http_statuses"`
	Interval     types.Number  `tfsdk:"interval"`
	TCPFailures  types.Int64   `tfsdk:"tcp_failures"`
	Timeouts     types.Int64   `tfsdk:"timeouts"`
}
