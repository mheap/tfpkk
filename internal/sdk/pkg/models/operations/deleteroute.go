// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteRouteRequest struct {
	RouteID string `pathParam:"style=simple,explode=false,name=route_id"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type DeleteRouteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
