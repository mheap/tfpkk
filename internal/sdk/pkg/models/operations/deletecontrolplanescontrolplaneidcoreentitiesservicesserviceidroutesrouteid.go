// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDRequest struct {
	// ID of the Route to lookup
	RouteID string `pathParam:"style=simple,explode=false,name=RouteId"`
	// ID of the Service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=ServiceId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDRequest) GetRouteID() string {
	if o == nil {
		return ""
	}
	return o.RouteID
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

type DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteControlPlanesControlPlaneIDCoreEntitiesServicesServiceIDRoutesRouteIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}