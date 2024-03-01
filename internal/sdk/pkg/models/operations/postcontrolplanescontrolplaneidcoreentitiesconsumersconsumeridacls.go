// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest struct {
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// Description of new ACL for creation
	CreateACLWithoutParents shared.CreateACLWithoutParents `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetCreateACLWithoutParents() shared.CreateACLWithoutParents {
	if o == nil {
		return shared.CreateACLWithoutParents{}
	}
	return o.CreateACLWithoutParents
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody - Invalid ACL
type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody struct {
}

type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse struct {
	// Successfully created ACL
	ACL *shared.ACL
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid ACL
	Object *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetACL() *shared.ACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetObject() *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}