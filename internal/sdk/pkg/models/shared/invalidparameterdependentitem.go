// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InvalidParameterDependentItemRule - invalid parameters rules
type InvalidParameterDependentItemRule string

const (
	InvalidParameterDependentItemRuleDependentFields InvalidParameterDependentItemRule = "dependent_fields"
)

func (e InvalidParameterDependentItemRule) ToPointer() *InvalidParameterDependentItemRule {
	return &e
}

func (e *InvalidParameterDependentItemRule) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dependent_fields":
		*e = InvalidParameterDependentItemRule(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvalidParameterDependentItemRule: %v", v)
	}
}

type InvalidParameterDependentItem struct {
	Dependents []interface{} `json:"dependents"`
	Field      string        `json:"field"`
	Reason     string        `json:"reason"`
	// invalid parameters rules
	Rule *InvalidParameterDependentItemRule `json:"rule"`
}