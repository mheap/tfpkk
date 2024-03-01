// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
)

func (r *PluginApplicationRegistrationDataSourceModel) RefreshFromSharedApplicationRegistrationPlugin(resp *shared.ApplicationRegistrationPlugin) {
	r.Config.AutoApprove = types.BoolPointerValue(resp.Config.AutoApprove)
	r.Config.Description = types.StringPointerValue(resp.Config.Description)
	r.Config.DisplayName = types.StringPointerValue(resp.Config.DisplayName)
	r.Config.ShowIssuer = types.BoolPointerValue(resp.Config.ShowIssuer)
	if resp.Consumer == nil {
		r.Consumer = nil
	} else {
		r.Consumer = &CreateACLConsumer{}
		r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
	}
	r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
	r.Enabled = types.BoolPointerValue(resp.Enabled)
	r.ID = types.StringPointerValue(resp.ID)
	r.Protocols = nil
	for _, v := range resp.Protocols {
		r.Protocols = append(r.Protocols, types.StringValue(string(v)))
	}
	if resp.Route == nil {
		r.Route = nil
	} else {
		r.Route = &CreateACLConsumer{}
		r.Route.ID = types.StringPointerValue(resp.Route.ID)
	}
	if resp.Service == nil {
		r.Service = nil
	} else {
		r.Service = &CreateACLConsumer{}
		r.Service.ID = types.StringPointerValue(resp.Service.ID)
	}
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
}
