// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateAPIProductResponse struct {
	// API product
	APIProduct *shared.APIProduct
	// Bad Request
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Unsupported Media Type
	UnsupportedMediaTypeError *shared.UnsupportedMediaTypeError
}
