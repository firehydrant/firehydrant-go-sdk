// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateIncidentLinkRequest struct {
	LinkID                              string                                         `pathParam:"style=simple,explode=false,name=link_id"`
	IncidentID                          string                                         `pathParam:"style=simple,explode=false,name=incident_id"`
	PutV1IncidentsIncidentIDLinksLinkID components.PutV1IncidentsIncidentIDLinksLinkID `request:"mediaType=application/json"`
}

func (o *UpdateIncidentLinkRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *UpdateIncidentLinkRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *UpdateIncidentLinkRequest) GetPutV1IncidentsIncidentIDLinksLinkID() components.PutV1IncidentsIncidentIDLinksLinkID {
	if o == nil {
		return components.PutV1IncidentsIncidentIDLinksLinkID{}
	}
	return o.PutV1IncidentsIncidentIDLinksLinkID
}

type UpdateIncidentLinkResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateIncidentLinkResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}