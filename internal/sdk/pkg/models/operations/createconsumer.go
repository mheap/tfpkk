// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateConsumerRequest struct {
	// Consumer request body
	ConsumerRequest *shared.ConsumerRequest `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateConsumerRequest) GetConsumerRequest() *shared.ConsumerRequest {
	if o == nil {
		return nil
	}
	return o.ConsumerRequest
}

func (o *CreateConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateConsumerResponseBody - Invalid Consumer
type CreateConsumerResponseBody struct {
}

type CreateConsumerResponse struct {
	// Successfully created consumer
	Consumer *shared.Consumer
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid Consumer
	Object *CreateConsumerResponseBody
}

func (o *CreateConsumerResponse) GetConsumer() *shared.Consumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateConsumerResponse) GetObject() *CreateConsumerResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
