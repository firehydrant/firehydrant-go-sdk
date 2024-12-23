// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetPriorityRequest struct {
	PrioritySlug string `pathParam:"style=simple,explode=false,name=priority_slug"`
}

func (o *GetPriorityRequest) GetPrioritySlug() string {
	if o == nil {
		return ""
	}
	return o.PrioritySlug
}

type GetPriorityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a specific priority
	PriorityEntity *components.PriorityEntity
}

func (o *GetPriorityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPriorityResponse) GetPriorityEntity() *components.PriorityEntity {
	if o == nil {
		return nil
	}
	return o.PriorityEntity
}
