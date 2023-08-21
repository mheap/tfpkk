// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListRoutesForServiceRequest struct {
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
	// ID or name of the related service
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Number of resources to be returned.
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

// ListRoutesForService200ApplicationJSON - A successful response listing routes
type ListRoutesForService200ApplicationJSON struct {
	Data []shared.Route `json:"data,omitempty"`
	// Offset is used to paginate through the API. Provide this value to the next list operation to fetch the next page
	Offset *string `json:"offset,omitempty"`
}

type ListRoutesForServiceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A successful response listing routes
	ListRoutesForService200ApplicationJSONObject *ListRoutesForService200ApplicationJSON
}
