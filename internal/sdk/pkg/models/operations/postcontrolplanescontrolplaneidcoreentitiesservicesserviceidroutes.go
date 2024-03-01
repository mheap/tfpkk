// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRequest struct {
	// Description of new Route for creation
	CreateRouteWithoutParents shared.CreateRouteWithoutParents `request:"mediaType=application/json"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRequest) GetCreateRouteWithoutParents() shared.CreateRouteWithoutParents {
	if o == nil {
		return shared.CreateRouteWithoutParents{}
	}
	return o.CreateRouteWithoutParents
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponseBody - Invalid Route
type PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponseBody struct {
}

type PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created Route
	Route *shared.Route
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid Route
	Object *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponseBody
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse) GetRoute() *shared.Route {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponse) GetObject() *PostControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
