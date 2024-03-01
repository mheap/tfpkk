// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateCertificateRequest struct {
	// Description of the new Certificate for creation
	CreateCertificate shared.CreateCertificate `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateCertificateRequest) GetCreateCertificate() shared.CreateCertificate {
	if o == nil {
		return shared.CreateCertificate{}
	}
	return o.CreateCertificate
}

func (o *CreateCertificateRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateCertificateResponseBody - Invalid Certificate
type CreateCertificateResponseBody struct {
}

type CreateCertificateResponse struct {
	// Successfully created Certificate
	Certificate *shared.Certificate
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid Certificate
	Object *CreateCertificateResponseBody
}

func (o *CreateCertificateResponse) GetCertificate() *shared.Certificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *CreateCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCertificateResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateCertificateResponse) GetObject() *CreateCertificateResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}