// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateTeamOnCallScheduleShiftRequest struct {
	TeamID                                           string                                                      `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID                                       string                                                      `pathParam:"style=simple,explode=false,name=schedule_id"`
	PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts components.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts `request:"mediaType=application/json"`
}

func (o *CreateTeamOnCallScheduleShiftRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *CreateTeamOnCallScheduleShiftRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

func (o *CreateTeamOnCallScheduleShiftRequest) GetPostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts() components.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts {
	if o == nil {
		return components.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts{}
	}
	return o.PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts
}

type CreateTeamOnCallScheduleShiftResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateTeamOnCallScheduleShiftResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}