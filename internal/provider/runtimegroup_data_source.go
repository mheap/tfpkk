// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"konnect/internal/sdk"
	"konnect/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RuntimeGroupDataSource{}
var _ datasource.DataSourceWithConfigure = &RuntimeGroupDataSource{}

func NewRuntimeGroupDataSource() datasource.DataSource {
	return &RuntimeGroupDataSource{}
}

// RuntimeGroupDataSource is the data source implementation.
type RuntimeGroupDataSource struct {
	client *sdk.Konnect
}

// RuntimeGroupDataSourceModel describes the data model.
type RuntimeGroupDataSourceModel struct {
	Config      *RuntimeGroupConfig `tfsdk:"config"`
	Description types.String        `tfsdk:"description"`
	ID          types.String        `tfsdk:"id"`
	Labels      types.String        `tfsdk:"labels"`
	Name        types.String        `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (r *RuntimeGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_runtime_group"
}

// Schema defines the schema for the data source.
func (r *RuntimeGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RuntimeGroup DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `Control Plane Endpoint.`,
					},
					"telemetry_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `Telemetry Endpoint.`,
					},
				},
				Description: `CP configuration object for related access endpoints.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the runtime group in Konnect.`,
			},
			"id": schema.StringAttribute{
				Optional:    true,
				Description: `The runtime group ID`,
			},
			"labels": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
				MarkdownDescription: `Parsed as JSON.` + "\n" +
					`Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the runtime group.`,
			},
		},
	}
}

func (r *RuntimeGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RuntimeGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RuntimeGroupDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	request := operations.GetRuntimeGroupRequest{
		ID: id,
	}
	res, err := r.client.RuntimeGroups.GetRuntimeGroup(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.RuntimeGroup == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.RuntimeGroup)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
