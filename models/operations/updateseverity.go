// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type UpdateSeverityRequest struct {
	SeveritySlug   string                    `pathParam:"style=simple,explode=false,name=severity_slug"`
	UpdateSeverity components.UpdateSeverity `request:"mediaType=application/json"`
}

func (o *UpdateSeverityRequest) GetSeveritySlug() string {
	if o == nil {
		return ""
	}
	return o.SeveritySlug
}

func (o *UpdateSeverityRequest) GetUpdateSeverity() components.UpdateSeverity {
	if o == nil {
		return components.UpdateSeverity{}
	}
	return o.UpdateSeverity
}
