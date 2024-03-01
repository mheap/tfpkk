// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateRouteRequest struct {
	// Description of the new Route for creation
	CreateRoute shared.CreateRoute `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateRouteRequest) GetCreateRoute() shared.CreateRoute {
	if o == nil {
		return shared.CreateRoute{}
	}
	return o.CreateRoute
}

func (o *CreateRouteRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateRouteResponseBody - Invalid Route
type CreateRouteResponseBody struct {
}

type CreateRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created Route
	Route *shared.Route
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Route
	Object *CreateRouteResponseBody
}

func (o *CreateRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRouteResponse) GetRoute() *shared.Route {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRouteResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateRouteResponse) GetObject() *CreateRouteResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
