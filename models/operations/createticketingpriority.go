// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateTicketingPriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a single ticketing priority
	TicketingPriorityEntity *components.TicketingPriorityEntity
}

func (o *CreateTicketingPriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTicketingPriorityResponse) GetTicketingPriorityEntity() *components.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}
