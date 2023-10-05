// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetPortalAuthenticationSettingsRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

type GetPortalAuthenticationSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Response for authentication settings endpoint
	PortalAuthenticationSettings *shared.PortalAuthenticationSettings
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}