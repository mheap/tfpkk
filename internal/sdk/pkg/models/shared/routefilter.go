// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RouteFilterDimension - The dimension to filter.
type RouteFilterDimension string

const (
	RouteFilterDimensionRoute RouteFilterDimension = "ROUTE"
)

func (e RouteFilterDimension) ToPointer() *RouteFilterDimension {
	return &e
}

func (e *RouteFilterDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ROUTE":
		*e = RouteFilterDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RouteFilterDimension: %v", v)
	}
}

type RouteFilter struct {
	// The dimension to filter.
	Dimension RouteFilterDimension `json:"dimension"`
	// The type of filter to apply.  `IN` filters will limit results to only the specified values, while `NOT_IN` filters will exclude the specified values.
	Type FilterType `json:"type"`
	// The routes to include in the results.  Because route UUIDs are only unique within a given runtime group, the filter values must be of the form "[runtime group UUID]:[route UUID]".
	//
	Values []string `json:"values"`
}
