// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateBasicauthPluginRequest struct {
	// Description of the Plugin
	CreateBasicAuthPlugin shared.CreateBasicAuthPlugin `request:"mediaType=application/json"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *UpdateBasicauthPluginRequest) GetCreateBasicAuthPlugin() shared.CreateBasicAuthPlugin {
	if o == nil {
		return shared.CreateBasicAuthPlugin{}
	}
	return o.CreateBasicAuthPlugin
}

func (o *UpdateBasicauthPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateBasicauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// UpdateBasicauthPluginResponseBody - Invalid Plugin
type UpdateBasicauthPluginResponseBody struct {
}

type UpdateBasicauthPluginResponse struct {
	// Successfully upserted Plugin
	BasicAuthPlugin *shared.BasicAuthPlugin
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Plugin
	Object *UpdateBasicauthPluginResponseBody
}

func (o *UpdateBasicauthPluginResponse) GetBasicAuthPlugin() *shared.BasicAuthPlugin {
	if o == nil {
		return nil
	}
	return o.BasicAuthPlugin
}

func (o *UpdateBasicauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBasicauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBasicauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBasicauthPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateBasicauthPluginResponse) GetObject() *UpdateBasicauthPluginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
