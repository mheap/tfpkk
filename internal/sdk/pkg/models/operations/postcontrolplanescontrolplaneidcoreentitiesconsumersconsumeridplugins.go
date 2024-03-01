// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsRequest struct {
	// ID of the Consumer to lookup
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerId"`
	// Description of new Plugin for creation
	CreatePluginWithoutParents shared.CreatePluginWithoutParents `request:"mediaType=application/json"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsRequest) GetCreatePluginWithoutParents() shared.CreatePluginWithoutParents {
	if o == nil {
		return shared.CreatePluginWithoutParents{}
	}
	return o.CreatePluginWithoutParents
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

// PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponseBody - Invalid Plugin
type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponseBody struct {
}

type PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully created Plugin
	Plugin *shared.Plugin
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid Plugin
	Object *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponseBody
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse) GetPlugin() *shared.Plugin {
	if o == nil {
		return nil
	}
	return o.Plugin
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponse) GetObject() *PostControlPlanesControlPlaneIDCoreEntitiesConsumersConsumerIDPluginsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}