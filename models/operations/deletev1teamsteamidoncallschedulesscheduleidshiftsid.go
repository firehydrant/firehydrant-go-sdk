// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest struct {
	ID         string `pathParam:"style=simple,explode=false,name=id"`
	TeamID     string `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule_id"`
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}