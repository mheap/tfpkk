// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetReportsRequest struct {
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
}

// GetReports429ApplicationProblemPlusJSON - The error response object.
type GetReports429ApplicationProblemPlusJSON struct {
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
	// The Konnect traceback code
	Instance *string `json:"instance,omitempty"`
	// The HTTP status code.
	Status *int64 `json:"status,omitempty"`
	// The error response code.
	Title *string `json:"title,omitempty"`
	// Documentation for this error.
	Type *string `json:"type,omitempty"`
}

// GetReports403ApplicationProblemPlusJSON - The error object.
type GetReports403ApplicationProblemPlusJSON struct {
	// Information about the error response.
	Detail *string `json:"detail,omitempty"`
	// Konnect traceback error code.
	Instance *string `json:"instance,omitempty"`
	// HTTP status code.
	Status *int64 `json:"status,omitempty"`
	// HTTP status code
	Title *string `json:"title,omitempty"`
	// Documentation for this error.
	Type *string `json:"type,omitempty"`
}

// GetReports200ApplicationJSON - A paginated list response for a collection of reports.
type GetReports200ApplicationJSON struct {
	Data []shared.Report `json:"data,omitempty"`
	// returns the pagination information
	Meta *shared.PaginatedMeta `json:"meta,omitempty"`
}

type GetReportsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A paginated list response for a collection of reports.
	GetReports200ApplicationJSONObject *GetReports200ApplicationJSON
	// Permission denied
	GetReports403ApplicationProblemPlusJSONObject *GetReports403ApplicationProblemPlusJSON
	// Too many requests
	GetReports429ApplicationProblemPlusJSONObject *GetReports429ApplicationProblemPlusJSON
}
