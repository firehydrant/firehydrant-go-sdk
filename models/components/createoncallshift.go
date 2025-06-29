// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateOnCallShift - Create a Signals on-call shift in a schedule.
type CreateOnCallShift struct {
	// The start time of the shift in ISO8601 format.
	StartTime string `json:"start_time"`
	// The end time of the shift in ISO8601 format.
	EndTime string `json:"end_time"`
	// The ID of the user who is on-call for the shift. If not provided, the shift will be unassigned.
	UserID *string `json:"user_id,omitempty"`
	// The ID of the on-call rotation you want to create the shift in. This parameter is optional for backwards compatibility but must be provided if the schedule has multiple rotations.
	RotationID *string `json:"rotation_id,omitempty"`
}

func (o *CreateOnCallShift) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

func (o *CreateOnCallShift) GetEndTime() string {
	if o == nil {
		return ""
	}
	return o.EndTime
}

func (o *CreateOnCallShift) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *CreateOnCallShift) GetRotationID() *string {
	if o == nil {
		return nil
	}
	return o.RotationID
}
