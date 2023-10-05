// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"konnect/internal/sdk"
	"konnect/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RouteResource{}
var _ resource.ResourceWithImportState = &RouteResource{}

func NewRouteResource() resource.Resource {
	return &RouteResource{}
}

// RouteResource defines the resource implementation.
type RouteResource struct {
	client *sdk.Konnect
}

// RouteResourceModel describes the resource data model.
type RouteResourceModel struct {
	CreatedAt               types.Int64    `tfsdk:"created_at"`
	Headers                 *RouteHeaders  `tfsdk:"headers"`
	Hosts                   []types.String `tfsdk:"hosts"`
	HTTPSRedirectStatusCode types.Int64    `tfsdk:"https_redirect_status_code"`
	ID                      types.String   `tfsdk:"id"`
	Methods                 []types.String `tfsdk:"methods"`
	Name                    types.String   `tfsdk:"name"`
	PathHandling            types.String   `tfsdk:"path_handling"`
	Paths                   []types.String `tfsdk:"paths"`
	PreserveHost            types.Bool     `tfsdk:"preserve_host"`
	Protocols               []types.String `tfsdk:"protocols"`
	RegexPriority           types.Int64    `tfsdk:"regex_priority"`
	RequestBuffering        types.Bool     `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool     `tfsdk:"response_buffering"`
	RuntimeGroupID          types.String   `tfsdk:"runtime_group_id"`
	Service                 *RouteService  `tfsdk:"service"`
	Snis                    []types.String `tfsdk:"snis"`
	StripPath               types.Bool     `tfsdk:"strip_path"`
	Tags                    []types.String `tfsdk:"tags"`
	UpdatedAt               types.Int64    `tfsdk:"updated_at"`
}

func (r *RouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

func (r *RouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Route Resource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"headers": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `One or more lists of values indexed by header name that will cause this route to match if present in the request. The ` + "`" + `Host` + "`" + ` header cannot be used with this attribute: hosts should be specified using the ` + "`" + `hosts` + "`" + ` attribute. When ` + "`" + `headers` + "`" + ` contains only one value and that value starts with the special prefix ` + "`" + `~*` + "`" + `, the value is interpreted as a regular expression.`,
			},
			"hosts": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of domain names that match this route. Note that the hosts value is case sensitive.`,
			},
			"https_redirect_status_code": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is ` + "`" + `HTTP` + "`" + ` instead of ` + "`" + `HTTPS` + "`" + `. ` + "`" + `Location` + "`" + ` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the ` + "`" + `https` + "`" + ` protocol.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
			},
			"methods": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of HTTP methods that match this route.`,
			},
			"name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.`,
			},
			"path_handling": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.`,
			},
			"paths": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of paths that match this route.`,
			},
			"preserve_host": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `When matching a route via one of the ` + "`" + `hosts` + "`" + ` domain names, use the request ` + "`" + `Host` + "`" + ` header in the upstream request headers. If set to ` + "`" + `false` + "`" + `, the upstream ` + "`" + `Host` + "`" + ` header will be that of the service's ` + "`" + `host` + "`" + `.`,
			},
			"protocols": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only ` + "`" + `https` + "`" + `, HTTP requests are answered with an upgrade error. When set to only ` + "`" + `http` + "`" + `, HTTPS requests are answered with an error.`,
			},
			"regex_priority": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same ` + "`" + `regex_priority` + "`" + `, the older one (lowest ` + "`" + `created_at` + "`" + `) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).`,
			},
			"request_buffering": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.`,
			},
			"response_buffering": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.`,
			},
			"runtime_group_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `The ID of your runtime group. This variable is available in the Konnect manager`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
				},
				Description: `The service this route is associated to. This is where the route proxies traffic to.`,
			},
			"snis": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of SNIs that match this route when using stream routing.`,
			},
			"strip_path": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `When matching a route via one of the ` + "`" + `paths` + "`" + `, strip the matching prefix from the upstream request URL.`,
			},
			"tags": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the route for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *RouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RouteResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	route := data.ToCreateSDKType()
	runtimeGroupID := data.RuntimeGroupID.ValueString()
	request := operations.CreateRouteRequest{
		Route:          route,
		RuntimeGroupID: runtimeGroupID,
	}
	res, err := r.client.Routes.CreateRoute(ctx, request)
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
	if res.Route == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.Route)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RouteResourceModel
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

	routeID := data.ID.ValueString()
	runtimeGroupID := data.RuntimeGroupID.ValueString()
	request := operations.GetRouteRequest{
		RouteID:        routeID,
		RuntimeGroupID: runtimeGroupID,
	}
	res, err := r.client.Routes.GetRoute(ctx, request)
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
	if res.Route == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Route)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RouteResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RouteResourceModel
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

	routeID := data.ID.ValueString()
	runtimeGroupID := data.RuntimeGroupID.ValueString()
	request := operations.DeleteRouteRequest{
		RouteID:        routeID,
		RuntimeGroupID: runtimeGroupID,
	}
	res, err := r.client.Routes.DeleteRoute(ctx, request)
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

func (r *RouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource route.")
}