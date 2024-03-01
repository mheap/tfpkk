// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDRequest struct {
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// Description of the JWT
	CreateJWTWithoutParents shared.CreateJWTWithoutParents `request:"mediaType=application/json"`
	// ID of the JWT to lookup
	JWTID string `pathParam:"style=simple,explode=false,name=JWTId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDRequest) GetCreateJWTWithoutParents() shared.CreateJWTWithoutParents {
	if o == nil {
		return shared.CreateJWTWithoutParents{}
	}
	return o.CreateJWTWithoutParents
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDRequest) GetJWTID() string {
	if o == nil {
		return ""
	}
	return o.JWTID
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponseBody - Invalid JWT
type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponseBody struct {
}

type PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully upserted JWT
	Jwt *shared.Jwt
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid JWT
	Object *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponseBody
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse) GetJwt() *shared.Jwt {
	if o == nil {
		return nil
	}
	return o.Jwt
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponse) GetObject() *PutControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDJWTJWTIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}