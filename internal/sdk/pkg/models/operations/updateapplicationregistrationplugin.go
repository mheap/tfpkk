// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateApplicationregistrationPluginRequest struct {
	// Description of the Plugin
	CreateApplicationRegistrationPlugin shared.CreateApplicationRegistrationPlugin `request:"mediaType=application/json"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *UpdateApplicationregistrationPluginRequest) GetCreateApplicationRegistrationPlugin() shared.CreateApplicationRegistrationPlugin {
	if o == nil {
		return shared.CreateApplicationRegistrationPlugin{}
	}
	return o.CreateApplicationRegistrationPlugin
}

func (o *UpdateApplicationregistrationPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateApplicationregistrationPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// UpdateApplicationregistrationPluginResponseBody - Invalid Plugin
type UpdateApplicationregistrationPluginResponseBody struct {
}

type UpdateApplicationregistrationPluginResponse struct {
	// Successfully upserted Plugin
	ApplicationRegistrationPlugin *shared.ApplicationRegistrationPlugin
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Plugin
	Object *UpdateApplicationregistrationPluginResponseBody
}

func (o *UpdateApplicationregistrationPluginResponse) GetApplicationRegistrationPlugin() *shared.ApplicationRegistrationPlugin {
	if o == nil {
		return nil
	}
	return o.ApplicationRegistrationPlugin
}

func (o *UpdateApplicationregistrationPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateApplicationregistrationPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateApplicationregistrationPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateApplicationregistrationPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateApplicationregistrationPluginResponse) GetObject() *UpdateApplicationregistrationPluginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}