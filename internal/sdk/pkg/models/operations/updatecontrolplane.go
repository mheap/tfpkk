// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateControlPlaneRequest struct {
	UpdateControlPlaneRequest shared.UpdateControlPlaneRequest `request:"mediaType=application/json"`
	// The control plane ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateControlPlaneRequest) GetUpdateControlPlaneRequest() shared.UpdateControlPlaneRequest {
	if o == nil {
		return shared.UpdateControlPlaneRequest{}
	}
	return o.UpdateControlPlaneRequest
}

func (o *UpdateControlPlaneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// UpdateControlPlaneControlPlanesStatus - The HTTP status code.
type UpdateControlPlaneControlPlanesStatus int64

const (
	UpdateControlPlaneControlPlanesStatusFiveHundredAndThree UpdateControlPlaneControlPlanesStatus = 503
)

func (e UpdateControlPlaneControlPlanesStatus) ToPointer() *UpdateControlPlaneControlPlanesStatus {
	return &e
}

func (e *UpdateControlPlaneControlPlanesStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = UpdateControlPlaneControlPlanesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateControlPlaneControlPlanesStatus: %v", v)
	}
}

// UpdateControlPlaneServiceUnavailableError - Error response for temporary service unavailability.
type UpdateControlPlaneServiceUnavailableError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status UpdateControlPlaneControlPlanesStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

func (o *UpdateControlPlaneServiceUnavailableError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *UpdateControlPlaneServiceUnavailableError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *UpdateControlPlaneServiceUnavailableError) GetStatus() UpdateControlPlaneControlPlanesStatus {
	if o == nil {
		return UpdateControlPlaneControlPlanesStatus(0)
	}
	return o.Status
}

func (o *UpdateControlPlaneServiceUnavailableError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

// UpdateControlPlaneStatus - The HTTP status code.
type UpdateControlPlaneStatus int64

const (
	UpdateControlPlaneStatusFiveHundred UpdateControlPlaneStatus = 500
)

func (e UpdateControlPlaneStatus) ToPointer() *UpdateControlPlaneStatus {
	return &e
}

func (e *UpdateControlPlaneStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 500:
		*e = UpdateControlPlaneStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateControlPlaneStatus: %v", v)
	}
}

// UpdateControlPlaneInternalServerError - The error response object.
type UpdateControlPlaneInternalServerError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status UpdateControlPlaneStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

func (o *UpdateControlPlaneInternalServerError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *UpdateControlPlaneInternalServerError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *UpdateControlPlaneInternalServerError) GetStatus() UpdateControlPlaneStatus {
	if o == nil {
		return UpdateControlPlaneStatus(0)
	}
	return o.Status
}

func (o *UpdateControlPlaneInternalServerError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

type UpdateControlPlaneResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// A response to updating a control plane.
	ControlPlane *shared.ControlPlane
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Internal Server Error
	InternalServerError *UpdateControlPlaneInternalServerError
	// Not Found
	NotFoundError *shared.NotFoundError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Service Unavailable
	ServiceUnavailableError *UpdateControlPlaneServiceUnavailableError
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *UpdateControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateControlPlaneResponse) GetControlPlane() *shared.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}

func (o *UpdateControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdateControlPlaneResponse) GetInternalServerError() *UpdateControlPlaneInternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateControlPlaneResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *UpdateControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateControlPlaneResponse) GetServiceUnavailableError() *UpdateControlPlaneServiceUnavailableError {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailableError
}

func (o *UpdateControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
