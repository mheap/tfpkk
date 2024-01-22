// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RouteDataSource{}
var _ datasource.DataSourceWithConfigure = &RouteDataSource{}

func NewRouteDataSource() datasource.DataSource {
	return &RouteDataSource{}
}

// RouteDataSource is the data source implementation.
type RouteDataSource struct {
	client *sdk.Konnect
}

// RouteDataSourceModel describes the data model.
type RouteDataSourceModel struct {
	ControlPlaneID          types.String    `tfsdk:"control_plane_id"`
	CreatedAt               types.Int64     `tfsdk:"created_at"`
	Destinations            []Destinations  `tfsdk:"destinations"`
	FilterTags              types.String    `tfsdk:"filter_tags"`
	Headers                 *Ordering       `tfsdk:"headers"`
	Hosts                   []types.String  `tfsdk:"hosts"`
	HTTPSRedirectStatusCode types.Int64     `tfsdk:"https_redirect_status_code"`
	ID                      types.String    `tfsdk:"id"`
	Methods                 []types.String  `tfsdk:"methods"`
	Name                    types.String    `tfsdk:"name"`
	Offset                  types.String    `tfsdk:"offset"`
	PathHandling            types.String    `tfsdk:"path_handling"`
	Paths                   []types.String  `tfsdk:"paths"`
	PreserveHost            types.Bool      `tfsdk:"preserve_host"`
	Protocols               []types.String  `tfsdk:"protocols"`
	RegexPriority           types.Int64     `tfsdk:"regex_priority"`
	RequestBuffering        types.Bool      `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool      `tfsdk:"response_buffering"`
	Service                 *PluginConsumer `tfsdk:"service"`
	Size                    types.Int64     `tfsdk:"size"`
	Snis                    []types.String  `tfsdk:"snis"`
	Sources                 []Destinations  `tfsdk:"sources"`
	StripPath               types.Bool      `tfsdk:"strip_path"`
	Tags                    []types.String  `tfsdk:"tags"`
	UpdatedAt               types.Int64     `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *RouteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

// Schema defines the schema for the data source.
func (r *RouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Route DataSource",

		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of your control plane. This variable is available in the Konnect manager`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"destinations": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
						},
						"port": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
				Description: `A list of IP destinations of incoming connections that match this route when using stream routing. Each entry is an object with fields ip (optionally in CIDR range notation) and/or port.`,
			},
			"filter_tags": schema.StringAttribute{
				Optional:    true,
				Description: `A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.`,
			},
			"headers": schema.SingleNestedAttribute{
				Computed:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `One or more lists of values indexed by header name that will cause this route to match if present in the request. The ` + "`" + `Host` + "`" + ` header cannot be used with this attribute: hosts should be specified using the ` + "`" + `hosts` + "`" + ` attribute. When ` + "`" + `headers` + "`" + ` contains only one value and that value starts with the special prefix ` + "`" + `~*` + "`" + `, the value is interpreted as a regular expression.`,
			},
			"hosts": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of domain names that match this route. Note that the hosts value is case sensitive.`,
			},
			"https_redirect_status_code": schema.Int64Attribute{
				Computed:    true,
				Description: `The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is ` + "`" + `HTTP` + "`" + ` instead of ` + "`" + `HTTPS` + "`" + `. ` + "`" + `Location` + "`" + ` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the ` + "`" + `https` + "`" + ` protocol. Default: 426`,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"methods": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of HTTP methods that match this route.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.`,
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Description: `Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources`,
			},
			"path_handling": schema.StringAttribute{
				Computed:    true,
				Description: `Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior. Default: "v0"`,
			},
			"paths": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of paths that match this route.`,
			},
			"preserve_host": schema.BoolAttribute{
				Computed:    true,
				Description: `When matching a route via one of the ` + "`" + `hosts` + "`" + ` domain names, use the request ` + "`" + `Host` + "`" + ` header in the upstream request headers. If set to ` + "`" + `false` + "`" + `, the upstream ` + "`" + `Host` + "`" + ` header will be that of the service's ` + "`" + `host` + "`" + `. Default: false`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only ` + "`" + `https` + "`" + `, HTTP requests are answered with an upgrade error. When set to only ` + "`" + `http` + "`" + `, HTTPS requests are answered with an error.`,
			},
			"regex_priority": schema.Int64Attribute{
				Computed:    true,
				Description: `A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same ` + "`" + `regex_priority` + "`" + `, the older one (lowest ` + "`" + `created_at` + "`" + `) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones). Default: 0`,
			},
			"request_buffering": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding. Default: true`,
			},
			"response_buffering": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding. Default: true`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `The service this route is associated to. This is where the route proxies traffic to.`,
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of resources to be returned. Default: 100`,
			},
			"snis": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of SNIs that match this route when using stream routing.`,
			},
			"sources": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed: true,
						},
						"port": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
				Description: `A list of IP sources of incoming connections that match this route when using stream routing. Each entry is an object with fields ip (optionally in CIDR range notation) and/or port.`,
			},
			"strip_path": schema.BoolAttribute{
				Computed:    true,
				Description: `When matching a route via one of the ` + "`" + `paths` + "`" + `, strip the matching prefix from the upstream request URL. Default: true`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the route for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *RouteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RouteDataSourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	filterTags := new(string)
	if !data.FilterTags.IsUnknown() && !data.FilterTags.IsNull() {
		*filterTags = data.FilterTags.ValueString()
	} else {
		filterTags = nil
	}
	offset := new(string)
	if !data.Offset.IsUnknown() && !data.Offset.IsNull() {
		*offset = data.Offset.ValueString()
	} else {
		offset = nil
	}
	routeID := data.ID.ValueString()
	size := new(int64)
	if !data.Size.IsUnknown() && !data.Size.IsNull() {
		*size = data.Size.ValueInt64()
	} else {
		size = nil
	}
	request := operations.GetRouteRequest{
		ControlPlaneID: controlPlaneID,
		FilterTags:     filterTags,
		Offset:         offset,
		RouteID:        routeID,
		Size:           size,
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
	data.RefreshFromSharedRoute(res.Route)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
