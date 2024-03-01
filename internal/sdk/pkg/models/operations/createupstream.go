// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateUpstreamRequest struct {
	// Description of the new Upstream for creation
	CreateUpstream shared.CreateUpstream `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateUpstreamRequest) GetCreateUpstream() shared.CreateUpstream {
	if o == nil {
		return shared.CreateUpstream{}
	}
	return o.CreateUpstream
}

func (o *CreateUpstreamRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateUpstreamResponseBody - Invalid Upstream
type CreateUpstreamResponseBody struct {
}

type CreateUpstreamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Successfully created Upstream
	Upstream *shared.Upstream
	// Invalid Upstream
	Object *CreateUpstreamResponseBody
}

func (o *CreateUpstreamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUpstreamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUpstreamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUpstreamResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateUpstreamResponse) GetUpstream() *shared.Upstream {
	if o == nil {
		return nil
	}
	return o.Upstream
}

func (o *CreateUpstreamResponse) GetObject() *CreateUpstreamResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}