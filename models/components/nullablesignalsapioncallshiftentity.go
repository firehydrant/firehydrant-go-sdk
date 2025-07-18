// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// NullableSignalsAPIOnCallShiftEntity - Signals_API_OnCallShiftEntity model
type NullableSignalsAPIOnCallShiftEntity struct {
	ID              *string                 `json:"id,omitempty"`
	User            *NullableSuccinctEntity `json:"user,omitempty"`
	CoverageRequest *string                 `json:"coverage_request,omitempty"`
	Color           *string                 `json:"color,omitempty"`
	TimeZone        *string                 `json:"time_zone,omitempty"`
	OnCallSchedule  *NullableSuccinctEntity `json:"on_call_schedule,omitempty"`
	OnCallRotation  *NullableSuccinctEntity `json:"on_call_rotation,omitempty"`
	Team            *NullableSuccinctEntity `json:"team,omitempty"`
	StartTime       *time.Time              `json:"start_time,omitempty"`
	EndTime         *time.Time              `json:"end_time,omitempty"`
}

func (n NullableSignalsAPIOnCallShiftEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NullableSignalsAPIOnCallShiftEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetUser() *NullableSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetCoverageRequest() *string {
	if o == nil {
		return nil
	}
	return o.CoverageRequest
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.TimeZone
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetOnCallSchedule() *NullableSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.OnCallSchedule
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetOnCallRotation() *NullableSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.OnCallRotation
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetTeam() *NullableSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Team
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *NullableSignalsAPIOnCallShiftEntity) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}
