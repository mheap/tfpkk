// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateAPIProductVersionRequest struct {
	UpdateAPIProductVersionDTO shared.UpdateAPIProductVersionDTO `request:"mediaType=application/json"`
	// The API product identifier
	APIProductID string `pathParam:"style=simple,explode=false,name=apiProductId"`
	// The API product version identifier
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateAPIProductVersionResponse struct {
	// API product
	APIProductVersion *shared.APIProductVersion
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
	UnauthorizedError *shared.UnauthorizedError
	// Unsupported Media Type
	UnsupportedMediaTypeError *shared.UnsupportedMediaTypeError
}
