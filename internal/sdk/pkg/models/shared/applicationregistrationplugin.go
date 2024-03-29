// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
)

type ApplicationRegistrationPluginConfig struct {
	// If enabled, all new Service Contracts requests are automatically approved.
	AutoApprove *bool `default:"false" json:"auto_approve"`
	// Unique description displayed in information about a Service in the Developer Portal.
	Description *string `json:"description,omitempty"`
	// Unique display name used for a Service in the Developer Portal.
	DisplayName *string `json:"display_name,omitempty"`
	// Displays the **Issuer URL** in the **Service Details** dialog.
	ShowIssuer *bool `default:"false" json:"show_issuer"`
}

func (a ApplicationRegistrationPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationRegistrationPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationRegistrationPluginConfig) GetAutoApprove() *bool {
	if o == nil {
		return nil
	}
	return o.AutoApprove
}

func (o *ApplicationRegistrationPluginConfig) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ApplicationRegistrationPluginConfig) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ApplicationRegistrationPluginConfig) GetShowIssuer() *bool {
	if o == nil {
		return nil
	}
	return o.ShowIssuer
}

// ApplicationRegistrationPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type ApplicationRegistrationPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ApplicationRegistrationPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ApplicationRegistrationPluginProtocols string

const (
	ApplicationRegistrationPluginProtocolsGrpc           ApplicationRegistrationPluginProtocols = "grpc"
	ApplicationRegistrationPluginProtocolsGrpcs          ApplicationRegistrationPluginProtocols = "grpcs"
	ApplicationRegistrationPluginProtocolsHTTP           ApplicationRegistrationPluginProtocols = "http"
	ApplicationRegistrationPluginProtocolsHTTPS          ApplicationRegistrationPluginProtocols = "https"
	ApplicationRegistrationPluginProtocolsTCP            ApplicationRegistrationPluginProtocols = "tcp"
	ApplicationRegistrationPluginProtocolsTLS            ApplicationRegistrationPluginProtocols = "tls"
	ApplicationRegistrationPluginProtocolsTLSPassthrough ApplicationRegistrationPluginProtocols = "tls_passthrough"
	ApplicationRegistrationPluginProtocolsUDP            ApplicationRegistrationPluginProtocols = "udp"
	ApplicationRegistrationPluginProtocolsWs             ApplicationRegistrationPluginProtocols = "ws"
	ApplicationRegistrationPluginProtocolsWss            ApplicationRegistrationPluginProtocols = "wss"
)

func (e ApplicationRegistrationPluginProtocols) ToPointer() *ApplicationRegistrationPluginProtocols {
	return &e
}

func (e *ApplicationRegistrationPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = ApplicationRegistrationPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplicationRegistrationPluginProtocols: %v", v)
	}
}

// ApplicationRegistrationPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type ApplicationRegistrationPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *ApplicationRegistrationPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ApplicationRegistrationPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type ApplicationRegistrationPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *ApplicationRegistrationPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// ApplicationRegistrationPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type ApplicationRegistrationPlugin struct {
	Config ApplicationRegistrationPluginConfig `json:"config"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *ApplicationRegistrationPluginConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled *bool   `default:"true" json:"enabled"`
	ID      *string `json:"id,omitempty"`
	name    string  `const:"application-registration" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []ApplicationRegistrationPluginProtocols `json:"protocols"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *ApplicationRegistrationPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *ApplicationRegistrationPluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (a ApplicationRegistrationPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationRegistrationPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationRegistrationPlugin) GetConfig() ApplicationRegistrationPluginConfig {
	if o == nil {
		return ApplicationRegistrationPluginConfig{}
	}
	return o.Config
}

func (o *ApplicationRegistrationPlugin) GetConsumer() *ApplicationRegistrationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *ApplicationRegistrationPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ApplicationRegistrationPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ApplicationRegistrationPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ApplicationRegistrationPlugin) GetName() string {
	return "application-registration"
}

func (o *ApplicationRegistrationPlugin) GetProtocols() []ApplicationRegistrationPluginProtocols {
	if o == nil {
		return []ApplicationRegistrationPluginProtocols{}
	}
	return o.Protocols
}

func (o *ApplicationRegistrationPlugin) GetRoute() *ApplicationRegistrationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *ApplicationRegistrationPlugin) GetService() *ApplicationRegistrationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ApplicationRegistrationPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
