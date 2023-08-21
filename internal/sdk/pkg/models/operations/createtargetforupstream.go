// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreateTargetForUpstreamRequestBodyUpstream - The unique identifier or the name of the upstream for which to update the target.
type CreateTargetForUpstreamRequestBodyUpstream struct {
	// The unique identifier or the name of the upstream for which to update the target.
	ID *string `json:"id,omitempty"`
}

type CreateTargetForUpstreamRequestBody struct {
	// An optional set of strings associated with the Target for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The target IP address
	Target *string `json:"target,omitempty"`
	// The unique identifier or the name of the upstream for which to update the target.
	//
	Upstream *CreateTargetForUpstreamRequestBodyUpstream `json:"upstream,omitempty"`
	// The weight this target gets within the upstream loadbalancer (`0`-`65535`). If the hostname resolves to an SRV record, the `weight` value will be overridden by the value from the DNS record.
	Weight *int64 `json:"weight,omitempty"`
}

type CreateTargetForUpstreamRequest struct {
	RequestBody *CreateTargetForUpstreamRequestBody `request:"mediaType=application/json"`
	// The ID of your runtime group. This variable is available in the Konnect manager
	RuntimeGroupID string `pathParam:"style=simple,explode=false,name=runtimeGroupId"`
	// ID or name of the related Upstream
	UpstreamID string `pathParam:"style=simple,explode=false,name=upstream_id"`
}

// CreateTargetForUpstream400ApplicationJSON - Invalid Target
type CreateTargetForUpstream400ApplicationJSON struct {
}

type CreateTargetForUpstreamResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully created Target
	Target *shared.Target
	// Invalid Target
	CreateTargetForUpstream400ApplicationJSONObject *CreateTargetForUpstream400ApplicationJSON
}
