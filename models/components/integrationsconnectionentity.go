// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// IntegrationsConnectionEntityDetails - Integration-specific details of this connection. As identified by the integration_slug, this object will be represented by that integration's ConnectionEntity.
type IntegrationsConnectionEntityDetails struct {
}

// IntegrationsConnectionEntity - Integrations_ConnectionEntity model
type IntegrationsConnectionEntity struct {
	ID               *string    `json:"id,omitempty"`
	IntegrationSlug  *string    `json:"integration_slug,omitempty"`
	IntegrationID    *string    `json:"integration_id,omitempty"`
	DisplayName      *string    `json:"display_name,omitempty"`
	ConfigurationURL *string    `json:"configuration_url,omitempty"`
	AuthorizedBy     *string    `json:"authorized_by,omitempty"`
	AuthorizedByID   *string    `json:"authorized_by_id,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	// Integration-specific details of this connection. As identified by the integration_slug, this object will be represented by that integration's ConnectionEntity.
	Details                *IntegrationsConnectionEntityDetails `json:"details,omitempty"`
	DefaultAuthorizedActor *NullableAuthorEntity                `json:"default_authorized_actor,omitempty"`
}

func (i IntegrationsConnectionEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IntegrationsConnectionEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IntegrationsConnectionEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationsConnectionEntity) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *IntegrationsConnectionEntity) GetIntegrationID() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationID
}

func (o *IntegrationsConnectionEntity) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *IntegrationsConnectionEntity) GetConfigurationURL() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationURL
}

func (o *IntegrationsConnectionEntity) GetAuthorizedBy() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizedBy
}

func (o *IntegrationsConnectionEntity) GetAuthorizedByID() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizedByID
}

func (o *IntegrationsConnectionEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IntegrationsConnectionEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *IntegrationsConnectionEntity) GetDetails() *IntegrationsConnectionEntityDetails {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *IntegrationsConnectionEntity) GetDefaultAuthorizedActor() *NullableAuthorEntity {
	if o == nil {
		return nil
	}
	return o.DefaultAuthorizedActor
}
