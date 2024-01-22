// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreateControlPlaneStatus - The HTTP status code.
type CreateControlPlaneStatus int64

const (
	CreateControlPlaneStatusFiveHundredAndThree CreateControlPlaneStatus = 503
)

func (e CreateControlPlaneStatus) ToPointer() *CreateControlPlaneStatus {
	return &e
}

func (e *CreateControlPlaneStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = CreateControlPlaneStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateControlPlaneStatus: %v", v)
	}
}

// CreateControlPlaneServiceUnavailableError - Error response for temporary service unavailability.
type CreateControlPlaneServiceUnavailableError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status CreateControlPlaneStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

func (o *CreateControlPlaneServiceUnavailableError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *CreateControlPlaneServiceUnavailableError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *CreateControlPlaneServiceUnavailableError) GetStatus() CreateControlPlaneStatus {
	if o == nil {
		return CreateControlPlaneStatus(0)
	}
	return o.Status
}

func (o *CreateControlPlaneServiceUnavailableError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

// Status - The HTTP status code.
type Status int64

const (
	StatusFiveHundred Status = 500
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 500:
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// CreateControlPlaneInternalServerError - The error response object.
type CreateControlPlaneInternalServerError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status Status `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

func (o *CreateControlPlaneInternalServerError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *CreateControlPlaneInternalServerError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *CreateControlPlaneInternalServerError) GetStatus() Status {
	if o == nil {
		return Status(0)
	}
	return o.Status
}

func (o *CreateControlPlaneInternalServerError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

type CreateControlPlaneResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Conflict
	ConflictError *shared.ConflictError
	// HTTP response content type for this operation
	ContentType string
	// A response to creating a control plane.
	ControlPlane *shared.ControlPlane
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Internal Server Error
	InternalServerError *CreateControlPlaneInternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Service Unavailable
	ServiceUnavailableError *CreateControlPlaneServiceUnavailableError
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *CreateControlPlaneResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}

func (o *CreateControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateControlPlaneResponse) GetControlPlane() *shared.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}

func (o *CreateControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *CreateControlPlaneResponse) GetInternalServerError() *CreateControlPlaneInternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateControlPlaneResponse) GetServiceUnavailableError() *CreateControlPlaneServiceUnavailableError {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailableError
}

func (o *CreateControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
