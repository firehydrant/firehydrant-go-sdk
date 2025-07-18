// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
)

type UpdateIncidentRetrospectiveRequest struct {
	RetrospectiveID             string                                 `pathParam:"style=simple,explode=false,name=retrospective_id"`
	IncidentID                  string                                 `pathParam:"style=simple,explode=false,name=incident_id"`
	UpdateIncidentRetrospective components.UpdateIncidentRetrospective `request:"mediaType=application/json"`
}

func (o *UpdateIncidentRetrospectiveRequest) GetRetrospectiveID() string {
	if o == nil {
		return ""
	}
	return o.RetrospectiveID
}

func (o *UpdateIncidentRetrospectiveRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *UpdateIncidentRetrospectiveRequest) GetUpdateIncidentRetrospective() components.UpdateIncidentRetrospective {
	if o == nil {
		return components.UpdateIncidentRetrospective{}
	}
	return o.UpdateIncidentRetrospective
}
