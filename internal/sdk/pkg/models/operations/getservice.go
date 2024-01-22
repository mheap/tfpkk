// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
	"net/http"
)

type GetServiceRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	FilterTags *string `queryParam:"style=form,explode=true,name=tags"`
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// UUID of the service to lookup
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
}

func (g GetServiceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetServiceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetServiceRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetServiceRequest) GetFilterTags() *string {
	if o == nil {
		return nil
	}
	return o.FilterTags
}

func (o *GetServiceRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetServiceRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetServiceRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

// GetServiceResponseBody - The service response object is returned when creating a new service or retreiving an existing service.
type GetServiceResponseBody struct {
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Unix epoch when the resource was last created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Service enabled boolean
	Enabled      *bool   `json:"enabled,omitempty"`
	Host         *string `json:"host,omitempty"`
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Path         *string `json:"path,omitempty"`
	Port         *int64  `json:"port,omitempty"`
	Protocol     *string `json:"protocol,omitempty"`
	ReadTimeout  *int64  `json:"read_timeout,omitempty"`
	Retries      *int64  `json:"retries,omitempty"`
	UpdatedAt    *int64  `json:"updated_at,omitempty"`
	WriteTimeout *int64  `json:"write_timeout,omitempty"`
}

func (o *GetServiceResponseBody) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *GetServiceResponseBody) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetServiceResponseBody) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *GetServiceResponseBody) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *GetServiceResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetServiceResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetServiceResponseBody) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *GetServiceResponseBody) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *GetServiceResponseBody) GetProtocol() *string {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *GetServiceResponseBody) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *GetServiceResponseBody) GetRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *GetServiceResponseBody) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetServiceResponseBody) GetWriteTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteTimeout
}

type GetServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The service response object is returned when creating a new service or retreiving an existing service.
	Object *GetServiceResponseBody
}

func (o *GetServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServiceResponse) GetObject() *GetServiceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
