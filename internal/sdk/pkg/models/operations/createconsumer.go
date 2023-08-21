// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// CreateConsumerRequestBody - Consumer request body
type CreateConsumerRequestBody struct {
	// Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or username with the request.
	//
	CustomID string `json:"custom_id"`
	// An optional set of strings associated with the Consumer for grouping and filtering.
	//
	Tags []string `json:"tags,omitempty"`
	// The unique username of the Consumer. You must send either this field or custom_id with the request.
	//
	Username string `json:"username"`
}

type CreateConsumerRequest struct {
	// Consumer request body
	RequestBody *CreateConsumerRequestBody `request:"mediaType=application/json"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

// CreateConsumer400ApplicationJSON - Invalid Consumer
type CreateConsumer400ApplicationJSON struct {
}

// CreateConsumer200ApplicationJSON - New consumer created response
type CreateConsumer200ApplicationJSON struct {
	// Unix epoch when the resource was created.
	//
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Field for the unique consumer ID
	CustomID *string `json:"custom_id,omitempty"`
	// The unique id of the consumer.
	ID *string `json:"id,omitempty"`
	// An optional set of strings associated with the Consumer for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was updated.
	//
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The unique username of the consumer.
	Username *string `json:"username,omitempty"`
}

type CreateConsumerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// New consumer created response
	CreateConsumer200ApplicationJSONObject *CreateConsumer200ApplicationJSON
	// Invalid Consumer
	CreateConsumer400ApplicationJSONObject *CreateConsumer400ApplicationJSON
}
