// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteIncidentRoleAssignmentRequest struct {
	IncidentID       string `pathParam:"style=simple,explode=false,name=incident_id"`
	RoleAssignmentID string `pathParam:"style=simple,explode=false,name=role_assignment_id"`
}

func (o *DeleteIncidentRoleAssignmentRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *DeleteIncidentRoleAssignmentRequest) GetRoleAssignmentID() string {
	if o == nil {
		return ""
	}
	return o.RoleAssignmentID
}

type DeleteIncidentRoleAssignmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Unassign a role from a user
	IncidentsRoleAssignmentEntity *components.IncidentsRoleAssignmentEntity
}

func (o *DeleteIncidentRoleAssignmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteIncidentRoleAssignmentResponse) GetIncidentsRoleAssignmentEntity() *components.IncidentsRoleAssignmentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRoleAssignmentEntity
}