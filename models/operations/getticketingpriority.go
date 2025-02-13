// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetTicketingPriorityRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTicketingPriorityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTicketingPriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single ticketing priority by ID
	TicketingPriorityEntity *components.TicketingPriorityEntity
}

func (o *GetTicketingPriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTicketingPriorityResponse) GetTicketingPriorityEntity() *components.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}
