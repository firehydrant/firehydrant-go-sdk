// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

type ChangeEventEntityAttachments struct {
}

// ChangeEventEntityLabels - An object of label key and values
type ChangeEventEntityLabels struct {
}

// ChangeEventEntity model
type ChangeEventEntity struct {
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
	RelatedChanges  []ChangeEntity           `json:"related_changes,omitempty"`
	Identities      []ChangeIdentityEntity   `json:"identities,omitempty"`
	Authors         []AuthorEntity           `json:"authors,omitempty"`
	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	Attachments []ChangeEventEntityAttachments `json:"attachments,omitempty"`
	// An object of label key and values
	Labels   *ChangeEventEntityLabels `json:"labels,omitempty"`
	Services []ServiceEntity          `json:"services,omitempty"`
}

func (c ChangeEventEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChangeEventEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChangeEventEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeEventEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *ChangeEventEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ChangeEventEntity) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *ChangeEventEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ChangeEventEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ChangeEventEntity) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *ChangeEventEntity) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *ChangeEventEntity) GetDurationMs() *int {
	if o == nil {
		return nil
	}
	return o.DurationMs
}

func (o *ChangeEventEntity) GetDurationIso8601() *string {
	if o == nil {
		return nil
	}
	return o.DurationIso8601
}

func (o *ChangeEventEntity) GetEnvironments() []EnvironmentEntryEntity {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *ChangeEventEntity) GetRelatedChanges() []ChangeEntity {
	if o == nil {
		return nil
	}
	return o.RelatedChanges
}

func (o *ChangeEventEntity) GetIdentities() []ChangeIdentityEntity {
	if o == nil {
		return nil
	}
	return o.Identities
}

func (o *ChangeEventEntity) GetAuthors() []AuthorEntity {
	if o == nil {
		return nil
	}
	return o.Authors
}

func (o *ChangeEventEntity) GetAttachments() []ChangeEventEntityAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *ChangeEventEntity) GetLabels() *ChangeEventEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ChangeEventEntity) GetServices() []ServiceEntity {
	if o == nil {
		return nil
	}
	return o.Services
}