// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateACLRequest struct {
	// Description of the new ACL for creation
	CreateACL shared.CreateACL `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *CreateACLRequest) GetCreateACL() shared.CreateACL {
	if o == nil {
		return shared.CreateACL{}
	}
	return o.CreateACL
}

func (o *CreateACLRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// CreateACLResponseBody - Invalid ACL
type CreateACLResponseBody struct {
}

type CreateACLResponse struct {
	// Successfully created ACL
	ACL *shared.ACL
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Invalid ACL
	Object *CreateACLResponseBody
}

func (o *CreateACLResponse) GetACL() *shared.ACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *CreateACLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateACLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateACLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateACLResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateACLResponse) GetObject() *CreateACLResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}