// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/utils"
	"net/http"
)

type GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest struct {
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (g GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody - A successful response listing ACLs
type GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody struct {
	Data []shared.ACL `json:"data,omitempty"`
	// Offset is used to paginate through the API. Provide this value to the next list operation to fetch the next page
	Offset *string `json:"offset,omitempty"`
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody) GetData() []shared.ACL {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

type GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A successful response listing ACLs
	Object *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponse) GetObject() *GetControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDAclsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
