// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetControlPlaneRequest struct {
	// The control plane ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetControlPlaneStatus - The HTTP status code.
type GetControlPlaneStatus int64

const (
	GetControlPlaneStatusFiveHundredAndThree GetControlPlaneStatus = 503
)

func (e GetControlPlaneStatus) ToPointer() *GetControlPlaneStatus {
	return &e
}

func (e *GetControlPlaneStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = GetControlPlaneStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetControlPlaneStatus: %v", v)
	}
}

// GetControlPlaneServiceUnavailableError - Error response for temporary service unavailability.
type GetControlPlaneServiceUnavailableError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status GetControlPlaneStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

func (o *GetControlPlaneServiceUnavailableError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *GetControlPlaneServiceUnavailableError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *GetControlPlaneServiceUnavailableError) GetStatus() GetControlPlaneStatus {
	if o == nil {
		return GetControlPlaneStatus(0)
	}
	return o.Status
}

func (o *GetControlPlaneServiceUnavailableError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

type GetControlPlaneResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// A response to retrieving a single control plane.
	ControlPlane *shared.ControlPlane
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Service Unavailable
	ServiceUnavailableError *GetControlPlaneServiceUnavailableError
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
}

func (o *GetControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *GetControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetControlPlaneResponse) GetControlPlane() *shared.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}

func (o *GetControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *GetControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetControlPlaneResponse) GetServiceUnavailableError() *GetControlPlaneServiceUnavailableError {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailableError
}

func (o *GetControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
