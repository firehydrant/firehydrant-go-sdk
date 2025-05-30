// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type UpdateIncidentChangeEventType string

const (
	UpdateIncidentChangeEventTypeCaused    UpdateIncidentChangeEventType = "caused"
	UpdateIncidentChangeEventTypeFixed     UpdateIncidentChangeEventType = "fixed"
	UpdateIncidentChangeEventTypeSuspect   UpdateIncidentChangeEventType = "suspect"
	UpdateIncidentChangeEventTypeDismissed UpdateIncidentChangeEventType = "dismissed"
)

func (e UpdateIncidentChangeEventType) ToPointer() *UpdateIncidentChangeEventType {
	return &e
}
func (e *UpdateIncidentChangeEventType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "caused":
		fallthrough
	case "fixed":
		fallthrough
	case "suspect":
		fallthrough
	case "dismissed":
		*e = UpdateIncidentChangeEventType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateIncidentChangeEventType: %v", v)
	}
}

// UpdateIncidentChangeEvent - Update a change attached to an incident
type UpdateIncidentChangeEvent struct {
	Type *UpdateIncidentChangeEventType `json:"type,omitempty"`
	// A short description about why this change event is related
	Why *string `json:"why,omitempty"`
}

func (o *UpdateIncidentChangeEvent) GetType() *UpdateIncidentChangeEventType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateIncidentChangeEvent) GetWhy() *string {
	if o == nil {
		return nil
	}
	return o.Why
}
