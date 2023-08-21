// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetPortalAuthenticationSettingsRequest struct {
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

type GetPortalAuthenticationSettings200ApplicationJSONOidcConfig struct {
	ClientID string `json:"client_id"`
	Issuer   string `json:"issuer"`
}

// GetPortalAuthenticationSettings200ApplicationJSON - Response for authentication settings endpoint
type GetPortalAuthenticationSettings200ApplicationJSON struct {
	// The organization has basic auth enabled.
	BasicAuthEnabled bool `json:"basic_auth_enabled"`
	// A Konnect Identity Admin assigns teams to a developer.
	KonnectMappingEnabled bool `json:"konnect_mapping_enabled"`
	// The organization has OIDC disabled.
	OidcAuthEnabled bool                                                         `json:"oidc_auth_enabled"`
	OidcConfig      *GetPortalAuthenticationSettings200ApplicationJSONOidcConfig `json:"oidc_config,omitempty"`
	// IdP groups determine the Portal Teams a developer has.
	OidcTeamMappingEnabled bool `json:"oidc_team_mapping_enabled"`
}

type GetPortalAuthenticationSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Response for authentication settings endpoint
	GetPortalAuthenticationSettings200ApplicationJSONObject *GetPortalAuthenticationSettings200ApplicationJSON
}
