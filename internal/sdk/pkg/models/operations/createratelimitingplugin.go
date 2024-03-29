// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateRatelimitingPluginRequest struct {
	// Create a new RateLimiting plugin
	CreateRateLimitingPlugin shared.CreateRateLimitingPlugin `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateRatelimitingPluginRequest) GetCreateRateLimitingPlugin() shared.CreateRateLimitingPlugin {
	if o == nil {
		return shared.CreateRateLimitingPlugin{}
	}
	return o.CreateRateLimitingPlugin
}

func (o *CreateRatelimitingPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateRatelimitingPluginResponseBody - Invalid Plugin
type CreateRatelimitingPluginResponseBody struct {
}

type CreateRatelimitingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created Plugin
	RateLimitingPlugin *shared.RateLimitingPlugin
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Plugin
	Object *CreateRatelimitingPluginResponseBody
}

func (o *CreateRatelimitingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRatelimitingPluginResponse) GetRateLimitingPlugin() *shared.RateLimitingPlugin {
	if o == nil {
		return nil
	}
	return o.RateLimitingPlugin
}

func (o *CreateRatelimitingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRatelimitingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRatelimitingPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateRatelimitingPluginResponse) GetObject() *CreateRatelimitingPluginResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
