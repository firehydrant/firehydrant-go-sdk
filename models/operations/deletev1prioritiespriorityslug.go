// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1PrioritiesPrioritySlugRequest struct {
	PrioritySlug string `pathParam:"style=simple,explode=false,name=priority_slug"`
}

func (o *DeleteV1PrioritiesPrioritySlugRequest) GetPrioritySlug() string {
	if o == nil {
		return ""
	}
	return o.PrioritySlug
}

type DeleteV1PrioritiesPrioritySlugResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a specific priority
	PriorityEntity *components.PriorityEntity
}

func (o *DeleteV1PrioritiesPrioritySlugResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1PrioritiesPrioritySlugResponse) GetPriorityEntity() *components.PriorityEntity {
	if o == nil {
		return nil
	}
	return o.PriorityEntity
}