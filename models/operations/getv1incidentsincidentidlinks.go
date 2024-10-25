// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1IncidentsIncidentIDLinksRequest struct {
	Page       *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage    *int   `queryParam:"style=form,explode=true,name=per_page"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetV1IncidentsIncidentIDLinksRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IncidentsIncidentIDLinksRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1IncidentsIncidentIDLinksRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetV1IncidentsIncidentIDLinksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all the editable, external incident links attached to an incident
	AttachmentsLinkEntityPaginated *components.AttachmentsLinkEntityPaginated
}

func (o *GetV1IncidentsIncidentIDLinksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDLinksResponse) GetAttachmentsLinkEntityPaginated() *components.AttachmentsLinkEntityPaginated {
	if o == nil {
		return nil
	}
	return o.AttachmentsLinkEntityPaginated
}