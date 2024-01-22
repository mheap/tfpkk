// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
)

// PluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type PluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Ordering struct {
}

// PluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type PluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// PluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified service. Leave unset for the plugin to activate regardless of the service being matched.
type PluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *PluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle.
type Plugin struct {
	// The configuration properties for the Plugin which can be found on the plugins documentation page in the [Kong Hub](https://docs.konghq.com/hub/).
	Config interface{} `json:"config,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *PluginConsumer `json:"consumer,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `default:"true" json:"enabled"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	// The name of the Plugin that's going to be added. Currently, the Plugin must be installed in every Kong instance separately.
	Name     *string   `json:"name,omitempty"`
	Ordering *Ordering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `tcp` and `tls`.
	Protocols []string `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *PluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified service. Leave unset for the plugin to activate regardless of the service being matched.
	Service *PluginService `json:"service,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (p Plugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Plugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Plugin) GetConfig() interface{} {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Plugin) GetConsumer() *PluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *Plugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Plugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Plugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Plugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *Plugin) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Plugin) GetOrdering() *Ordering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *Plugin) GetProtocols() []string {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *Plugin) GetRoute() *PluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *Plugin) GetService() *PluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *Plugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
