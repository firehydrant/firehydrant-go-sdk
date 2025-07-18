// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"github.com/firehydrant/firehydrant-go-sdk/types"
	"time"
)

// ListMttxMetricsTagMatchStrategy - A matching strategy for the tags provided
type ListMttxMetricsTagMatchStrategy string

const (
	ListMttxMetricsTagMatchStrategyAny      ListMttxMetricsTagMatchStrategy = "any"
	ListMttxMetricsTagMatchStrategyMatchAll ListMttxMetricsTagMatchStrategy = "match_all"
	ListMttxMetricsTagMatchStrategyExclude  ListMttxMetricsTagMatchStrategy = "exclude"
)

func (e ListMttxMetricsTagMatchStrategy) ToPointer() *ListMttxMetricsTagMatchStrategy {
	return &e
}
func (e *ListMttxMetricsTagMatchStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "match_all":
		fallthrough
	case "exclude":
		*e = ListMttxMetricsTagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListMttxMetricsTagMatchStrategy: %v", v)
	}
}

type ListMttxMetricsSortBy string

const (
	ListMttxMetricsSortByCountAsc        ListMttxMetricsSortBy = "count_asc"
	ListMttxMetricsSortByMttrAsc         ListMttxMetricsSortBy = "mttr_asc"
	ListMttxMetricsSortByMttaAsc         ListMttxMetricsSortBy = "mtta_asc"
	ListMttxMetricsSortByMttdAsc         ListMttxMetricsSortBy = "mttd_asc"
	ListMttxMetricsSortByMttmAsc         ListMttxMetricsSortBy = "mttm_asc"
	ListMttxMetricsSortByHealthinessAsc  ListMttxMetricsSortBy = "healthiness_asc"
	ListMttxMetricsSortByCountDesc       ListMttxMetricsSortBy = "count_desc"
	ListMttxMetricsSortByMttrDesc        ListMttxMetricsSortBy = "mttr_desc"
	ListMttxMetricsSortByMttaDesc        ListMttxMetricsSortBy = "mtta_desc"
	ListMttxMetricsSortByMttdDesc        ListMttxMetricsSortBy = "mttd_desc"
	ListMttxMetricsSortByMttmDesc        ListMttxMetricsSortBy = "mttm_desc"
	ListMttxMetricsSortByHealthinessDesc ListMttxMetricsSortBy = "healthiness_desc"
)

func (e ListMttxMetricsSortBy) ToPointer() *ListMttxMetricsSortBy {
	return &e
}
func (e *ListMttxMetricsSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count_asc":
		fallthrough
	case "mttr_asc":
		fallthrough
	case "mtta_asc":
		fallthrough
	case "mttd_asc":
		fallthrough
	case "mttm_asc":
		fallthrough
	case "healthiness_asc":
		fallthrough
	case "count_desc":
		fallthrough
	case "mttr_desc":
		fallthrough
	case "mtta_desc":
		fallthrough
	case "mttd_desc":
		fallthrough
	case "mttm_desc":
		fallthrough
	case "healthiness_desc":
		*e = ListMttxMetricsSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListMttxMetricsSortBy: %v", v)
	}
}

type ListMttxMetricsGroupBy string

const (
	ListMttxMetricsGroupByServices        ListMttxMetricsGroupBy = "services"
	ListMttxMetricsGroupByEnvironments    ListMttxMetricsGroupBy = "environments"
	ListMttxMetricsGroupByFunctionalities ListMttxMetricsGroupBy = "functionalities"
	ListMttxMetricsGroupByTeams           ListMttxMetricsGroupBy = "teams"
	ListMttxMetricsGroupBySeverities      ListMttxMetricsGroupBy = "severities"
	ListMttxMetricsGroupByUsers           ListMttxMetricsGroupBy = "users"
	ListMttxMetricsGroupByIncidentTypes   ListMttxMetricsGroupBy = "incident_types"
	ListMttxMetricsGroupByStartedDay      ListMttxMetricsGroupBy = "started_day"
	ListMttxMetricsGroupByStartedWeek     ListMttxMetricsGroupBy = "started_week"
	ListMttxMetricsGroupByStartedMonth    ListMttxMetricsGroupBy = "started_month"
	ListMttxMetricsGroupByCustomFields    ListMttxMetricsGroupBy = "custom_fields"
)

func (e ListMttxMetricsGroupBy) ToPointer() *ListMttxMetricsGroupBy {
	return &e
}
func (e *ListMttxMetricsGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "services":
		fallthrough
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "teams":
		fallthrough
	case "severities":
		fallthrough
	case "users":
		fallthrough
	case "incident_types":
		fallthrough
	case "started_day":
		fallthrough
	case "started_week":
		fallthrough
	case "started_month":
		fallthrough
	case "custom_fields":
		*e = ListMttxMetricsGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListMttxMetricsGroupBy: %v", v)
	}
}

type ListMttxMetricsRequestBody struct {
	GroupBy []ListMttxMetricsGroupBy `multipartForm:"name=group_by"`
}

func (o *ListMttxMetricsRequestBody) GetGroupBy() []ListMttxMetricsGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

type ListMttxMetricsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A JSON string that defines 'logic' and 'user_data'
	Conditions *string `queryParam:"style=form,explode=true,name=conditions"`
	// A comma separated list of environment IDs or 'is_empty' to filter for incidents with no impacted environments
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs or 'is_empty' to filter for incidents with no impacted services
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of functionality IDs or 'is_empty' to filter for incidents with no impacted functionalities
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A comma separated list of infrastructure IDs. Returns incidents that do not have the following infrastructure ids associated with them.
	ExcludedInfrastructureIds *string `queryParam:"style=form,explode=true,name=excluded_infrastructure_ids"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of IDs for assigned teams or 'is_empty' to filter for incidents with no active team assignments
	AssignedTeams *string `queryParam:"style=form,explode=true,name=assigned_teams"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// Filters for incidents that started on or after this date
	StartDate types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// Filters for incidents that started on or before this date
	EndDate types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// Filters for incidents that were resolved at or after this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_after"`
	// Filters for incidents that were resolved at or before this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_before"`
	// Filters for incidents that were closed at or after this time
	ClosedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=closed_at_or_after"`
	// Filters for incidents that were closed at or before this time
	ClosedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=closed_at_or_before"`
	// Filters for incidents that were created at or after this time
	CreatedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=created_at_or_after"`
	// Filters for incidents that were created at or before this time
	CreatedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=created_at_or_before"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search incidents by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A text value of priority
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A text value of severity
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A matching strategy for the tags provided
	TagMatchStrategy *ListMttxMetricsTagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// Return archived incidents
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Filters for incidents that were updated after this date
	UpdatedAfter *time.Time `queryParam:"style=form,explode=true,name=updated_after"`
	// Filters for incidents that were updated before this date
	UpdatedBefore *time.Time `queryParam:"style=form,explode=true,name=updated_before"`
	// A comma separated list of incident type IDs
	IncidentTypeID *string `queryParam:"style=form,explode=true,name=incident_type_id"`
	// A comma separated list of retrospective template IDs
	RetrospectiveTemplates *string                `queryParam:"style=form,explode=true,name=retrospective_templates"`
	CustomFieldID          *string                `queryParam:"style=form,explode=true,name=custom_field_id"`
	SortBy                 *ListMttxMetricsSortBy `queryParam:"style=form,explode=true,name=sort_by"`
	// Comma-separated list of measurements to include in the response
	Measurements *string `queryParam:"style=form,explode=true,name=measurements"`
	// Comma-separated list of label key / values in the format of 'key=value,key2=value2'
	Labels *string `queryParam:"style=form,explode=true,name=labels"`
	// Comma-separated list of user IDs for the incident openers
	IncidentOpeners *string `queryParam:"style=form,explode=true,name=incident_openers"`
	// Comma-separated list of ticket status states
	TicketStatus *string                     `queryParam:"style=form,explode=true,name=ticket_status"`
	RequestBody  *ListMttxMetricsRequestBody `request:"mediaType=multipart/form-data"`
}

func (l ListMttxMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMttxMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMttxMetricsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListMttxMetricsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListMttxMetricsRequest) GetConditions() *string {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *ListMttxMetricsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *ListMttxMetricsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *ListMttxMetricsRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *ListMttxMetricsRequest) GetExcludedInfrastructureIds() *string {
	if o == nil {
		return nil
	}
	return o.ExcludedInfrastructureIds
}

func (o *ListMttxMetricsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *ListMttxMetricsRequest) GetAssignedTeams() *string {
	if o == nil {
		return nil
	}
	return o.AssignedTeams
}

func (o *ListMttxMetricsRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ListMttxMetricsRequest) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}

func (o *ListMttxMetricsRequest) GetEndDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.EndDate
}

func (o *ListMttxMetricsRequest) GetResolvedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrAfter
}

func (o *ListMttxMetricsRequest) GetResolvedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrBefore
}

func (o *ListMttxMetricsRequest) GetClosedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClosedAtOrAfter
}

func (o *ListMttxMetricsRequest) GetClosedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClosedAtOrBefore
}

func (o *ListMttxMetricsRequest) GetCreatedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrAfter
}

func (o *ListMttxMetricsRequest) GetCreatedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrBefore
}

func (o *ListMttxMetricsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListMttxMetricsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListMttxMetricsRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *ListMttxMetricsRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *ListMttxMetricsRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *ListMttxMetricsRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *ListMttxMetricsRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *ListMttxMetricsRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

func (o *ListMttxMetricsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListMttxMetricsRequest) GetTagMatchStrategy() *ListMttxMetricsTagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *ListMttxMetricsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *ListMttxMetricsRequest) GetUpdatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAfter
}

func (o *ListMttxMetricsRequest) GetUpdatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedBefore
}

func (o *ListMttxMetricsRequest) GetIncidentTypeID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentTypeID
}

func (o *ListMttxMetricsRequest) GetRetrospectiveTemplates() *string {
	if o == nil {
		return nil
	}
	return o.RetrospectiveTemplates
}

func (o *ListMttxMetricsRequest) GetCustomFieldID() *string {
	if o == nil {
		return nil
	}
	return o.CustomFieldID
}

func (o *ListMttxMetricsRequest) GetSortBy() *ListMttxMetricsSortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *ListMttxMetricsRequest) GetMeasurements() *string {
	if o == nil {
		return nil
	}
	return o.Measurements
}

func (o *ListMttxMetricsRequest) GetLabels() *string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ListMttxMetricsRequest) GetIncidentOpeners() *string {
	if o == nil {
		return nil
	}
	return o.IncidentOpeners
}

func (o *ListMttxMetricsRequest) GetTicketStatus() *string {
	if o == nil {
		return nil
	}
	return o.TicketStatus
}

func (o *ListMttxMetricsRequest) GetRequestBody() *ListMttxMetricsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}
