// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Konnect/internal/sdk/pkg/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type PortalTeamsAssignRoleRequestBodyEntityRegion string

const (
	PortalTeamsAssignRoleRequestBodyEntityRegionUs       PortalTeamsAssignRoleRequestBodyEntityRegion = "us"
	PortalTeamsAssignRoleRequestBodyEntityRegionEu       PortalTeamsAssignRoleRequestBodyEntityRegion = "eu"
	PortalTeamsAssignRoleRequestBodyEntityRegionWildcard PortalTeamsAssignRoleRequestBodyEntityRegion = "*"
)

func (e PortalTeamsAssignRoleRequestBodyEntityRegion) ToPointer() *PortalTeamsAssignRoleRequestBodyEntityRegion {
	return &e
}

func (e *PortalTeamsAssignRoleRequestBodyEntityRegion) UnmarshalJSON(data []byte) error {
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
		*e = PortalTeamsAssignRoleRequestBodyEntityRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PortalTeamsAssignRoleRequestBodyEntityRegion: %v", v)
	}
}

// PortalTeamsAssignRoleRequestBody - An assigned role associates a service and an action to a team.
type PortalTeamsAssignRoleRequestBody struct {
	EntityID       *string                                       `json:"entity_id,omitempty"`
	EntityRegion   *PortalTeamsAssignRoleRequestBodyEntityRegion `json:"entity_region,omitempty"`
	EntityTypeName *string                                       `json:"entity_type_name,omitempty"`
	RoleName       *string                                       `json:"role_name,omitempty"`
}

type PortalTeamsAssignRoleRequest struct {
	// Assign a role to a team.
	RequestBody *PortalTeamsAssignRoleRequestBody `request:"mediaType=application/json"`
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// ID of the team.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
}

type PortalTeamsAssignRoleResponse struct {
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Conflict
	ConflictError *shared.ConflictError
	ContentType   string
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
	// A get action response of a single assigned role.
	PortalAssignedRole *shared.PortalAssignedRole
	StatusCode         int
	RawResponse        *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}
