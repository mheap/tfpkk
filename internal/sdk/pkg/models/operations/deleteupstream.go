// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUpstreamRequest struct {
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
	// The unique identifier of the Upstream associated to the Certificate to be retrieved.
	UpstreamID string `pathParam:"style=simple,explode=false,name=upstream_id"`
}

type DeleteUpstreamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
