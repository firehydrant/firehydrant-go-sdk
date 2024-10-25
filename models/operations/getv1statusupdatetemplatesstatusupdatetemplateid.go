// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1StatusUpdateTemplatesStatusUpdateTemplateIDRequest struct {
	StatusUpdateTemplateID string `pathParam:"style=simple,explode=false,name=status_update_template_id"`
}

func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDRequest) GetStatusUpdateTemplateID() string {
	if o == nil {
		return ""
	}
	return o.StatusUpdateTemplateID
}

type GetV1StatusUpdateTemplatesStatusUpdateTemplateIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get a single status update template by ID
	StatusUpdateTemplateEntity *components.StatusUpdateTemplateEntity
}

func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDResponse) GetStatusUpdateTemplateEntity() *components.StatusUpdateTemplateEntity {
	if o == nil {
		return nil
	}
	return o.StatusUpdateTemplateEntity
}