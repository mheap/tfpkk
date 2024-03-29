// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest struct {
	// Description of the Plugin
	CreatePluginWithoutParents shared.CreatePluginWithoutParents `request:"mediaType=application/json"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetCreatePluginWithoutParents() shared.CreatePluginWithoutParents {
	if o == nil {
		return shared.CreatePluginWithoutParents{}
	}
	return o.CreatePluginWithoutParents
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponseBody - Invalid Plugin
type PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponseBody struct {
}

type PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully upserted Plugin
	Plugin *shared.Plugin
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid Plugin
	Object *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponseBody
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetPlugin() *shared.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponse) GetObject() *PutControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDPluginsPluginIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
