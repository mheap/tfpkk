// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_objectplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayControlPlaneResource{}
var _ resource.ResourceWithImportState = &GatewayControlPlaneResource{}

func NewGatewayControlPlaneResource() resource.Resource {
	return &GatewayControlPlaneResource{}
}

// GatewayControlPlaneResource defines the resource implementation.
type GatewayControlPlaneResource struct {
	client *sdk.Konnect
}

// GatewayControlPlaneResourceModel describes the resource data model.
type GatewayControlPlaneResourceModel struct {
	AuthType     types.String        `tfsdk:"auth_type"`
	CloudGateway types.Bool          `tfsdk:"cloud_gateway"`
	ClusterType  types.String        `tfsdk:"cluster_type"`
	Config       *ControlPlaneConfig `tfsdk:"config"`
	Description  types.String        `tfsdk:"description"`
	ID           types.String        `tfsdk:"id"`
	Labels       types.String        `tfsdk:"labels"`
	Name         types.String        `tfsdk:"name"`
}

func (r *GatewayControlPlaneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_control_plane"
}

func (r *GatewayControlPlaneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayControlPlane Resource",

		Attributes: map[string]schema.Attribute{
			"auth_type": schema.StringAttribute{
				Optional:    true,
				Description: `The auth type value of the cluster associated with the Runtime Group. must be one of ["pinned_client_certs", "pki_client_certs"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"pinned_client_certs",
						"pki_client_certs",
					),
				},
			},
			"cloud_gateway": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Whether this control-plane can be used for cloud-gateways. Requires replacement if changed. `,
			},
			"cluster_type": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The ClusterType value of the cluster associated with the Control Plane. Requires replacement if changed. ; must be one of ["CLUSTER_TYPE_HYBRID", "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER", "CLUSTER_TYPE_CONTROL_PLANE_GROUP"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"CLUSTER_TYPE_HYBRID",
						"CLUSTER_TYPE_K8S_INGRESS_CONTROLLER",
						"CLUSTER_TYPE_CONTROL_PLANE_GROUP",
					),
				},
			},
			"config": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Control Plane Endpoint.`,
					},
					"telemetry_endpoint": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Telemetry Endpoint.`,
					},
				},
				Description: `CP configuration object for related access endpoints.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The description of the control plane in Konnect.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The control plane ID`,
			},
			"labels": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Labels to facilitate tagged search on control planes. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'. Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the control plane.`,
			},
		},
	}
}

func (r *GatewayControlPlaneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayControlPlaneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayControlPlaneResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	authType := new(shared.AuthType)
	if !data.AuthType.IsUnknown() && !data.AuthType.IsNull() {
		*authType = shared.AuthType(data.AuthType.ValueString())
	} else {
		authType = nil
	}
	cloudGateway := new(bool)
	if !data.CloudGateway.IsUnknown() && !data.CloudGateway.IsNull() {
		*cloudGateway = data.CloudGateway.ValueBool()
	} else {
		cloudGateway = nil
	}
	clusterType := new(shared.ClusterType)
	if !data.ClusterType.IsUnknown() && !data.ClusterType.IsNull() {
		*clusterType = shared.ClusterType(data.ClusterType.ValueString())
	} else {
		clusterType = nil
	}
	description := new(string)
	if !data.Description.IsUnknown() && !data.Description.IsNull() {
		*description = data.Description.ValueString()
	} else {
		description = nil
	}
	var labels interface{}
	if !data.Labels.IsUnknown() && !data.Labels.IsNull() {
		_ = json.Unmarshal([]byte(data.Labels.ValueString()), &labels)
	}
	name := data.Name.ValueString()
	request := shared.CreateControlPlaneRequest{
		AuthType:     authType,
		CloudGateway: cloudGateway,
		ClusterType:  clusterType,
		Description:  description,
		Labels:       labels,
		Name:         name,
	}
	res, err := r.client.ControlPlanes.CreateControlPlane(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res.ControlPlane)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id := data.ID.ValueString()
	request1 := operations.GetControlPlaneRequest{
		ID: id,
	}
	res1, err := r.client.ControlPlanes.GetControlPlane(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res1.ControlPlane)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayControlPlaneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayControlPlaneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	request := operations.GetControlPlaneRequest{
		ID: id,
	}
	res, err := r.client.ControlPlanes.GetControlPlane(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res.ControlPlane)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayControlPlaneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayControlPlaneResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	updateControlPlaneRequest := *data.ToSharedUpdateControlPlaneRequest()
	id := data.ID.ValueString()
	request := operations.UpdateControlPlaneRequest{
		UpdateControlPlaneRequest: updateControlPlaneRequest,
		ID:                        id,
	}
	res, err := r.client.ControlPlanes.UpdateControlPlane(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res.ControlPlane)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id1 := data.ID.ValueString()
	request1 := operations.GetControlPlaneRequest{
		ID: id1,
	}
	res1, err := r.client.ControlPlanes.GetControlPlane(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res1.ControlPlane)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayControlPlaneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayControlPlaneResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	request := operations.DeleteControlPlaneRequest{
		ID: id,
	}
	res, err := r.client.ControlPlanes.DeleteControlPlane(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayControlPlaneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
