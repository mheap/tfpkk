// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"konnect/internal/sdk/pkg/models/shared"
)

func (r *RuntimeGroupDataSourceModel) RefreshFromGetResponse(resp *shared.RuntimeGroup) {
	if r.Config == nil {
		r.Config = &RuntimeGroupConfig{}
	}
	if resp.Config == nil {
		r.Config = nil
	} else {
		r.Config = &RuntimeGroupConfig{}
		if resp.Config.ControlPlaneEndpoint != nil {
			r.Config.ControlPlaneEndpoint = types.StringValue(*resp.Config.ControlPlaneEndpoint)
		} else {
			r.Config.ControlPlaneEndpoint = types.StringNull()
		}
		if resp.Config.TelemetryEndpoint != nil {
			r.Config.TelemetryEndpoint = types.StringValue(*resp.Config.TelemetryEndpoint)
		} else {
			r.Config.TelemetryEndpoint = types.StringNull()
		}
	}
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.Labels == nil {
		r.Labels = types.StringNull()
	} else {
		labelsResult, _ := json.Marshal(resp.Labels)
		r.Labels = types.StringValue(string(labelsResult))
	}
	if resp.Name != nil {
		r.Name = types.StringValue(*resp.Name)
	} else {
		r.Name = types.StringNull()
	}
}
