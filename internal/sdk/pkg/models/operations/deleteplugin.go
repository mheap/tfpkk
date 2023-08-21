// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeletePluginRequest struct {
	// The unique identifier of the Plugin to create or update.
	PluginIDOrInstanceName string `pathParam:"style=simple,explode=false,name=plugin_id_or_instance_name"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type DeletePluginResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
