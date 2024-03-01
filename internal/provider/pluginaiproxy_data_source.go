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
var _ datasource.DataSource = &PluginAIProxyDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginAIProxyDataSource{}

func NewPluginAIProxyDataSource() datasource.DataSource {
	return &PluginAIProxyDataSource{}
}

// PluginAIProxyDataSource is the data source implementation.
type PluginAIProxyDataSource struct {
	client *sdk.Konnect
}

// PluginAIProxyDataSourceModel describes the data model.
type PluginAIProxyDataSourceModel struct {
	Config         CreateAIProxyPluginConfig `tfsdk:"config"`
	Consumer       *CreateACLConsumer        `tfsdk:"consumer"`
	ControlPlaneID types.String              `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64               `tfsdk:"created_at"`
	Enabled        types.Bool                `tfsdk:"enabled"`
	ID             types.String              `tfsdk:"id"`
	Protocols      []types.String            `tfsdk:"protocols"`
	Route          *CreateACLConsumer        `tfsdk:"route"`
	Service        *CreateACLConsumer        `tfsdk:"service"`
	Tags           []types.String            `tfsdk:"tags"`
}

// Metadata returns the data source type name.
func (r *PluginAIProxyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_ai_proxy"
}

// Schema defines the schema for the data source.
func (r *PluginAIProxyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginAIProxy DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"auth": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"header_name": schema.StringAttribute{
								Computed:    true,
								Description: `If AI model requires authentication via Authorization or API key header, specify its name here.`,
							},
							"header_value": schema.StringAttribute{
								Computed:    true,
								Description: `Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.`,
							},
							"param_location": schema.StringAttribute{
								Computed:    true,
								Description: `Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body. must be one of ["query", "body"]`,
							},
							"param_name": schema.StringAttribute{
								Computed:    true,
								Description: `If AI model requires authentication via query parameter, specify its name here.`,
							},
							"param_value": schema.StringAttribute{
								Computed:    true,
								Description: `Specify the full parameter value for 'param_name'.`,
							},
						},
					},
					"logging": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"log_payloads": schema.BoolAttribute{
								Computed:    true,
								Description: `If enabled, will log the request and response body into the Kong log plugin(s) output.`,
							},
							"log_statistics": schema.BoolAttribute{
								Computed:    true,
								Description: `If enabled and supported by the driver, will add model usage and token metrics into the Kong log plugin(s) output.`,
							},
						},
					},
					"model": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Computed:    true,
								Description: `Model name to execute.`,
							},
							"options": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"anthropic_version": schema.StringAttribute{
										Computed:    true,
										Description: `Defines the schema/API version, if using Anthropic provider.`,
									},
									"azure_api_version": schema.StringAttribute{
										Computed:    true,
										Description: `'api-version' for Azure OpenAI instances.`,
									},
									"azure_deployment_id": schema.StringAttribute{
										Computed:    true,
										Description: `Deployment ID for Azure OpenAI instances.`,
									},
									"azure_instance": schema.StringAttribute{
										Computed:    true,
										Description: `Instance name for Azure OpenAI hosted models.`,
									},
									"llama2_format": schema.StringAttribute{
										Computed:    true,
										Description: `If using llama2 provider, select the upstream message format. must be one of ["raw", "openai", "ollama"]`,
									},
									"max_tokens": schema.Int64Attribute{
										Computed:    true,
										Description: `Defines the max_tokens, if using chat or completion models.`,
									},
									"mistral_format": schema.StringAttribute{
										Computed:    true,
										Description: `If using mistral provider, select the upstream message format. must be one of ["openai", "ollama"]`,
									},
									"temperature": schema.NumberAttribute{
										Computed:    true,
										Description: `Defines the matching temperature, if using chat or completion models.`,
									},
									"top_k": schema.Int64Attribute{
										Computed:    true,
										Description: `Defines the top-k most likely tokens, if supported.`,
									},
									"top_p": schema.NumberAttribute{
										Computed:    true,
										Description: `Defines the top-p probability mass, if supported.`,
									},
									"upstream_url": schema.StringAttribute{
										Computed:    true,
										Description: `Manually specify or override the full URL to the AI operation endpoints, when calling (self-)hosted models, or for running via a private endpoint.`,
									},
								},
								Description: `Key/value settings for the model`,
							},
							"provider": schema.StringAttribute{
								Computed:    true,
								Description: `AI provider request format - Kong translates requests to and from the specified backend compatible formats. must be one of ["openai", "azure", "anthropic", "cohere", "mistral", "llama2"]`,
							},
						},
					},
					"route_type": schema.StringAttribute{
						Computed:    true,
						Description: `The model's operation implementation, for this provider. must be one of ["llm/v1/chat", "llm/v1/completions"]`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the Plugin to lookup`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
		},
	}
}

func (r *PluginAIProxyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginAIProxyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginAIProxyDataSourceModel
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.GetAiproxyPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetAiproxyPlugin(ctx, request)
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
	if res.AIProxyPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAIProxyPlugin(res.AIProxyPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
