// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully fetched Plugin
	Plugin *shared.Plugin
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetPlugin() *shared.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
