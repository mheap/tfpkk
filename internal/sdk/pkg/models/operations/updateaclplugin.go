// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateACLPluginRequest struct {
	// Description of the Plugin
	CreateACLPlugin shared.CreateACLPlugin `request:"mediaType=application/json"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *UpdateACLPluginRequest) GetCreateACLPlugin() shared.CreateACLPlugin {
	if o == nil {
		return shared.CreateACLPlugin{}
	}
	return o.CreateACLPlugin
}

func (o *UpdateACLPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateACLPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// UpdateACLPluginResponseBody - Invalid Plugin
type UpdateACLPluginResponseBody struct {
}

type UpdateACLPluginResponse struct {
	// Successfully upserted Plugin
	ACLPlugin *shared.ACLPlugin
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Plugin
	Object *UpdateACLPluginResponseBody
}

func (o *UpdateACLPluginResponse) GetACLPlugin() *shared.ACLPlugin {
	if o == nil {
		return nil
	}
	return o.ACLPlugin
}

func (o *UpdateACLPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateACLPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateACLPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateACLPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdateACLPluginResponse) GetObject() *UpdateACLPluginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}