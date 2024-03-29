// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSniRequest struct {
	// Description of the new SNI for creation
	CreateSNI shared.CreateSNI `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateSniRequest) GetCreateSNI() shared.CreateSNI {
	if o == nil {
		return shared.CreateSNI{}
	}
	return o.CreateSNI
}

func (o *CreateSniRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateSniResponseBody - Invalid SNI
type CreateSniResponseBody struct {
}

type CreateSniResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created SNI
	Sni *shared.Sni
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid SNI
	Object *CreateSniResponseBody
}

func (o *CreateSniResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSniResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *CreateSniResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSniResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSniResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateSniResponse) GetObject() *CreateSniResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
