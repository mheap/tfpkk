// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAPIProductDocumentRequest struct {
	// The API product identifier
	APIProductID string `pathParam:"style=simple,explode=false,name=apiProductId"`
	// The document identifier related to the API product
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPIProductDocumentResponse struct {
	// API product document
	APIProductDocument *shared.APIProductDocument
	ContentType        string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	StatusCode    int
	RawResponse   *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}
