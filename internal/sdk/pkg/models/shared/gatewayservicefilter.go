// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GatewayServiceFilterDimension - The dimension to filter.
type GatewayServiceFilterDimension string

const (
	GatewayServiceFilterDimensionGatewayService GatewayServiceFilterDimension = "GATEWAY_SERVICE"
)

func (e GatewayServiceFilterDimension) ToPointer() *GatewayServiceFilterDimension {
	return &e
}

func (e *GatewayServiceFilterDimension) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GATEWAY_SERVICE":
		*e = GatewayServiceFilterDimension(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GatewayServiceFilterDimension: %v", v)
	}
}

type GatewayServiceFilter struct {
	// The dimension to filter.
	Dimension GatewayServiceFilterDimension `json:"dimension"`
	// The type of filter to apply.  `IN` filters will limit results to only the specified values, while `NOT_IN` filters will exclude the specified values.
	Type FilterType `json:"type"`
	// The gateway services to include in the results.  Because gateway service UUIDs are only unique within a given runtime group, the filter values must be of the form "[runtime group UUID]:[service UUID]".
	//
	Values []string `json:"values"`
}
