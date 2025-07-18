// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// IntegrationsIntegrationEntity - Integrations_IntegrationEntity model
type IntegrationsIntegrationEntity struct {
	ID          *string                                          `json:"id,omitempty"`
	Slug        *string                                          `json:"slug,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	Description *string                                          `json:"description,omitempty"`
	SetupURL    *string                                          `json:"setup_url,omitempty"`
	CreatedAt   *time.Time                                       `json:"created_at,omitempty"`
	Connections []IntegrationsConnectionEntity                   `json:"connections,omitempty"`
	Enabled     *bool                                            `json:"enabled,omitempty"`
	Installed   *bool                                            `json:"installed,omitempty"`
	Deprecated  *bool                                            `json:"deprecated,omitempty"`
	Logo        *NullableIntegrationsIntegrationEntityLogoEntity `json:"logo,omitempty"`
	NatIP       *string                                          `json:"nat_ip,omitempty"`
}

func (i IntegrationsIntegrationEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IntegrationsIntegrationEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IntegrationsIntegrationEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationsIntegrationEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *IntegrationsIntegrationEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IntegrationsIntegrationEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IntegrationsIntegrationEntity) GetSetupURL() *string {
	if o == nil {
		return nil
	}
	return o.SetupURL
}

func (o *IntegrationsIntegrationEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IntegrationsIntegrationEntity) GetConnections() []IntegrationsConnectionEntity {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *IntegrationsIntegrationEntity) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *IntegrationsIntegrationEntity) GetInstalled() *bool {
	if o == nil {
		return nil
	}
	return o.Installed
}

func (o *IntegrationsIntegrationEntity) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *IntegrationsIntegrationEntity) GetLogo() *NullableIntegrationsIntegrationEntityLogoEntity {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *IntegrationsIntegrationEntity) GetNatIP() *string {
	if o == nil {
		return nil
	}
	return o.NatIP
}
