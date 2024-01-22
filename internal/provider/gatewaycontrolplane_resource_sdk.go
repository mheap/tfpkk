// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
)

func (r *GatewayControlPlaneResourceModel) RefreshFromSharedControlPlane(resp *shared.ControlPlane) {
	if resp.Config == nil {
		r.Config = nil
	} else {
		r.Config = &ControlPlaneConfig{}
		r.Config.ControlPlaneEndpoint = types.StringPointerValue(resp.Config.ControlPlaneEndpoint)
		r.Config.TelemetryEndpoint = types.StringPointerValue(resp.Config.TelemetryEndpoint)
	}
	r.Description = types.StringPointerValue(resp.Description)
	r.ID = types.StringPointerValue(resp.ID)
	if resp.Labels == nil {
		r.Labels = types.StringNull()
	} else {
		labelsResult, _ := json.Marshal(resp.Labels)
		r.Labels = types.StringValue(string(labelsResult))
	}
	r.Name = types.StringPointerValue(resp.Name)
}

func (r *GatewayControlPlaneResourceModel) ToSharedUpdateControlPlaneRequest() *shared.UpdateControlPlaneRequest {
	authType := new(shared.UpdateControlPlaneRequestAuthType)
	if !r.AuthType.IsUnknown() && !r.AuthType.IsNull() {
		*authType = shared.UpdateControlPlaneRequestAuthType(r.AuthType.ValueString())
	} else {
		authType = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	var labels interface{}
	if !r.Labels.IsUnknown() && !r.Labels.IsNull() {
		_ = json.Unmarshal([]byte(r.Labels.ValueString()), &labels)
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	out := shared.UpdateControlPlaneRequest{
		AuthType:    authType,
		Description: description,
		Labels:      labels,
		Name:        name,
	}
	return &out
}
