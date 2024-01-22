// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
)

var _ provider.Provider = &KonnectProvider{}

type KonnectProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// KonnectProviderModel describes the provider data model.
type KonnectProviderModel struct {
	ServerURL                types.String `tfsdk:"server_url"`
	KonnectAccessToken       types.String `tfsdk:"konnect_access_token"`
	PersonalAccessToken      types.String `tfsdk:"personal_access_token"`
	SystemAccountAccessToken types.String `tfsdk:"system_account_access_token"`
}

func (p *KonnectProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "konnect"
	resp.Version = p.version
}

func (p *KonnectProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Konnect API: The Konnect platform API`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://global.api.konghq.com/v2)",
				Optional:            true,
				Required:            false,
			},
			"konnect_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"personal_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"system_account_access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *KonnectProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data KonnectProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://global.api.konghq.com/v2"
	}

	konnectAccessToken := new(string)
	if !data.KonnectAccessToken.IsUnknown() && !data.KonnectAccessToken.IsNull() {
		*konnectAccessToken = data.KonnectAccessToken.ValueString()
	} else {
		konnectAccessToken = nil
	}
	personalAccessToken := new(string)
	if !data.PersonalAccessToken.IsUnknown() && !data.PersonalAccessToken.IsNull() {
		*personalAccessToken = data.PersonalAccessToken.ValueString()
	} else {
		personalAccessToken = nil
	}
	systemAccountAccessToken := new(string)
	if !data.SystemAccountAccessToken.IsUnknown() && !data.SystemAccountAccessToken.IsNull() {
		*systemAccountAccessToken = data.SystemAccountAccessToken.ValueString()
	} else {
		systemAccountAccessToken = nil
	}
	security := shared.Security{
		KonnectAccessToken:       konnectAccessToken,
		PersonalAccessToken:      personalAccessToken,
		SystemAccountAccessToken: systemAccountAccessToken,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *KonnectProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIProductResource,
		NewAPIProductVersionResource,
		NewConsumerResource,
		NewGatewayControlPlaneResource,
		NewPluginResource,
		NewPortalResource,
		NewPortalAuthResource,
		NewRouteResource,
		NewServiceResource,
	}
}

func (p *KonnectProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAPIProductDataSource,
		NewAPIProductVersionDataSource,
		NewConsumerDataSource,
		NewGatewayControlPlaneDataSource,
		NewPluginDataSource,
		NewPortalAuthDataSource,
		NewRouteDataSource,
		NewServiceDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KonnectProvider{
			version: version,
		}
	}
}
