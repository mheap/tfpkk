// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// UpsertConsumerRequestBody - Consumer request body
type UpsertConsumerRequestBody struct {
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

type UpsertConsumerRequest struct {
	// Consumer request body
	RequestBody *UpsertConsumerRequestBody `request:"mediaType=application/json"`
	// The unique identifier or the username of the Consumer to retrieve.
	ConsumerID string `pathParam:"style=simple,explode=false,name=consumer_id"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type UpsertConsumer200ApplicationJSONData struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database.
	CustomID *string `json:"custom_id,omitempty"`
	// The unique identifier or the name attribute of the consumer.
	ID *string `json:"id,omitempty"`
	// An optional set of strings associated with the Consumer for grouping and filtering.
	//
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was created.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The unique username of the consumer. You must send either this field or` custom_i`d with the request.
	Username *string `json:"username,omitempty"`
}

// UpsertConsumer200ApplicationJSON - The consumer object response body
type UpsertConsumer200ApplicationJSON struct {
	Data []UpsertConsumer200ApplicationJSONData `json:"data,omitempty"`
	// Pagination information
	Next *string `json:"next,omitempty"`
}

type UpsertConsumerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The consumer object response body
	UpsertConsumer200ApplicationJSONObject *UpsertConsumer200ApplicationJSON
}
