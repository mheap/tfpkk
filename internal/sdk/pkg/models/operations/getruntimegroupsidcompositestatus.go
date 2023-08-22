// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRuntimeGroupsIDCompositeStatusRequest struct {
	// ID of a composite runtime group
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus - The HTTP status code.
type GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus int64

const (
	GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatusFiveHundredAndThree GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus = 503
)

func (e GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus) ToPointer() *GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus {
	return &e
}

func (e *GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus: %v", v)
	}
}

// GetRuntimeGroupsIDCompositeStatusServiceUnavailableError - Error response for temporary service unavailability.
type GetRuntimeGroupsIDCompositeStatusServiceUnavailableError struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// The HTTP status code.
	Status GetRuntimeGroupsIDCompositeStatusServiceUnavailableErrorStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
}

type GetRuntimeGroupsIDCompositeStatusResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Status of a composite runtime group, including existing conflicts.
	CompositeStatus *shared.CompositeStatus
	ContentType     string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	StatusCode    int
	RawResponse   *http.Response
	// Service Unavailable
	ServiceUnavailableError *GetRuntimeGroupsIDCompositeStatusServiceUnavailableError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}
