// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1IncidentsIncidentIDRoleAssignmentsRequest struct {
	IncidentID                               string                                              `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDRoleAssignments components.PostV1IncidentsIncidentIDRoleAssignments `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsRequest) GetPostV1IncidentsIncidentIDRoleAssignments() components.PostV1IncidentsIncidentIDRoleAssignments {
	if o == nil {
		return components.PostV1IncidentsIncidentIDRoleAssignments{}
	}
	return o.PostV1IncidentsIncidentIDRoleAssignments
}

type PostV1IncidentsIncidentIDRoleAssignmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Assign a role to a user for this incident
	IncidentsRoleAssignmentEntity *components.IncidentsRoleAssignmentEntity
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsResponse) GetIncidentsRoleAssignmentEntity() *components.IncidentsRoleAssignmentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRoleAssignmentEntity
}