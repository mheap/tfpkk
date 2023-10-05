// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAPIProductRequest struct {
	// API product identifier
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPIProductResponse struct {
	// API product
	APIProduct *shared.APIProduct
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
}
