// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type ExportCsvRequest struct {
	// The ID of the report to export
	ReportID string `pathParam:"style=simple,explode=false,name=reportId"`
	// The timezone to use for the report.  This impacts several things:
	//
	// - The timestamps returned, if time is a dimension.
	// - Day boundaries for reports, meaning that if the report implicitly has a granularity of DAILY or higher, the same query may return different data depending on the timezone selected.
	//
	Tz *string `queryParam:"style=form,explode=true,name=tz"`
	// If true, export timestamps in "programmer friendly" format (ISO-8601, UTC timezone).
	// If false, export in "spreadsheet friendly" format (local time without timezone specifier).
	//
	// Note: this does NOT impact the timezone of the report, merely the format of the time column if one is present.
	//
	UtcFormat *bool `queryParam:"style=form,explode=true,name=utc_format"`
}

// ExportCsv429ApplicationProblemPlusJSON - The error response object.
type ExportCsv429ApplicationProblemPlusJSON struct {
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

// ExportCsv403ApplicationProblemPlusJSON - The error object.
type ExportCsv403ApplicationProblemPlusJSON struct {
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

type ExportCsvResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	ContentType     string
	// Not Found
	NotFoundError *shared.NotFoundError
	StatusCode    int
	RawResponse   *http.Response
	// Report exported successfully
	ExportCsv200TextCsvString *string
	// Permission denied
	ExportCsv403ApplicationProblemPlusJSONObject *ExportCsv403ApplicationProblemPlusJSON
	// Too many requests
	ExportCsv429ApplicationProblemPlusJSONObject *ExportCsv429ApplicationProblemPlusJSON
}
