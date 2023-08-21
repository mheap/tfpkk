// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteRuntimeGroupIDCoreEntitiesConsumerGroupsGroupIDConsumersRequest struct {
	// The UUID of the consumer group.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type DeleteRuntimeGroupIDCoreEntitiesConsumerGroupsGroupIDConsumersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
