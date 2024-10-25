// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1TicketingPrioritiesIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteV1TicketingPrioritiesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteV1TicketingPrioritiesIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a single ticketing priority by ID
	TicketingPriorityEntity *components.TicketingPriorityEntity
}

func (o *DeleteV1TicketingPrioritiesIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1TicketingPrioritiesIDResponse) GetTicketingPriorityEntity() *components.TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.TicketingPriorityEntity
}