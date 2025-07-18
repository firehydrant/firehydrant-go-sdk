// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// GetSignalsMttxAnalyticsGroupBy - String that determines how records are grouped
type GetSignalsMttxAnalyticsGroupBy string

const (
	GetSignalsMttxAnalyticsGroupBySignalRules  GetSignalsMttxAnalyticsGroupBy = "signal_rules"
	GetSignalsMttxAnalyticsGroupByTeams        GetSignalsMttxAnalyticsGroupBy = "teams"
	GetSignalsMttxAnalyticsGroupByServices     GetSignalsMttxAnalyticsGroupBy = "services"
	GetSignalsMttxAnalyticsGroupByEnvironments GetSignalsMttxAnalyticsGroupBy = "environments"
	GetSignalsMttxAnalyticsGroupByTags         GetSignalsMttxAnalyticsGroupBy = "tags"
)

func (e GetSignalsMttxAnalyticsGroupBy) ToPointer() *GetSignalsMttxAnalyticsGroupBy {
	return &e
}
func (e *GetSignalsMttxAnalyticsGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "signal_rules":
		fallthrough
	case "teams":
		fallthrough
	case "services":
		fallthrough
	case "environments":
		fallthrough
	case "tags":
		*e = GetSignalsMttxAnalyticsGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSignalsMttxAnalyticsGroupBy: %v", v)
	}
}

// GetSignalsMttxAnalyticsSortBy - String that determines how records are sorted
type GetSignalsMttxAnalyticsSortBy string

const (
	GetSignalsMttxAnalyticsSortByTotalOpenedAlerts   GetSignalsMttxAnalyticsSortBy = "total_opened_alerts"
	GetSignalsMttxAnalyticsSortByTotalAckedAlerts    GetSignalsMttxAnalyticsSortBy = "total_acked_alerts"
	GetSignalsMttxAnalyticsSortByTotalIncidents      GetSignalsMttxAnalyticsSortBy = "total_incidents"
	GetSignalsMttxAnalyticsSortByAckedPercentage     GetSignalsMttxAnalyticsSortBy = "acked_percentage"
	GetSignalsMttxAnalyticsSortByIncidentsPercentage GetSignalsMttxAnalyticsSortBy = "incidents_percentage"
)

func (e GetSignalsMttxAnalyticsSortBy) ToPointer() *GetSignalsMttxAnalyticsSortBy {
	return &e
}
func (e *GetSignalsMttxAnalyticsSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total_opened_alerts":
		fallthrough
	case "total_acked_alerts":
		fallthrough
	case "total_incidents":
		fallthrough
	case "acked_percentage":
		fallthrough
	case "incidents_percentage":
		*e = GetSignalsMttxAnalyticsSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSignalsMttxAnalyticsSortBy: %v", v)
	}
}

// GetSignalsMttxAnalyticsSortDirection - String that determines how records are sorted
type GetSignalsMttxAnalyticsSortDirection string

const (
	GetSignalsMttxAnalyticsSortDirectionAsc  GetSignalsMttxAnalyticsSortDirection = "asc"
	GetSignalsMttxAnalyticsSortDirectionDesc GetSignalsMttxAnalyticsSortDirection = "desc"
)

func (e GetSignalsMttxAnalyticsSortDirection) ToPointer() *GetSignalsMttxAnalyticsSortDirection {
	return &e
}
func (e *GetSignalsMttxAnalyticsSortDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSignalsMttxAnalyticsSortDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSignalsMttxAnalyticsSortDirection: %v", v)
	}
}

type GetSignalsMttxAnalyticsRequest struct {
	// A comma separated list of signal rule IDs
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A comma separated list of user IDs
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// String that determines how records are grouped
	GroupBy *GetSignalsMttxAnalyticsGroupBy `queryParam:"style=form,explode=true,name=group_by"`
	// String that determines how records are sorted
	SortBy *GetSignalsMttxAnalyticsSortBy `queryParam:"style=form,explode=true,name=sort_by"`
	// String that determines how records are sorted
	SortDirection *GetSignalsMttxAnalyticsSortDirection `queryParam:"style=form,explode=true,name=sort_direction"`
	// The start date to return metrics from
	StartDate *time.Time `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *time.Time `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetSignalsMttxAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSignalsMttxAnalyticsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetSignalsMttxAnalyticsRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *GetSignalsMttxAnalyticsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetSignalsMttxAnalyticsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetSignalsMttxAnalyticsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetSignalsMttxAnalyticsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetSignalsMttxAnalyticsRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *GetSignalsMttxAnalyticsRequest) GetGroupBy() *GetSignalsMttxAnalyticsGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *GetSignalsMttxAnalyticsRequest) GetSortBy() *GetSignalsMttxAnalyticsSortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *GetSignalsMttxAnalyticsRequest) GetSortDirection() *GetSignalsMttxAnalyticsSortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *GetSignalsMttxAnalyticsRequest) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetSignalsMttxAnalyticsRequest) GetEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndDate
}
