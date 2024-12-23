// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteTicketingPriorityRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteTicketingPriorityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteTicketingPriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a single ticketing priority by ID
	TicketingPriorityEntity *components.TicketingPriorityEntity
}

func (o *DeleteTicketingPriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteTicketingPriorityResponse) GetTicketingPriorityEntity() *components.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}
