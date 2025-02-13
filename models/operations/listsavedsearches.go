// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/models/components"
	"fmt"
)

type ResourceType string

const (
	ResourceTypeChangeEvents          ResourceType = "change_events"
	ResourceTypeIncidents             ResourceType = "incidents"
	ResourceTypeServices              ResourceType = "services"
	ResourceTypeScheduledMaintenances ResourceType = "scheduled_maintenances"
	ResourceTypeTicketTasks           ResourceType = "ticket_tasks"
	ResourceTypeTicketFollowUps       ResourceType = "ticket_follow_ups"
	ResourceTypeAnalytics             ResourceType = "analytics"
	ResourceTypeImpactAnalytics       ResourceType = "impact_analytics"
	ResourceTypeAlerts                ResourceType = "alerts"
	ResourceTypeIncidentEvents        ResourceType = "incident_events"
)

func (e ResourceType) ToPointer() *ResourceType {
	return &e
}
func (e *ResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "change_events":
		fallthrough
	case "incidents":
		fallthrough
	case "services":
		fallthrough
	case "scheduled_maintenances":
		fallthrough
	case "ticket_tasks":
		fallthrough
	case "ticket_follow_ups":
		fallthrough
	case "analytics":
		fallthrough
	case "impact_analytics":
		fallthrough
	case "alerts":
		fallthrough
	case "incident_events":
		*e = ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResourceType: %v", v)
	}
}

type ListSavedSearchesRequest struct {
	ResourceType ResourceType `pathParam:"style=simple,explode=false,name=resource_type"`
	// The user ID used to filter saved searches.
	UserID *string `queryParam:"style=form,explode=true,name=user_id"`
	// Filter saved searches with a query on their name
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListSavedSearchesRequest) GetResourceType() ResourceType {
	if o == nil {
		return ResourceType("")
	}
	return o.ResourceType
}

func (o *ListSavedSearchesRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *ListSavedSearchesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListSavedSearchesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListSavedSearchesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListSavedSearchesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists save searches
	SavedSearchEntity *components.SavedSearchEntity
}

func (o *ListSavedSearchesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListSavedSearchesResponse) GetSavedSearchEntity() *components.SavedSearchEntity {
	if o == nil {
		return nil
	}
	return o.SavedSearchEntity
}
