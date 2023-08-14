// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetOrderRequest struct {
	// The ID of the order get.
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderID"`
}

type GetOrderResponse struct {
	ContentType string
	// OK
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}