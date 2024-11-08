// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"firehydrant/types"
)

type GetMeanTimeReportRequest struct {
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// The start date to return incidents from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return incidents from
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A comma separated list of priorities
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A comma separated list of severities
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
}

func (g GetMeanTimeReportRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeanTimeReportRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeanTimeReportRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetMeanTimeReportRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetMeanTimeReportRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetMeanTimeReportRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetMeanTimeReportRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetMeanTimeReportRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *GetMeanTimeReportRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetMeanTimeReportRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *GetMeanTimeReportRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *GetMeanTimeReportRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *GetMeanTimeReportRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *GetMeanTimeReportRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *GetMeanTimeReportRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

type GetMeanTimeReportResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a report with time bucketed analytics data
	ReportEntity *components.ReportEntity
}

func (o *GetMeanTimeReportResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMeanTimeReportResponse) GetReportEntity() *components.ReportEntity {
	if o == nil {
		return nil
	}
	return o.ReportEntity
}