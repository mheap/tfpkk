// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateConsumerGroupRequest struct {
	// Description of the new Consumer Group for creation
	CreateConsumerGroup shared.CreateConsumerGroup `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateConsumerGroupRequest) GetCreateConsumerGroup() shared.CreateConsumerGroup {
	if o == nil {
		return shared.CreateConsumerGroup{}
	}
	return o.CreateConsumerGroup
}

func (o *CreateConsumerGroupRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateConsumerGroupResponseBody - Invalid Consumer Group
type CreateConsumerGroupResponseBody struct {
}

type CreateConsumerGroupResponse struct {
	// Successfully created Consumer Group
	ConsumerGroup *shared.ConsumerGroup
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Consumer Group
	Object *CreateConsumerGroupResponseBody
}

func (o *CreateConsumerGroupResponse) GetConsumerGroup() *shared.ConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateConsumerGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConsumerGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConsumerGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateConsumerGroupResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateConsumerGroupResponse) GetObject() *CreateConsumerGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
