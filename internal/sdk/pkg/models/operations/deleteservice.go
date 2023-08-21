// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteServiceRequest struct {
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
	// ID or name of the service to delete
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

type DeleteServiceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
