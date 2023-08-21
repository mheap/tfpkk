// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RuntimeGroupFilterDimension - The dimension to filter.
type RuntimeGroupFilterDimension string

const (
	RuntimeGroupFilterDimensionRuntimeGroup RuntimeGroupFilterDimension = "RUNTIME_GROUP"
)

func (e RuntimeGroupFilterDimension) ToPointer() *RuntimeGroupFilterDimension {
	return &e
}

func (e *RuntimeGroupFilterDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RUNTIME_GROUP":
		*e = RuntimeGroupFilterDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RuntimeGroupFilterDimension: %v", v)
	}
}

type RuntimeGroupFilter struct {
	// The dimension to filter.
	Dimension RuntimeGroupFilterDimension `json:"dimension"`
	// The type of filter to apply.  `IN` filters will limit results to only the specified values, while `NOT_IN` filters will exclude the specified values.
	Type FilterType `json:"type"`
	// The UUIDs of the runtime groups to include in the results.
	//
	Values []string `json:"values"`
}
