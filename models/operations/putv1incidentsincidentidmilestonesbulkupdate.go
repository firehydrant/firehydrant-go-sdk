// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PutV1IncidentsIncidentIDMilestonesBulkUpdateRequest struct {
	IncidentID                                   string                                                  `pathParam:"style=simple,explode=false,name=incident_id"`
	PutV1IncidentsIncidentIDMilestonesBulkUpdate components.PutV1IncidentsIncidentIDMilestonesBulkUpdate `request:"mediaType=application/json"`
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateRequest) GetPutV1IncidentsIncidentIDMilestonesBulkUpdate() components.PutV1IncidentsIncidentIDMilestonesBulkUpdate {
	if o == nil {
		return components.PutV1IncidentsIncidentIDMilestonesBulkUpdate{}
	}
	return o.PutV1IncidentsIncidentIDMilestonesBulkUpdate
}

type PutV1IncidentsIncidentIDMilestonesBulkUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update milestone times in bulk for a given incident. All milestone
	// times for an incident must occur in chronological order
	// corresponding to the configured order of milestones. If the result
	// of this request would cause any milestone(s) to appear out of place,
	// a 422 response will instead be returned. This includes milestones
	// not explicitly submitted or updated in this request.
	//
	IncidentsMilestoneEntityPaginated *components.IncidentsMilestoneEntityPaginated
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateResponse) GetIncidentsMilestoneEntityPaginated() *components.IncidentsMilestoneEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsMilestoneEntityPaginated
}