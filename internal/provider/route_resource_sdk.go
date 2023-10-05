// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"konnect/internal/sdk/pkg/models/shared"
)

func (r *RouteResourceModel) ToCreateSDKType() *shared.Route {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	var headers *shared.RouteHeaders
	if r.Headers != nil {
		headers = &shared.RouteHeaders{}
	}
	var hosts []string = nil
	for _, hostsItem := range r.Hosts {
		hosts = append(hosts, hostsItem.ValueString())
	}
	httpsRedirectStatusCode := new(int64)
	if !r.HTTPSRedirectStatusCode.IsUnknown() && !r.HTTPSRedirectStatusCode.IsNull() {
		*httpsRedirectStatusCode = r.HTTPSRedirectStatusCode.ValueInt64()
	} else {
		httpsRedirectStatusCode = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var methods []string = nil
	for _, methodsItem := range r.Methods {
		methods = append(methods, methodsItem.ValueString())
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	pathHandling := new(string)
	if !r.PathHandling.IsUnknown() && !r.PathHandling.IsNull() {
		*pathHandling = r.PathHandling.ValueString()
	} else {
		pathHandling = nil
	}
	var paths []string = nil
	for _, pathsItem := range r.Paths {
		paths = append(paths, pathsItem.ValueString())
	}
	preserveHost := new(bool)
	if !r.PreserveHost.IsUnknown() && !r.PreserveHost.IsNull() {
		*preserveHost = r.PreserveHost.ValueBool()
	} else {
		preserveHost = nil
	}
	var protocols []string = nil
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, protocolsItem.ValueString())
	}
	regexPriority := new(int64)
	if !r.RegexPriority.IsUnknown() && !r.RegexPriority.IsNull() {
		*regexPriority = r.RegexPriority.ValueInt64()
	} else {
		regexPriority = nil
	}
	requestBuffering := new(bool)
	if !r.RequestBuffering.IsUnknown() && !r.RequestBuffering.IsNull() {
		*requestBuffering = r.RequestBuffering.ValueBool()
	} else {
		requestBuffering = nil
	}
	responseBuffering := new(bool)
	if !r.ResponseBuffering.IsUnknown() && !r.ResponseBuffering.IsNull() {
		*responseBuffering = r.ResponseBuffering.ValueBool()
	} else {
		responseBuffering = nil
	}
	var service *shared.RouteService
	if r.Service != nil {
		id1 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id1 = r.Service.ID.ValueString()
		} else {
			id1 = nil
		}
		service = &shared.RouteService{
			ID: id1,
		}
	}
	var snis []string = nil
	for _, snisItem := range r.Snis {
		snis = append(snis, snisItem.ValueString())
	}
	stripPath := new(bool)
	if !r.StripPath.IsUnknown() && !r.StripPath.IsNull() {
		*stripPath = r.StripPath.ValueBool()
	} else {
		stripPath = nil
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	out := shared.Route{
		CreatedAt:               createdAt,
		Headers:                 headers,
		Hosts:                   hosts,
		HTTPSRedirectStatusCode: httpsRedirectStatusCode,
		ID:                      id,
		Methods:                 methods,
		Name:                    name,
		PathHandling:            pathHandling,
		Paths:                   paths,
		PreserveHost:            preserveHost,
		Protocols:               protocols,
		RegexPriority:           regexPriority,
		RequestBuffering:        requestBuffering,
		ResponseBuffering:       responseBuffering,
		Service:                 service,
		Snis:                    snis,
		StripPath:               stripPath,
		Tags:                    tags,
		UpdatedAt:               updatedAt,
	}
	return &out
}

func (r *RouteResourceModel) ToGetSDKType() *shared.Route {
	out := r.ToCreateSDKType()
	return out
}

func (r *RouteResourceModel) ToDeleteSDKType() *shared.Route {
	out := r.ToCreateSDKType()
	return out
}

func (r *RouteResourceModel) RefreshFromGetResponse(resp *shared.Route) {
	if resp.CreatedAt != nil {
		r.CreatedAt = types.Int64Value(*resp.CreatedAt)
	} else {
		r.CreatedAt = types.Int64Null()
	}
	if r.Headers == nil {
		r.Headers = &RouteHeaders{}
	}
	if resp.Headers == nil {
		r.Headers = nil
	} else {
		r.Headers = &RouteHeaders{}
	}
	r.Hosts = nil
	for _, v := range resp.Hosts {
		r.Hosts = append(r.Hosts, types.StringValue(v))
	}
	if resp.HTTPSRedirectStatusCode != nil {
		r.HTTPSRedirectStatusCode = types.Int64Value(*resp.HTTPSRedirectStatusCode)
	} else {
		r.HTTPSRedirectStatusCode = types.Int64Null()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	r.Methods = nil
	for _, v := range resp.Methods {
		r.Methods = append(r.Methods, types.StringValue(v))
	}
	if resp.Name != nil {
		r.Name = types.StringValue(*resp.Name)
	} else {
		r.Name = types.StringNull()
	}
	if resp.PathHandling != nil {
		r.PathHandling = types.StringValue(*resp.PathHandling)
	} else {
		r.PathHandling = types.StringNull()
	}
	r.Paths = nil
	for _, v := range resp.Paths {
		r.Paths = append(r.Paths, types.StringValue(v))
	}
	if resp.PreserveHost != nil {
		r.PreserveHost = types.BoolValue(*resp.PreserveHost)
	} else {
		r.PreserveHost = types.BoolNull()
	}
	r.Protocols = nil
	for _, v := range resp.Protocols {
		r.Protocols = append(r.Protocols, types.StringValue(v))
	}
	if resp.RegexPriority != nil {
		r.RegexPriority = types.Int64Value(*resp.RegexPriority)
	} else {
		r.RegexPriority = types.Int64Null()
	}
	if resp.RequestBuffering != nil {
		r.RequestBuffering = types.BoolValue(*resp.RequestBuffering)
	} else {
		r.RequestBuffering = types.BoolNull()
	}
	if resp.ResponseBuffering != nil {
		r.ResponseBuffering = types.BoolValue(*resp.ResponseBuffering)
	} else {
		r.ResponseBuffering = types.BoolNull()
	}
	if r.Service == nil {
		r.Service = &RouteService{}
	}
	if resp.Service == nil {
		r.Service = nil
	} else {
		r.Service = &RouteService{}
		if resp.Service.ID != nil {
			r.Service.ID = types.StringValue(*resp.Service.ID)
		} else {
			r.Service.ID = types.StringNull()
		}
	}
	r.Snis = nil
	for _, v := range resp.Snis {
		r.Snis = append(r.Snis, types.StringValue(v))
	}
	if resp.StripPath != nil {
		r.StripPath = types.BoolValue(*resp.StripPath)
	} else {
		r.StripPath = types.BoolNull()
	}
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.Int64Value(*resp.UpdatedAt)
	} else {
		r.UpdatedAt = types.Int64Null()
	}
}

func (r *RouteResourceModel) RefreshFromCreateResponse(resp *shared.Route) {
	r.RefreshFromGetResponse(resp)
}
