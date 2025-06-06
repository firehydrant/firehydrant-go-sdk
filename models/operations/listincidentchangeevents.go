// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
)

// ListIncidentChangeEventsType - The type of the relation to the incident
type ListIncidentChangeEventsType string

const (
	ListIncidentChangeEventsTypeCaused    ListIncidentChangeEventsType = "caused"
	ListIncidentChangeEventsTypeFixed     ListIncidentChangeEventsType = "fixed"
	ListIncidentChangeEventsTypeSuspect   ListIncidentChangeEventsType = "suspect"
	ListIncidentChangeEventsTypeDismissed ListIncidentChangeEventsType = "dismissed"
)

func (e ListIncidentChangeEventsType) ToPointer() *ListIncidentChangeEventsType {
	return &e
}
func (e *ListIncidentChangeEventsType) UnmarshalJSON(data []byte) error {
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
		*e = ListIncidentChangeEventsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListIncidentChangeEventsType: %v", v)
	}
}

type ListIncidentChangeEventsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// The type of the relation to the incident
	Type       *ListIncidentChangeEventsType `queryParam:"style=form,explode=true,name=type"`
	IncidentID string                        `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *ListIncidentChangeEventsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListIncidentChangeEventsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListIncidentChangeEventsRequest) GetType() *ListIncidentChangeEventsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListIncidentChangeEventsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}
