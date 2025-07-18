// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type UpdateOnCallShiftRequest struct {
	ID                string                       `pathParam:"style=simple,explode=false,name=id"`
	TeamID            string                       `pathParam:"style=simple,explode=false,name=team_id"`
	ScheduleID        string                       `pathParam:"style=simple,explode=false,name=schedule_id"`
	UpdateOnCallShift components.UpdateOnCallShift `request:"mediaType=application/json"`
}

func (o *UpdateOnCallShiftRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateOnCallShiftRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *UpdateOnCallShiftRequest) GetScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ScheduleID
}

func (o *UpdateOnCallShiftRequest) GetUpdateOnCallShift() components.UpdateOnCallShift {
	if o == nil {
		return components.UpdateOnCallShift{}
	}
	return o.UpdateOnCallShift
}
