// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteKeyRequest struct {
	// The unique identifier or the name of the Key to retrieve.
	KeyIDOrName string `pathParam:"style=simple,explode=false,name=key_id_or_name"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type DeleteKeyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
