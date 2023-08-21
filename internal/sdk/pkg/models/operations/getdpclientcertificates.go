// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetDpClientCertificatesRequest struct {
	// The specific page number in the collection results.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page_number"`
	// The number of items to include in a page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
}

type GetDpClientCertificates200ApplicationJSONItems struct {
	// JSON escaped string of the certificate.
	Cert *string `json:"cert,omitempty"`
	// Date certificate was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unique ID of the certificate entity.
	ID *string `json:"id,omitempty"`
	// Date certificate was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

type GetDpClientCertificates200ApplicationJSONPage struct {
	TotalCount *int64 `json:"total_count,omitempty"`
}

// GetDpClientCertificates200ApplicationJSON - Example response
type GetDpClientCertificates200ApplicationJSON struct {
	Items []GetDpClientCertificates200ApplicationJSONItems `json:"items,omitempty"`
	Page  *GetDpClientCertificates200ApplicationJSONPage   `json:"page,omitempty"`
}

type GetDpClientCertificatesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Example response
	GetDpClientCertificates200ApplicationJSONObject *GetDpClientCertificates200ApplicationJSON
}
