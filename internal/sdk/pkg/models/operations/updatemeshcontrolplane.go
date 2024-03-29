// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

var UpdateMeshControlPlaneServerList = []string{
	"https://us.api.konghq.com/v1",
	"https://us.api.konghq.com/v1",
}

type UpdateMeshControlPlaneRequest struct {
	UpdateMeshControlPlaneRequest shared.UpdateMeshControlPlaneRequest `request:"mediaType=application/json"`
	// Id of the Konnect resource
	CpID string `pathParam:"style=simple,explode=false,name=cpId"`
}

func (o *UpdateMeshControlPlaneRequest) GetUpdateMeshControlPlaneRequest() shared.UpdateMeshControlPlaneRequest {
	if o == nil {
		return shared.UpdateMeshControlPlaneRequest{}
	}
	return o.UpdateMeshControlPlaneRequest
}

func (o *UpdateMeshControlPlaneRequest) GetCpID() string {
	if o == nil {
		return ""
	}
	return o.CpID
}

type UpdateMeshControlPlaneResponse struct {
	// Validation Error
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// A response to updating a control plane.
	MeshControlPlane *shared.MeshControlPlane
	// Not found
	NotFoundError *shared.NotFoundError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized Error
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateMeshControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdateMeshControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeshControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdateMeshControlPlaneResponse) GetMeshControlPlane() *shared.MeshControlPlane {
	if o == nil {
		return nil
	}
	return o.MeshControlPlane
}

func (o *UpdateMeshControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *UpdateMeshControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeshControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeshControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
