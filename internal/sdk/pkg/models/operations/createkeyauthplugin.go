// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateKeyauthPluginRequest struct {
	// Create a new KeyAuth plugin
	CreateKeyAuthPlugin shared.CreateKeyAuthPlugin `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateKeyauthPluginRequest) GetCreateKeyAuthPlugin() shared.CreateKeyAuthPlugin {
	if o == nil {
		return shared.CreateKeyAuthPlugin{}
	}
	return o.CreateKeyAuthPlugin
}

func (o *CreateKeyauthPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateKeyauthPluginResponseBody - Invalid Plugin
type CreateKeyauthPluginResponseBody struct {
}

type CreateKeyauthPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created Plugin
	KeyAuthPlugin *shared.KeyAuthPlugin
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Plugin
	Object *CreateKeyauthPluginResponseBody
}

func (o *CreateKeyauthPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateKeyauthPluginResponse) GetKeyAuthPlugin() *shared.KeyAuthPlugin {
	if o == nil {
		return nil
	}
	return o.KeyAuthPlugin
}

func (o *CreateKeyauthPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateKeyauthPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateKeyauthPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateKeyauthPluginResponse) GetObject() *CreateKeyauthPluginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
