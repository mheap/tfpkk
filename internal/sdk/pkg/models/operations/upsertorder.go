// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpsertOrderResponse struct {
	ContentType string
	// OK
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}
