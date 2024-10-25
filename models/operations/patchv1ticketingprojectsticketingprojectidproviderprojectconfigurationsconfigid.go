// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest struct {
	TicketingProjectID string `pathParam:"style=simple,explode=false,name=ticketing_project_id"`
	ConfigID           string `pathParam:"style=simple,explode=false,name=config_id"`
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetTicketingProjectID() string {
	if o == nil {
		return ""
	}
	return o.TicketingProjectID
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDRequest) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

type PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update configuration for a ticketing project
	TicketingProjectConfigEntity *components.TicketingProjectConfigEntity
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDResponse) GetTicketingProjectConfigEntity() *components.TicketingProjectConfigEntity {
	if o == nil {
		return nil
	}
	return o.TicketingProjectConfigEntity
}