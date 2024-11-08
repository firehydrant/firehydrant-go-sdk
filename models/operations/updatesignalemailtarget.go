// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateSignalEmailTargetRequest struct {
	ID                           string                                  `pathParam:"style=simple,explode=false,name=id"`
	PatchV1SignalsEmailTargetsID components.PatchV1SignalsEmailTargetsID `request:"mediaType=application/json"`
}

func (o *UpdateSignalEmailTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSignalEmailTargetRequest) GetPatchV1SignalsEmailTargetsID() components.PatchV1SignalsEmailTargetsID {
	if o == nil {
		return components.PatchV1SignalsEmailTargetsID{}
	}
	return o.PatchV1SignalsEmailTargetsID
}

type UpdateSignalEmailTargetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateSignalEmailTargetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}