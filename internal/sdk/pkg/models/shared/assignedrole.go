// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AssignedRoleEntityRegion - Region of the entity.
type AssignedRoleEntityRegion string

const (
	AssignedRoleEntityRegionUs       AssignedRoleEntityRegion = "us"
	AssignedRoleEntityRegionEu       AssignedRoleEntityRegion = "eu"
	AssignedRoleEntityRegionWildcard AssignedRoleEntityRegion = "*"
)

func (e AssignedRoleEntityRegion) ToPointer() *AssignedRoleEntityRegion {
	return &e
}

func (e *AssignedRoleEntityRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us":
		fallthrough
	case "eu":
		fallthrough
	case "*":
		*e = AssignedRoleEntityRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssignedRoleEntityRegion: %v", v)
	}
}

// AssignedRole - An assigned role is a role that has been assigned to a user or team.
type AssignedRole struct {
	// A RBAC entity ID.
	EntityID *string `json:"entity_id,omitempty"`
	// Region of the entity.
	EntityRegion *AssignedRoleEntityRegion `json:"entity_region,omitempty"`
	// Name of the entity type the role is being assigned to.
	EntityTypeName *string `json:"entity_type_name,omitempty"`
	// The ID of the role assignment.
	ID *string `json:"id,omitempty"`
	// Name of the role being assigned.
	RoleName *string `json:"role_name,omitempty"`
}
