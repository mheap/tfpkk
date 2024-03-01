// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDRequest struct {
	// ID of the ACL to lookup
	ACLID string `pathParam:"style=simple,explode=false,name=ACLId"`
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// Description of the ACL
	CreateACLWithoutParents shared.CreateACLWithoutParents `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDRequest) GetCreateACLWithoutParents() shared.CreateACLWithoutParents {
	if o == nil {
		return shared.CreateACLWithoutParents{}
	}
	return o.CreateACLWithoutParents
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponseBody - Invalid ACL
type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponseBody struct {
}

type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse struct {
	// Successfully upserted ACL
	ACL *shared.ACL
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid ACL
	Object *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponseBody
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse) GetACL() *shared.ACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponse) GetObject() *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsACLIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
