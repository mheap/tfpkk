// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"konnect/internal/sdk/pkg/models/shared"
	"time"
)

func (r *APIProductVersionResourceModel) ToCreateSDKType() *shared.CreateAPIProductVersionDTO {
	deprecated := new(bool)
	if !r.Deprecated.IsUnknown() && !r.Deprecated.IsNull() {
		*deprecated = r.Deprecated.ValueBool()
	} else {
		deprecated = nil
	}
	var gatewayService *shared.GatewayServicePayload
	controlPlaneID := r.GatewayService.ControlPlaneID.ValueString()
	id := r.GatewayService.ID.ValueString()
	gatewayService = &shared.GatewayServicePayload{
		ControlPlaneID: controlPlaneID,
		ID:             id,
	}
	name := r.Name.ValueString()
	publishStatus := new(shared.CreateAPIProductVersionDTOPublishStatus)
	if !r.PublishStatus.IsUnknown() && !r.PublishStatus.IsNull() {
		*publishStatus = shared.CreateAPIProductVersionDTOPublishStatus(r.PublishStatus.ValueString())
	} else {
		publishStatus = nil
	}
	out := shared.CreateAPIProductVersionDTO{
		Deprecated:     deprecated,
		GatewayService: gatewayService,
		Name:           name,
		PublishStatus:  publishStatus,
	}
	return &out
}

func (r *APIProductVersionResourceModel) ToGetSDKType() *shared.CreateAPIProductVersionDTO {
	out := r.ToCreateSDKType()
	return out
}

func (r *APIProductVersionResourceModel) ToUpdateSDKType() *shared.UpdateAPIProductVersionDTO {
	deprecated := new(bool)
	if !r.Deprecated.IsUnknown() && !r.Deprecated.IsNull() {
		*deprecated = r.Deprecated.ValueBool()
	} else {
		deprecated = nil
	}
	var gatewayService *shared.GatewayServicePayload
	controlPlaneID := r.GatewayService.ControlPlaneID.ValueString()
	id := r.GatewayService.ID.ValueString()
	gatewayService = &shared.GatewayServicePayload{
		ControlPlaneID: controlPlaneID,
		ID:             id,
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	publishStatus := new(shared.UpdateAPIProductVersionDTOPublishStatus)
	if !r.PublishStatus.IsUnknown() && !r.PublishStatus.IsNull() {
		*publishStatus = shared.UpdateAPIProductVersionDTOPublishStatus(r.PublishStatus.ValueString())
	} else {
		publishStatus = nil
	}
	out := shared.UpdateAPIProductVersionDTO{
		Deprecated:     deprecated,
		GatewayService: gatewayService,
		Name:           name,
		PublishStatus:  publishStatus,
	}
	return &out
}

func (r *APIProductVersionResourceModel) ToDeleteSDKType() *shared.CreateAPIProductVersionDTO {
	out := r.ToCreateSDKType()
	return out
}

func (r *APIProductVersionResourceModel) RefreshFromGetResponse(resp *shared.APIProductVersion) {
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.Deprecated != nil {
		r.Deprecated = types.BoolValue(*resp.Deprecated)
	} else {
		r.Deprecated = types.BoolNull()
	}
	r.GatewayService.ControlPlaneID = types.StringValue(resp.GatewayService.ControlPlaneID)
	r.GatewayService.ID = types.StringValue(resp.GatewayService.ID)
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.PublishStatus = types.StringValue(string(resp.PublishStatus))
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	} else {
		r.UpdatedAt = types.StringNull()
	}
}

func (r *APIProductVersionResourceModel) RefreshFromCreateResponse(resp *shared.APIProductVersion) {
	r.RefreshFromGetResponse(resp)
}

func (r *APIProductVersionResourceModel) RefreshFromUpdateResponse(resp *shared.APIProductVersion) {
	r.RefreshFromGetResponse(resp)
}
