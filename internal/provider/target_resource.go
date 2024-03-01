// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TargetResource{}
var _ resource.ResourceWithImportState = &TargetResource{}

func NewTargetResource() resource.Resource {
	return &TargetResource{}
}

// TargetResource defines the resource implementation.
type TargetResource struct {
	client *sdk.Konnect
}

// TargetResourceModel describes the resource data model.
type TargetResourceModel struct {
	ControlPlaneID types.String       `tfsdk:"control_plane_id"`
	CreatedAt      types.Number       `tfsdk:"created_at"`
	ID             types.String       `tfsdk:"id"`
	Tags           []types.String     `tfsdk:"tags"`
	Target         types.String       `tfsdk:"target"`
	Upstream       *CreateACLConsumer `tfsdk:"upstream"`
	Weight         types.Int64        `tfsdk:"weight"`
}

func (r *TargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_target"
}

func (r *TargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Target Resource",

		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.NumberAttribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the Upstream to lookup`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Target for grouping and filtering.`,
			},
			"target": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the ` + "`" + `port` + "`" + ` value will be overridden by the value from the DNS record.`,
			},
			"upstream": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: `The weight this target gets within the upstream loadbalancer (` + "`" + `0` + "`" + `-` + "`" + `65535` + "`" + `). If the hostname resolves to an SRV record, the ` + "`" + `weight` + "`" + ` value will be overridden by the value from the DNS record. Default: 100`,
			},
		},
	}
}

func (r *TargetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TargetResourceModel
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

	createTargetWithoutParents := *data.ToSharedCreateTargetWithoutParents()
	upstreamID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.CreateTargetRequest{
		CreateTargetWithoutParents: createTargetWithoutParents,
		UpstreamID:                 upstreamID,
		ControlPlaneID:             controlPlaneID,
	}
	res, err := r.client.Targets.CreateTarget(ctx, request)
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
	if res.Target == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TargetResourceModel
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

	targetID := data.ID.ValueString()
	upstreamID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.ReadTargetRequest{
		TargetID:       targetID,
		UpstreamID:     upstreamID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Targets.ReadTarget(ctx, request)
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
	if res.Target == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TargetResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	createTargetWithoutParents := *data.ToSharedCreateTargetWithoutParents()
	targetID := data.ID.ValueString()
	upstreamID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.UpdateTargetRequest{
		CreateTargetWithoutParents: createTargetWithoutParents,
		TargetID:                   targetID,
		UpstreamID:                 upstreamID,
		ControlPlaneID:             controlPlaneID,
	}
	res, err := r.client.Targets.UpdateTarget(ctx, request)
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
	if res.Target == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TargetResourceModel
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

	targetID := data.ID.ValueString()
	upstreamID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.DeleteTargetRequest{
		TargetID:       targetID,
		UpstreamID:     upstreamID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Targets.DeleteTarget(ctx, request)
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

func (r *TargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource target. Reason: composite imports strings not supported.")
}
