// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateAPIProductVersionDTOPublishStatus - The publish status of the API product version.
type UpdateAPIProductVersionDTOPublishStatus string

const (
	UpdateAPIProductVersionDTOPublishStatusUnpublished UpdateAPIProductVersionDTOPublishStatus = "unpublished"
	UpdateAPIProductVersionDTOPublishStatusPublished   UpdateAPIProductVersionDTOPublishStatus = "published"
)

func (e UpdateAPIProductVersionDTOPublishStatus) ToPointer() *UpdateAPIProductVersionDTOPublishStatus {
	return &e
}

func (e *UpdateAPIProductVersionDTOPublishStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unpublished":
		fallthrough
	case "published":
		*e = UpdateAPIProductVersionDTOPublishStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateAPIProductVersionDTOPublishStatus: %v", v)
	}
}

// UpdateAPIProductVersionDTO - The request schema for updating a version of an API product.
type UpdateAPIProductVersionDTO struct {
	// Indicates if the version of the API product is deprecated.
	Deprecated     *bool                  `json:"deprecated,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
	// The version name of the API product version.
	Name *string `json:"name,omitempty"`
	// When set to `true`, and all the following conditions are true:
	// - version of the API product deprecation has changed from `false` -> `true`
	// - version of the API product is published
	//
	// then consumers of the now deprecated verion of the API product will be notified.
	//
	Notify *bool `json:"notify,omitempty"`
	// The publish status of the API product version.
	PublishStatus *UpdateAPIProductVersionDTOPublishStatus `json:"publish_status,omitempty"`
}

func (o *UpdateAPIProductVersionDTO) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *UpdateAPIProductVersionDTO) GetGatewayService() *GatewayServicePayload {
	if o == nil {
		return nil
	}
	return o.GatewayService
}

func (o *UpdateAPIProductVersionDTO) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateAPIProductVersionDTO) GetNotify() *bool {
	if o == nil {
		return nil
	}
	return o.Notify
}

func (o *UpdateAPIProductVersionDTO) GetPublishStatus() *UpdateAPIProductVersionDTOPublishStatus {
	if o == nil {
		return nil
	}
	return o.PublishStatus
}
