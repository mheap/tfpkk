// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// APIProductFilterDimension - The dimension to filter.
type APIProductFilterDimension string

const (
	APIProductFilterDimensionAPIProduct APIProductFilterDimension = "API_PRODUCT"
)

func (e APIProductFilterDimension) ToPointer() *APIProductFilterDimension {
	return &e
}

func (e *APIProductFilterDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API_PRODUCT":
		*e = APIProductFilterDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIProductFilterDimension: %v", v)
	}
}

type APIProductFilter struct {
	// The dimension to filter.
	Dimension APIProductFilterDimension `json:"dimension"`
	// The type of filter to apply.  `IN` filters will limit results to only the specified values, while `NOT_IN` filters will exclude the specified values.
	Type FilterType `json:"type"`
	// The UUIDs of the API products to include in the results.
	//
	Values []string `json:"values"`
}
