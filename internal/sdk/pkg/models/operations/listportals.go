// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListPortalsRequest struct {
	// Determines which page of the entities to retrieve.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// The maximum number of items to include per page. The last page of a collection may include fewer items.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// Sorts a collection of portals. Supported sort attributes are:
	//   - created_at
	//   - updated_at
	//
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
}

func (o *ListPortalsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListPortalsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListPortalsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListPortalsResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// HTTP response content type for this operation
	ContentType string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// A paginated list of portals in the current region of an organization.
	ListPortalsResponse *shared.ListPortalsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *ListPortalsResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *ListPortalsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPortalsResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *ListPortalsResponse) GetListPortalsResponse() *shared.ListPortalsResponse {
	if o == nil {
		return nil
	}
	return o.ListPortalsResponse
}

func (o *ListPortalsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPortalsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListPortalsResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
