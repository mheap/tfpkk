// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdatePortalRequest struct {
	// Update a portal
	UpdatePortalRequest shared.UpdatePortalRequest `request:"mediaType=application/json"`
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
}

type UpdatePortalResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError    *shared.UnauthorizedError
	UpdatePortalResponse *shared.UpdatePortalResponse
}
