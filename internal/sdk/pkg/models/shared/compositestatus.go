// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// CompositeStatusState - The state of the composite runtime group.
type CompositeStatusState string

const (
	CompositeStatusStateCompositeStateOk       CompositeStatusState = "COMPOSITE_STATE_OK"
	CompositeStatusStateCompositeStateConflict CompositeStatusState = "COMPOSITE_STATE_CONFLICT"
	CompositeStatusStateCompositeStateUnknown  CompositeStatusState = "COMPOSITE_STATE_UNKNOWN"
)

func (e CompositeStatusState) ToPointer() *CompositeStatusState {
	return &e
}

func (e *CompositeStatusState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPOSITE_STATE_OK":
		fallthrough
	case "COMPOSITE_STATE_CONFLICT":
		fallthrough
	case "COMPOSITE_STATE_UNKNOWN":
		*e = CompositeStatusState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompositeStatusState: %v", v)
	}
}

// CompositeStatus - The Composite Status object contains information about the status of a composite runtime group.
type CompositeStatus struct {
	Conflicts []CompositeConflict `json:"conflicts,omitempty"`
	// An ISO-8604 timestamp representation of composite runtime group status creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The composite runtime group ID.
	ID *string `json:"id,omitempty"`
	// The state of the composite runtime group.
	State *CompositeStatusState `json:"state,omitempty"`
	// An ISO-8604 timestamp representation of composite runtime group status update date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
