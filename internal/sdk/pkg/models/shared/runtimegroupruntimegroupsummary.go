// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RuntimeGroupRuntimeGroupSummary - The runtime group object contains information about a Kong control plane.
type RuntimeGroupRuntimeGroupSummary struct {
	// CP configuration object for related access endpoints.
	Config *RuntimeGroupConfig `json:"config,omitempty"`
	// An ISO-8604 timestamp representation of runtime group creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The description of the runtime group in Konnect.
	Description *string `json:"description,omitempty"`
	// The runtime group ID.
	ID *string `json:"id,omitempty"`
	// Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'.
	Labels interface{} `json:"labels,omitempty"`
	// The name of the runtime group.
	Name *string `json:"name,omitempty"`
	// An ISO-8604 timestamp representation of runtime group update date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
