// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
	"time"
)

// ServiceEntityLabels - An object of label key and values
type ServiceEntityLabels struct {
}

// ServiceEntityManagedBySettings - Indicates the settings of the catalog that manages this service
type ServiceEntityManagedBySettings struct {
}

// ServiceEntity model
type ServiceEntity struct {
	ID            *string    `json:"id,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Slug          *string    `json:"slug,omitempty"`
	ServiceTier   *int       `json:"service_tier,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	AllowedParams []string   `json:"allowed_params,omitempty"`
	// An object of label key and values
	Labels                *ServiceEntityLabels `json:"labels,omitempty"`
	AlertOnAdd            *bool                `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam *bool                `json:"auto_add_responding_team,omitempty"`
	// List of active incident guids
	ActiveIncidents []string `json:"active_incidents,omitempty"`
	// List of checklists associated with a service
	Checklists      []ChecklistTemplateEntity `json:"checklists,omitempty"`
	CompletedChecks *int                      `json:"completed_checks,omitempty"`
	// Information about known linkages to representations of services outside of FireHydrant.
	ExternalResources []ExternalResourceEntity `json:"external_resources,omitempty"`
	// List of functionalities attached to the service
	Functionalities []FunctionalityEntity                    `json:"functionalities,omitempty"`
	LastImport      *NullableImportsImportableResourceEntity `json:"last_import,omitempty"`
	// List of links attached to this service.
	Links []LinksEntity `json:"links,omitempty"`
	// If set, this field indicates that the service is managed by an integration and thus cannot be set manually
	ManagedBy *string `json:"managed_by,omitempty"`
	// Indicates the settings of the catalog that manages this service
	ManagedBySettings         *ServiceEntityManagedBySettings `json:"managed_by_settings,omitempty"`
	Owner                     *NullableTeamEntityLite         `json:"owner,omitempty"`
	ServiceChecklistUpdatedAt *time.Time                      `json:"service_checklist_updated_at,omitempty"`
	// List of teams attached to the service
	Teams     []TeamEntityLite      `json:"teams,omitempty"`
	UpdatedBy *NullableAuthorEntity `json:"updated_by,omitempty"`
}

func (s ServiceEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ServiceEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ServiceEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *ServiceEntity) GetServiceTier() *int {
	if o == nil {
		return nil
	}
	return o.ServiceTier
}

func (o *ServiceEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ServiceEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ServiceEntity) GetAllowedParams() []string {
	if o == nil {
		return nil
	}
	return o.AllowedParams
}

func (o *ServiceEntity) GetLabels() *ServiceEntityLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ServiceEntity) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *ServiceEntity) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}

func (o *ServiceEntity) GetActiveIncidents() []string {
	if o == nil {
		return nil
	}
	return o.ActiveIncidents
}

func (o *ServiceEntity) GetChecklists() []ChecklistTemplateEntity {
	if o == nil {
		return nil
	}
	return o.Checklists
}

func (o *ServiceEntity) GetCompletedChecks() *int {
	if o == nil {
		return nil
	}
	return o.CompletedChecks
}

func (o *ServiceEntity) GetExternalResources() []ExternalResourceEntity {
	if o == nil {
		return nil
	}
	return o.ExternalResources
}

func (o *ServiceEntity) GetFunctionalities() []FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *ServiceEntity) GetLastImport() *NullableImportsImportableResourceEntity {
	if o == nil {
		return nil
	}
	return o.LastImport
}

func (o *ServiceEntity) GetLinks() []LinksEntity {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *ServiceEntity) GetManagedBy() *string {
	if o == nil {
		return nil
	}
	return o.ManagedBy
}

func (o *ServiceEntity) GetManagedBySettings() *ServiceEntityManagedBySettings {
	if o == nil {
		return nil
	}
	return o.ManagedBySettings
}

func (o *ServiceEntity) GetOwner() *NullableTeamEntityLite {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *ServiceEntity) GetServiceChecklistUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ServiceChecklistUpdatedAt
}

func (o *ServiceEntity) GetTeams() []TeamEntityLite {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *ServiceEntity) GetUpdatedBy() *NullableAuthorEntity {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}
