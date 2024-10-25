// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

// AlertsAlertEntityLabels - Arbitrary key:value pairs of labels.
type AlertsAlertEntityLabels struct {
}

// AlertsAlertEntity - Alerts_AlertEntity model
type AlertsAlertEntity struct {
	ID              *string    `json:"id,omitempty"`
	Summary         *string    `json:"summary,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Priority        *string    `json:"priority,omitempty"`
	IntegrationName *string    `json:"integration_name,omitempty"`
	StartsAt        *time.Time `json:"starts_at,omitempty"`
	EndsAt          *time.Time `json:"ends_at,omitempty"`
	DurationMs      *int       `json:"duration_ms,omitempty"`
	DurationIso8601 *string    `json:"duration_iso8601,omitempty"`
	Status          *string    `json:"status,omitempty"`
	RemoteID        *string    `json:"remote_id,omitempty"`
	RemoteURL       *string    `json:"remote_url,omitempty"`
	// Arbitrary key:value pairs of labels.
	Labels       *AlertsAlertEntityLabels             `json:"labels,omitempty"`
	Environments []SuccinctEntity                     `json:"environments,omitempty"`
	Services     []SuccinctEntity                     `json:"services,omitempty"`
	Tags         []string                             `json:"tags,omitempty"`
	SourceIcon   *string                              `json:"source_icon,omitempty"`
	SignalID     *string                              `json:"signal_id,omitempty"`
	SignalRule   *SignalsAPIRuleEntity                `json:"signal_rule,omitempty"`
	TeamName     *string                              `json:"team_name,omitempty"`
	TeamID       *string                              `json:"team_id,omitempty"`
	Position     *int                                 `json:"position,omitempty"`
	Incidents    []PublicAPIV1IncidentsSuccinctEntity `json:"incidents,omitempty"`
	Events       []AlertsSirenEventEntity             `json:"events,omitempty"`
	IsExpired    *bool                                `json:"is_expired,omitempty"`
}

func (a AlertsAlertEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AlertsAlertEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AlertsAlertEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AlertsAlertEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *AlertsAlertEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AlertsAlertEntity) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *AlertsAlertEntity) GetIntegrationName() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationName
}

func (o *AlertsAlertEntity) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *AlertsAlertEntity) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *AlertsAlertEntity) GetDurationMs() *int {
	if o == nil {
		return nil
	}
	return o.DurationMs
}

func (o *AlertsAlertEntity) GetDurationIso8601() *string {
	if o == nil {
		return nil
	}
	return o.DurationIso8601
}

func (o *AlertsAlertEntity) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AlertsAlertEntity) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *AlertsAlertEntity) GetRemoteURL() *string {
	if o == nil {
		return nil
	}
	return o.RemoteURL
}

func (o *AlertsAlertEntity) GetLabels() *AlertsAlertEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *AlertsAlertEntity) GetEnvironments() []SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *AlertsAlertEntity) GetServices() []SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *AlertsAlertEntity) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AlertsAlertEntity) GetSourceIcon() *string {
	if o == nil {
		return nil
	}
	return o.SourceIcon
}

func (o *AlertsAlertEntity) GetSignalID() *string {
	if o == nil {
		return nil
	}
	return o.SignalID
}

func (o *AlertsAlertEntity) GetSignalRule() *SignalsAPIRuleEntity {
	if o == nil {
		return nil
	}
	return o.SignalRule
}

func (o *AlertsAlertEntity) GetTeamName() *string {
	if o == nil {
		return nil
	}
	return o.TeamName
}

func (o *AlertsAlertEntity) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *AlertsAlertEntity) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *AlertsAlertEntity) GetIncidents() []PublicAPIV1IncidentsSuccinctEntity {
	if o == nil {
		return nil
	}
	return o.Incidents
}

func (o *AlertsAlertEntity) GetEvents() []AlertsSirenEventEntity {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *AlertsAlertEntity) GetIsExpired() *bool {
	if o == nil {
		return nil
	}
	return o.IsExpired
}