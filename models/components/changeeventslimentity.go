// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

// ChangeEventSlimEntityLabels - An object of label key and values
type ChangeEventSlimEntityLabels struct {
}

type ChangeEventSlimEntity struct {
	ID              *string                  `json:"id,omitempty"`
	Summary         *string                  `json:"summary,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	ExternalID      *string                  `json:"external_id,omitempty"`
	CreatedAt       *time.Time               `json:"created_at,omitempty"`
	UpdatedAt       *time.Time               `json:"updated_at,omitempty"`
	StartsAt        *time.Time               `json:"starts_at,omitempty"`
	EndsAt          *time.Time               `json:"ends_at,omitempty"`
	DurationMs      *int                     `json:"duration_ms,omitempty"`
	DurationIso8601 *string                  `json:"duration_iso8601,omitempty"`
	Environments    []EnvironmentEntryEntity `json:"environments,omitempty"`
	Authors         []AuthorEntity           `json:"authors,omitempty"`
	// An object of label key and values
	Labels   *ChangeEventSlimEntityLabels `json:"labels,omitempty"`
	Services []ServiceEntity              `json:"services,omitempty"`
}

func (c ChangeEventSlimEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChangeEventSlimEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChangeEventSlimEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeEventSlimEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *ChangeEventSlimEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ChangeEventSlimEntity) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *ChangeEventSlimEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ChangeEventSlimEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ChangeEventSlimEntity) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *ChangeEventSlimEntity) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *ChangeEventSlimEntity) GetDurationMs() *int {
	if o == nil {
		return nil
	}
	return o.DurationMs
}

func (o *ChangeEventSlimEntity) GetDurationIso8601() *string {
	if o == nil {
		return nil
	}
	return o.DurationIso8601
}

func (o *ChangeEventSlimEntity) GetEnvironments() []EnvironmentEntryEntity {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *ChangeEventSlimEntity) GetAuthors() []AuthorEntity {
	if o == nil {
		return nil
	}
	return o.Authors
}

func (o *ChangeEventSlimEntity) GetLabels() *ChangeEventSlimEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ChangeEventSlimEntity) GetServices() []ServiceEntity {
	if o == nil {
		return nil
	}
	return o.Services
}