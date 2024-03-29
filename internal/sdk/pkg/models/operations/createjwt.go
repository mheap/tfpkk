// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateJwtRequest struct {
	// Description of the new JWT for creation
	CreateJWT shared.CreateJWT `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateJwtRequest) GetCreateJWT() shared.CreateJWT {
	if o == nil {
		return shared.CreateJWT{}
	}
	return o.CreateJWT
}

func (o *CreateJwtRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateJwtResponseBody - Invalid JWT
type CreateJwtResponseBody struct {
}

type CreateJwtResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created JWT
	Jwt *shared.Jwt
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid JWT
	Object *CreateJwtResponseBody
}

func (o *CreateJwtResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateJwtResponse) GetJwt() *shared.Jwt {
	if o == nil {
		return nil
	}
	return o.Jwt
}

func (o *CreateJwtResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateJwtResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateJwtResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateJwtResponse) GetObject() *CreateJwtResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
