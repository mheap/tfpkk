// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDRequest struct {
	// ID of the Basic-auth credential to lookup
	BasicAuthID string `pathParam:"style=simple,explode=false,name=BasicAuthId"`
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDRequest) GetBasicAuthID() string {
	if o == nil {
		return ""
	}
	return o.BasicAuthID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDResponse struct {
	// Successfully fetched Basic-auth credential
	BasicAuth *shared.BasicAuth
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDResponse) GetBasicAuth() *shared.BasicAuth {
	if o == nil {
		return nil
	}
	return o.BasicAuth
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDBasicAuthBasicAuthIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
